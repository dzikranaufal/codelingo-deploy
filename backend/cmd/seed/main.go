// Seed data bahasa C: 4 materi (learning path) x 5 lesson x 10 soal = 200 soal.
//
// Isi soal ada di file data_*.go pada folder yang sama (package main).
//
// Penggunaan:
//
//	go run ./cmd/seed               # aman diulang: path yang sudah ada dilewati
//	go run ./cmd/seed -reset        # HAPUS semua path bahasa C lalu seed ulang (minta konfirmasi)
//	go run ./cmd/seed -reset -yes   # sama, tanpa konfirmasi
//
// Tidak ada perubahan skema, jadi TIDAK perlu migrasi ulang.
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"codelingo/backend/internal/config"
	"codelingo/backend/internal/database"
	"codelingo/backend/internal/models"

	"gorm.io/gorm"
)

const (
	questionType       = "multiple_choice"
	xpPerAnswer        = 10
	lessonsPerPath     = 5
	questionsPerLesson = 10
	optionsPerQuestion = 4
)

type seedQuestion struct {
	Text        string
	Explanation string
	Options     []string
	Correct     int // indeks opsi yang benar, dimulai dari 0
}

type seedLesson struct {
	Title       string
	Description string
	Questions   []seedQuestion
}

type seedPath struct {
	Name        string
	Description string
	OrderIndex  int
	Lessons     []seedLesson
}

// q: helper singkat untuk membuat soal.
func q(text, explanation string, correct int, options ...string) seedQuestion {
	return seedQuestion{Text: text, Explanation: explanation, Options: options, Correct: correct}
}

// Urutan di sini = urutan materi yang tampil (OrderIndex di tiap file data harus sama).
var paths = []seedPath{
	basicConcepts,
	conditionalsLoops,
	arraysStrings,
	functionsPointers,
}

// validate memastikan data seed konsisten SEBELUM menyentuh database.
func validate() error {
	seenPath := map[string]bool{}
	seenOrder := map[int]bool{}

	for _, sp := range paths {
		if seenPath[sp.Name] {
			return fmt.Errorf("nama path ganda: %q", sp.Name)
		}
		seenPath[sp.Name] = true
		if seenOrder[sp.OrderIndex] {
			return fmt.Errorf("path %q: OrderIndex %d dipakai dua kali", sp.Name, sp.OrderIndex)
		}
		seenOrder[sp.OrderIndex] = true

		if len(sp.Lessons) != lessonsPerPath {
			return fmt.Errorf("path %q: %d lesson, seharusnya %d", sp.Name, len(sp.Lessons), lessonsPerPath)
		}

		seenQuestion := map[string]bool{}
		for li, sl := range sp.Lessons {
			if len(sl.Questions) != questionsPerLesson {
				return fmt.Errorf("path %q lesson %d (%q): %d soal, seharusnya %d",
					sp.Name, li+1, sl.Title, len(sl.Questions), questionsPerLesson)
			}
			for qi, sq := range sl.Questions {
				where := fmt.Sprintf("path %q lesson %d soal %d", sp.Name, li+1, qi+1)

				if sq.Text == "" || sq.Explanation == "" {
					return fmt.Errorf("%s: teks atau penjelasan kosong", where)
				}
				if seenQuestion[sq.Text] {
					return fmt.Errorf("%s: teks soal ganda dalam satu path", where)
				}
				seenQuestion[sq.Text] = true

				if len(sq.Options) != optionsPerQuestion {
					return fmt.Errorf("%s: %d opsi, seharusnya %d", where, len(sq.Options), optionsPerQuestion)
				}
				if sq.Correct < 0 || sq.Correct >= len(sq.Options) {
					return fmt.Errorf("%s: indeks jawaban benar %d di luar jangkauan", where, sq.Correct)
				}
				seenOpt := map[string]bool{}
				for _, o := range sq.Options {
					if o == "" {
						return fmt.Errorf("%s: ada opsi kosong", where)
					}
					if seenOpt[o] {
						return fmt.Errorf("%s: opsi ganda %q", where, o)
					}
					seenOpt[o] = true
				}
			}
		}
	}
	return nil
}

// ensureLanguage memastikan bahasa C ada.
func ensureLanguage(db *gorm.DB) (models.Language, error) {
	var lang models.Language
	err := db.Where("slug = ?", "c").First(&lang).Error
	if err == nil {
		return lang, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return lang, err
	}
	lang = models.Language{Name: "C", Slug: "c", Description: "Belajar pemrograman dengan bahasa C.", IsAvailable: true}
	return lang, db.Create(&lang).Error
}

// resetLanguagePaths menghapus SEMUA learning path milik bahasa C beserta
// lesson, soal, dan opsinya, dalam satu transaksi (gagal = rollback total).
// Jika ada tabel lain (mis. progres user) yang merujuk data ini lewat foreign key,
// penghapusan akan gagal dan tidak ada yang berubah.
func resetLanguagePaths(db *gorm.DB, lang models.Language) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var oldPaths []models.LearningPath
		if err := tx.Where("language_id = ?", lang.ID).Find(&oldPaths).Error; err != nil {
			return err
		}
		if len(oldPaths) == 0 {
			log.Println("reset: tidak ada path bahasa C yang perlu dihapus")
			return nil
		}

		pathIDs := make([]any, 0, len(oldPaths))
		for _, p := range oldPaths {
			pathIDs = append(pathIDs, p.ID)
		}

		var links []models.LearningPathLesson
		if err := tx.Where("learning_path_id IN ?", pathIDs).Find(&links).Error; err != nil {
			return err
		}
		lessonIDs := make([]any, 0, len(links))
		for _, l := range links {
			lessonIDs = append(lessonIDs, l.LessonID)
		}

		questionCount := 0
		if len(lessonIDs) > 0 {
			var qs []models.Question
			if err := tx.Where("lesson_id IN ?", lessonIDs).Find(&qs).Error; err != nil {
				return err
			}
			questionCount = len(qs)
			if len(qs) > 0 {
				questionIDs := make([]any, 0, len(qs))
				for _, x := range qs {
					questionIDs = append(questionIDs, x.ID)
				}
				if err := tx.Unscoped().Where("question_id IN ?", questionIDs).Delete(&models.AnswerOption{}).Error; err != nil {
					return err
				}
				if err := tx.Unscoped().Where("lesson_id IN ?", lessonIDs).Delete(&models.Question{}).Error; err != nil {
					return err
				}
			}
		}

		if err := tx.Unscoped().Where("learning_path_id IN ?", pathIDs).Delete(&models.LearningPathLesson{}).Error; err != nil {
			return err
		}
		if len(lessonIDs) > 0 {
			if err := tx.Unscoped().Where("id IN ?", lessonIDs).Delete(&models.Lesson{}).Error; err != nil {
				return err
			}
		}
		if err := tx.Unscoped().Where("id IN ?", pathIDs).Delete(&models.LearningPath{}).Error; err != nil {
			return err
		}

		log.Printf("reset: dihapus %d path, %d lesson, %d soal", len(oldPaths), len(lessonIDs), questionCount)
		return nil
	})
}

// seedPathData membuat satu path beserta lesson, soal, dan opsinya dalam SATU transaksi.
func seedPathData(db *gorm.DB, lang models.Language, sp seedPath) (created bool, questions int, err error) {
	err = db.Transaction(func(tx *gorm.DB) error {
		var existing int64
		if err := tx.Model(&models.LearningPath{}).Where("name = ?", sp.Name).Count(&existing).Error; err != nil {
			return err
		}
		if existing > 0 {
			return nil // sudah ada, lewati
		}

		path := models.LearningPath{
			Name:        sp.Name,
			Description: sp.Description,
			LanguageID:  &lang.ID,
			OrderIndex:  sp.OrderIndex,
		}
		if err := tx.Create(&path).Error; err != nil {
			return err
		}

		for i, sl := range sp.Lessons {
			lesson := models.Lesson{Title: sl.Title, Description: sl.Description, Level: 1}
			if err := tx.Create(&lesson).Error; err != nil {
				return err
			}

			// Urutan lesson disimpan di tabel penghubung, bukan di lessons.
			link := models.LearningPathLesson{
				LearningPathID: path.ID,
				LessonID:       lesson.ID,
				OrderIndex:     i + 1,
			}
			if err := tx.Create(&link).Error; err != nil {
				return err
			}

			for _, sq := range sl.Questions {
				question := models.Question{
					LessonID:     lesson.ID,
					QuestionText: sq.Text,
					QuestionType: questionType,
					XPReward:     xpPerAnswer,
					Explanation:  sq.Explanation,
				}
				if err := tx.Create(&question).Error; err != nil {
					return err
				}
				questions++

				for j, text := range sq.Options {
					option := models.AnswerOption{
						QuestionID: question.ID,
						OptionText: text,
						IsCorrect:  j == sq.Correct,
					}
					if err := tx.Create(&option).Error; err != nil {
						return err
					}
				}
			}
		}
		created = true
		return nil
	})
	return created, questions, err
}

func main() {
	reset := flag.Bool("reset", false, "hapus SEMUA learning path bahasa C (beserta lesson & soal) sebelum seed")
	yes := flag.Bool("yes", false, "lewati konfirmasi saat memakai -reset")
	flag.Parse()

	if err := validate(); err != nil {
		log.Fatalf("data seed tidak valid: %v", err)
	}

	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("gagal terhubung ke database: %v", err)
	}

	lang, err := ensureLanguage(db)
	if err != nil {
		log.Fatalf("gagal menyiapkan bahasa C: %v", err)
	}

	if *reset {
		if !*yes {
			fmt.Println("PERINGATAN: semua learning path bahasa C beserta lesson dan soalnya akan DIHAPUS dari database ini.")
			fmt.Print("Ketik YA (huruf besar) untuk lanjut: ")
			var answer string
			fmt.Scanln(&answer)
			if answer != "YA" {
				log.Fatal("dibatalkan, tidak ada data yang diubah")
			}
		}
		if err := resetLanguagePaths(db, lang); err != nil {
			log.Fatalf("reset gagal (rollback, tidak ada data yang berubah): %v", err)
		}
	}

	createdPaths, createdQuestions, failed := 0, 0, 0
	for _, sp := range paths {
		created, n, err := seedPathData(db, lang, sp)
		switch {
		case err != nil:
			failed++
			log.Printf("GAGAL path %q (rollback, tidak ada data tersimpan untuk path ini): %v", sp.Name, err)
		case !created:
			log.Printf("dilewati: path %q sudah ada", sp.Name)
		default:
			createdPaths++
			createdQuestions += n
			log.Printf("dibuat: path %q (%d lesson, %d soal)", sp.Name, len(sp.Lessons), n)
		}
	}

	log.Printf("seed selesai: %d path baru, %d soal baru, %d gagal", createdPaths, createdQuestions, failed)
	if failed > 0 {
		log.Fatal("ada path yang gagal di-seed, cek log di atas")
	}
}