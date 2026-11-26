package main

// Materi 1: Basic Concepts (5 lesson x 10 soal = 50 soal)
var basicConcepts = seedPath{
	Name:        "Basic Concepts",
	Description: "Dasar-dasar bahasa C: struktur program, variabel, tipe data, input/output, dan operator.",
	OrderIndex:  1,
	Lessons: []seedLesson{
		// ------------------------------------------------------------ Lesson 1
		{
			Title:       "Struktur Program dan Output",
			Description: "Kerangka program C, fungsi main, dan mencetak teks dengan printf.",
			Questions: []seedQuestion{
				q("Fungsi apa yang pertama kali dijalankan ketika program C dimulai?",
					"Eksekusi program C selalu dimulai dari fungsi main().", 1,
					"start()", "main()", "init()", "run()"),
				q("Apa fungsi baris #include <stdio.h> di awal program?",
					"Baris ini menyertakan header standar input/output, sehingga fungsi seperti printf() dan scanf() dapat dipakai.", 0,
					"Menyertakan fungsi input/output standar seperti printf dan scanf",
					"Mendeklarasikan variabel global",
					"Mengakhiri program",
					"Membuat komentar"),
				q("Tanda apa yang mengakhiri sebagian besar pernyataan (statement) di C?",
					"Setiap pernyataan di C diakhiri dengan titik koma (;).", 1,
					":", ";", ".", ","),
				q("Fungsi apa yang dipakai untuk menampilkan teks ke layar di C?",
					"printf() dari stdio.h dipakai untuk mencetak teks terformat ke layar.", 2,
					"echo()", "print()", "printf()", "cout"),
				q(`Apa fungsi escape sequence \n di dalam string printf?`,
					`\n menyisipkan karakter baris baru, sehingga teks berikutnya dicetak di baris selanjutnya.`, 3,
					"Menambah satu spasi",
					"Menghapus karakter sebelumnya",
					"Mengakhiri program",
					"Pindah ke baris baru"),
				q("Manakah cara menulis komentar satu baris di C (standar C99 ke atas)?",
					"// dipakai untuk komentar satu baris, sedangkan /* ... */ untuk komentar beberapa baris.", 0,
					"// ini komentar",
					"# ini komentar",
					"-- ini komentar",
					"<!-- ini komentar -->"),
				q(`Apa output kode berikut?

printf("Belajar ");
printf("C");`,
					"printf tidak menambah baris baru secara otomatis. Spasi setelah kata Belajar ikut tercetak, hasilnya: Belajar C.", 2,
					"C Belajar",
					"BelajarC",
					"Belajar C",
					"Belajar lalu C di baris baru"),
				q("Apa arti return 0; di akhir fungsi main()?",
					"Nilai 0 dikembalikan ke sistem operasi sebagai tanda bahwa program selesai tanpa error.", 3,
					"Mencetak angka 0 ke layar",
					"Program diulang dari awal",
					"Menghapus semua variabel",
					"Program selesai dengan sukses"),
				q("Apa yang terjadi jika kita lupa menulis titik koma di akhir sebuah pernyataan?",
					"Compiler akan melaporkan kesalahan sintaks (syntax error) dan program tidak bisa dikompilasi.", 1,
					"Program tetap berjalan normal",
					"Compile error karena kesalahan sintaks",
					"Compiler menambahkannya otomatis",
					"Program mencetak angka 0"),
				q(`Manakah penulisan yang benar untuk mencetak teks Hai di C?`,
					"C bersifat case-sensitive: nama fungsi harus ditulis persis printf, huruf kecil semua.", 2,
					`Printf("Hai");`,
					`PRINTF("Hai");`,
					`printf("Hai");`,
					`print f("Hai");`),
			},
		},

		// ------------------------------------------------------------ Lesson 2
		{
			Title:       "Variabel dan Konstanta",
			Description: "Mendeklarasikan variabel, aturan penamaan, dan membuat nilai tetap.",
			Questions: []seedQuestion{
				q("Apa yang dimaksud dengan deklarasi variabel?",
					"Deklarasi memberi tahu compiler tipe dan nama variabel agar disediakan ruang di memori.", 3,
					"Menghapus variabel dari memori",
					"Mencetak nilai variabel",
					"Mengubah tipe variabel",
					"Menentukan tipe dan nama variabel agar compiler menyediakan memori"),
				q("Manakah yang merupakan inisialisasi variabel (deklarasi sekaligus memberi nilai awal)?",
					"int x = 5; mendeklarasikan x sekaligus memberinya nilai awal 5.", 2,
					"int x;", "x = 5;", "int x = 5;", "5 = x;"),
				q("Manakah nama variabel yang valid di C?",
					"Nama variabel boleh berisi huruf, angka, dan underscore, tetapi tidak boleh diawali angka dan tidak boleh mengandung spasi atau tanda minus.", 0,
					"total_harga", "2total", "total-harga", "total harga"),
				q("Apakah umur dan Umur dianggap variabel yang sama di C?",
					"C bersifat case-sensitive, sehingga umur dan Umur adalah dua nama yang berbeda.", 1,
					"Ya, keduanya sama persis",
					"Tidak, keduanya berbeda",
					"Sama, tetapi hanya di dalam main()",
					"Compile error jika keduanya dipakai"),
				q("Kata kunci apa yang membuat nilai sebuah variabel tidak bisa diubah setelah diinisialisasi?",
					"const menandai variabel sebagai konstanta: nilainya tidak boleh diubah setelah diberi nilai awal.", 3,
					"static", "final", "fixed", "const"),
				q("Manakah cara mendefinisikan konstanta PI dengan preprosesor?",
					"#define dipakai tanpa tanda sama dengan dan tanpa titik koma di akhir baris.", 1,
					"const PI 3.14;",
					"#define PI 3.14",
					"define PI = 3.14;",
					"#const PI 3.14"),
				q("Manakah yang TIDAK boleh dipakai sebagai nama variabel karena merupakan kata kunci (keyword) C?",
					"Kata kunci seperti int, float, if, dan while sudah punya arti khusus sehingga tidak boleh dipakai sebagai nama variabel.", 0,
					"int", "angka", "nilai", "total"),
				q(`Apa output kode berikut?

int a = 5;
int b = a;
a = 10;
printf("%d", b);`,
					"b menyalin nilai a pada saat itu (5). Mengubah a setelahnya tidak memengaruhi b.", 2,
					"10", "15", "5", "0"),
				q(`Apa yang terjadi pada kode berikut?

const int x = 5;
x = 10;`,
					"Variabel const tidak boleh diubah, sehingga compiler melaporkan error.", 3,
					"x menjadi 10",
					"x tetap 5 tanpa pesan apa pun",
					"Program berhenti saat dijalankan",
					"Compile error karena variabel const tidak boleh diubah"),
				q("Variabel lokal yang dideklarasikan tetapi belum diberi nilai awal berisi apa?",
					"Variabel lokal yang belum diinisialisasi berisi nilai sampah (garbage) yang tidak bisa diprediksi.", 1,
					"Selalu 0",
					"Nilai sampah yang tidak bisa diprediksi",
					"Selalu NULL",
					"Nilai maksimum dari tipenya"),
			},
		},

		// ------------------------------------------------------------ Lesson 3
		{
			Title:       "Tipe Data",
			Description: "int, float, double, char, sizeof, pembagian bilangan bulat, dan casting.",
			Questions: []seedQuestion{
				q("Tipe data apa yang paling cocok untuk menyimpan bilangan pecahan seperti 3.14?",
					"float dan double dipakai untuk bilangan pecahan, sedangkan int hanya untuk bilangan bulat.", 0,
					"float", "int", "char", "void"),
				q(`Manakah penulisan nilai karakter yang benar di C?`,
					"Karakter tunggal ditulis dengan tanda kutip tunggal ('A'), sedangkan kutip ganda dipakai untuk string.", 3,
					`"A"`, `A`, `"AB"`, `'A'`),
				q("Apa perbedaan utama double dibandingkan float?",
					"double memiliki presisi dan jangkauan lebih besar daripada float, tetapi memakai memori lebih banyak.", 1,
					"double hanya untuk bilangan bulat",
					"double memiliki presisi lebih tinggi daripada float",
					"double hanya menyimpan karakter",
					"Tidak ada perbedaan sama sekali"),
				q("Operator apa yang dipakai untuk mengetahui ukuran sebuah tipe data dalam byte?",
					"sizeof menghasilkan ukuran tipe data atau variabel dalam satuan byte.", 2,
					"length", "size()", "sizeof", "bytes()"),
				q(`Apa output kode berikut?

printf("%d", 9 / 2);`,
					"Kedua operand bertipe int sehingga hasilnya int. Bagian pecahan dibuang: 9 / 2 = 4.", 0,
					"4", "4.5", "5", "4.0"),
				q(`Apa output kode berikut?

printf("%.1f", 9 / 2.0);`,
					"Salah satu operand bertipe double, sehingga pembagian dilakukan sebagai bilangan pecahan: 4.5.", 2,
					"4.0", "4", "4.5", "5.0"),
				q(`Apa output kode berikut?

int x = (int) 3.9;
printf("%d", x);`,
					"Casting ke int membuang bagian pecahan (bukan membulatkan), sehingga 3.9 menjadi 3.", 1,
					"4", "3", "3.9", "0"),
				q("Manakah tipe data berikut yang TIDAK tersedia sebagai tipe bawaan C?",
					"C tidak punya tipe string bawaan. String direpresentasikan sebagai array char.", 3,
					"int", "char", "float", "string"),
				q("Mengapa char bisa dipakai seperti bilangan bulat kecil?",
					"Karakter disimpan sebagai kode angka (misalnya ASCII), sehingga 'A' bernilai 65.", 0,
					"Karena karakter disimpan sebagai kode angka (misalnya ASCII)",
					"Karena char sama dengan float",
					"Karena char tidak memakai memori",
					"Karena char menyimpan string"),
				q(`Apa output kode berikut?

char c = 'A';
printf("%d", c);`,
					"%d mencetak kode ASCII dari karakter. Kode untuk 'A' adalah 65.", 2,
					"A", "a", "65", "1"),
			},
		},

		// ------------------------------------------------------------ Lesson 4
		{
			Title:       "Input dan Format Output",
			Description: "Membaca input dengan scanf dan mengatur tampilan dengan printf.",
			Questions: []seedQuestion{
				q("Fungsi apa yang dipakai untuk membaca input dari keyboard di C?",
					"scanf() dari stdio.h membaca input terformat dari keyboard.", 1,
					"input()", "scanf()", "read()", "cin"),
				q("Format specifier apa yang dipakai untuk membaca atau mencetak bilangan bulat int?",
					"%d untuk int, %f untuk float, %c untuk char, dan %s untuk string.", 3,
					"%f", "%c", "%s", "%d"),
				q("Format specifier apa yang dipakai untuk mencetak satu karakter?",
					"%c dipakai untuk satu karakter, sedangkan %s untuk string.", 0,
					"%c", "%s", "%d", "%f"),
				q(`Mengapa scanf("%d", &umur); memakai tanda & sebelum umur?`,
					"scanf membutuhkan alamat memori variabel untuk menyimpan hasil input, dan operator & mengambil alamat tersebut.", 2,
					"Agar nilai input otomatis dikalikan dua",
					"Agar input dibaca sebagai teks",
					"Karena scanf membutuhkan alamat memori variabel untuk menyimpan hasil input",
					"Karena umur adalah konstanta"),
				q(`Apa output kode berikut?

printf("%.2f", 3.14159);`,
					"%.2f mencetak bilangan pecahan dengan 2 angka di belakang koma: 3.14.", 1,
					"3.1", "3.14", "3.142", "3.14159"),
				q(`Apa output kode berikut?

printf("Nilai: %d", 80);`,
					"%d diganti dengan nilai 80, sehingga tercetak Nilai: 80.", 0,
					"Nilai: 80", "Nilai: %d", "80", "Nilai: d"),
				q(`Apa output kode berikut?

printf("%d + %d = %d", 2, 3, 2 + 3);`,
					"Ketiga %d diganti berurutan dengan 2, 3, dan 5, sehingga tercetak 2 + 3 = 5.", 3,
					"5 = 2 + 3", "2 + 3 = 23", "2 + 3 = 2 + 3", "2 + 3 = 5"),
				q(`Escape sequence apa yang menyisipkan tab (jarak horizontal) pada output?`,
					`\t menyisipkan karakter tab, sedangkan \n untuk baris baru.`, 2,
					`\n`, `\b`, `\t`, `\0`),
				q("Bagaimana cara mencetak tanda persen (%) dengan printf?",
					"Untuk mencetak karakter %, tulis %% di dalam format string.", 1,
					`printf("%");`,
					`printf("%%");`,
					`printf("\%");`,
					`printf("%p");`),
				q(`Jika pengguna memasukkan 4 dan 6, apa output kode berikut?

int a, b;
scanf("%d %d", &a, &b);
printf("%d", a + b);`,
					"scanf membaca dua bilangan: a = 4 dan b = 6, sehingga a + b = 10.", 2,
					"46", "24", "10", "Error"),
			},
		},

		// ------------------------------------------------------------ Lesson 5
		{
			Title:       "Operator dan Ekspresi",
			Description: "Operator aritmatika, penugasan, increment, perbandingan, dan logika.",
			Questions: []seedQuestion{
				q("Berapa hasil 17 % 5 di C?",
					"Operator % (modulo) menghasilkan sisa pembagian. 17 dibagi 5 = 3 sisa 2.", 2,
					"3", "3.4", "2", "12"),
				q("Berapa hasil ekspresi (2 + 3) * 4?",
					"Tanda kurung didahulukan: 2 + 3 = 5, lalu 5 * 4 = 20.", 0,
					"20", "14", "24", "9"),
				q(`Apa output kode berikut?

int x = 10;
x -= 4;
printf("%d", x);`,
					"x -= 4 sama dengan x = x - 4, sehingga x menjadi 6.", 1,
					"14", "6", "40", "-6"),
				q(`Apa output kode berikut?

int i = 5;
printf("%d", i++);`,
					"i++ (post-increment) memakai nilai i terlebih dahulu (5), baru menambahnya. Jadi yang tercetak 5, dan i menjadi 6 setelahnya.", 3,
					"4", "6", "10", "5"),
				q(`Apa output kode berikut?

int i = 5;
printf("%d", ++i);`,
					"++i (pre-increment) menambah i terlebih dahulu menjadi 6, baru nilainya dipakai.", 2,
					"5", "7", "6", "4"),
				q("Apa hasil ekspresi 5 > 3 di C?",
					"Ekspresi perbandingan menghasilkan 1 jika benar dan 0 jika salah.", 0,
					"1", "0", "5", "true"),
				q("Operator logika apa yang bernilai benar jika salah satu atau kedua kondisi bernilai benar?",
					"|| adalah operator ATAU (OR), sedangkan && adalah DAN (AND).", 1,
					"&&", "||", "!", "=="),
				q(`Apa output kode berikut?

int a = 7, b = 2;
printf("%d", a / b + a % b);`,
					"7 / 2 = 3 (pembagian bilangan bulat) dan 7 % 2 = 1, sehingga 3 + 1 = 4.", 3,
					"3", "5", "4.5", "4"),
				q("Apa fungsi operator = di C?",
					"= adalah operator penugasan (assignment) yang menyimpan nilai ke variabel, sedangkan == dipakai untuk membandingkan.", 0,
					"Menyimpan nilai ke variabel (penugasan)",
					"Membandingkan dua nilai",
					"Memeriksa ketidaksamaan",
					"Menambah nilai variabel"),
				q(`Apa output kode berikut?

int x = 3;
int y = (x > 1) && (x < 3);
printf("%d", y);`,
					"x > 1 benar, tetapi x < 3 salah. Karena && butuh keduanya benar, hasilnya 0.", 1,
					"1", "0", "3", "Error"),
			},
		},
	},
}