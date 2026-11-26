package main

// Materi 2: Conditionals and Loops (5 lesson x 10 soal = 50 soal)
var conditionalsLoops = seedPath{
	Name:        "Conditionals and Loops",
	Description: "Mengambil keputusan dan mengulang perintah: if-else, switch, for, while, do-while, break, dan continue.",
	OrderIndex:  2,
	Lessons: []seedLesson{
		// ------------------------------------------------------------ Lesson 1
		{
			Title:       "if, else, dan else if",
			Description: "Mengambil keputusan berdasarkan kondisi.",
			Questions: []seedQuestion{
				q("Apa yang dilakukan pernyataan if?",
					"if menjalankan blok kode hanya ketika kondisinya bernilai benar.", 2,
					"Mengulang kode berkali-kali",
					"Menghentikan program",
					"Menjalankan blok kode hanya jika kondisinya benar",
					"Mendeklarasikan variabel"),
				q("Di C, nilai apa yang dianggap SALAH (false) dalam sebuah kondisi?",
					"Nilai 0 dianggap salah, sedangkan semua nilai selain 0 dianggap benar.", 0,
					"0", "1", "-1", "Semua nilai negatif"),
				q(`Apa output kode berikut?

int x = 7;
if (x % 2 == 0) printf("Genap");
else printf("Ganjil");`,
					"7 % 2 = 1, bukan 0, sehingga kondisi salah dan bagian else dijalankan.", 3,
					"Genap", "GenapGanjil", "Tidak ada output", "Ganjil"),
				q(`Apa output kode berikut?

int n = 15;
if (n > 10) printf("A");
else if (n > 5) printf("B");
else printf("C");`,
					"Rangkaian if-else if berhenti pada kondisi pertama yang benar. n > 10 benar, jadi hanya A yang dicetak.", 1,
					"AB", "A", "B", "C"),
				q(`Apa output kode berikut?

int x = 5;
if (x = 0) printf("Nol");
else printf("Bukan nol");`,
					"x = 0 adalah penugasan, bukan perbandingan. Ekspresinya bernilai 0 (salah), sehingga bagian else dijalankan.", 2,
					"Nol", "Compile error", "Bukan nol", "Tidak ada output"),
				q("Kapan blok else dijalankan?",
					"else dijalankan ketika kondisi pada if di atasnya bernilai salah.", 0,
					"Ketika kondisi if bernilai salah",
					"Selalu, setelah if",
					"Ketika kondisi if bernilai benar",
					"Hanya jika ada break"),
				q(`Apa output kode berikut jika x = 2?

int x = 2;
if (x > 5) printf("A");
printf("B");`,
					`Tanpa kurung kurawal, hanya pernyataan tepat setelah if yang menjadi bagiannya. printf("B") selalu dijalankan.`, 1,
					"A", "B", "AB", "Tidak ada output"),
				q("Mengapa disarankan memakai kurung kurawal { } pada if walau hanya satu pernyataan?",
					"Kurung kurawal membuat batas blok jelas dan mencegah bug ketika pernyataan baru ditambahkan nanti.", 3,
					"Agar program berjalan lebih cepat",
					"Karena if tanpa kurung kurawal selalu error",
					"Agar kondisinya otomatis benar",
					"Agar jelas pernyataan mana yang termasuk blok dan terhindar dari bug"),
				q(`Apa output kode berikut?

int nilai = 85;
if (nilai >= 90) printf("A");
else if (nilai >= 80) printf("B");
else printf("C");`,
					"85 tidak memenuhi >= 90, tetapi memenuhi >= 80, sehingga yang tercetak B.", 2,
					"A", "C", "B", "AB"),
				q("Manakah kondisi yang benar untuk memeriksa apakah x berada di antara 1 dan 10 (inklusif)?",
					"Gabungkan dua perbandingan dengan &&. Penulisan 1 <= x <= 10 tidak bekerja seperti di matematika: (1 <= x) menghasilkan 0 atau 1, lalu hasil itu dibandingkan dengan 10 sehingga selalu benar.", 1,
					"1 <= x <= 10",
					"x >= 1 && x <= 10",
					"x >= 1 || x <= 10",
					"x in (1, 10)"),
			},
		},

		// ------------------------------------------------------------ Lesson 2
		{
			Title:       "switch-case dan Operator Ternary",
			Description: "Percabangan banyak pilihan dan bentuk ringkas if-else.",
			Questions: []seedQuestion{
				q("Kapan switch-case paling cocok dipakai?",
					"switch cocok untuk memilih satu dari banyak kemungkinan nilai tetap (bilangan bulat atau karakter) dari sebuah ekspresi.", 1,
					"Untuk mengulang kode sebanyak n kali",
					"Untuk memilih satu dari banyak kemungkinan nilai tetap suatu ekspresi",
					"Untuk mendeklarasikan fungsi",
					"Untuk membaca input pengguna"),
				q("Apa fungsi break di dalam sebuah case pada switch?",
					"break keluar dari switch. Tanpanya eksekusi akan terus jatuh ke case berikutnya (fall-through).", 3,
					"Mengulang switch dari awal",
					"Menghentikan program",
					"Menghapus case berikutnya",
					"Keluar dari switch agar tidak lanjut ke case berikutnya"),
				q("Apa fungsi label default pada switch?",
					"default dijalankan jika tidak ada case yang cocok. Label ini opsional.", 0,
					"Dijalankan jika tidak ada case yang cocok",
					"Wajib ada di setiap switch",
					"Dijalankan pertama kali",
					"Menghentikan switch"),
				q(`Apa output kode berikut?

int x = 2;
switch (x) {
    case 1: printf("A"); break;
    case 2: printf("B"); break;
    default: printf("C");
}`,
					"x bernilai 2 sehingga case 2 dijalankan dan mencetak B, lalu break keluar dari switch.", 2,
					"A", "C", "B", "BC"),
				q(`Apa output kode berikut?

int x = 1;
switch (x) {
    case 1: printf("A");
    case 2: printf("B"); break;
    case 3: printf("C");
}`,
					"case 1 tidak punya break sehingga eksekusi jatuh ke case 2 (fall-through): tercetak A lalu B, kemudian break.", 1,
					"A", "AB", "ABC", "B"),
				q("Tipe data apa yang TIDAK bisa dipakai sebagai ekspresi switch di C?",
					"switch hanya menerima ekspresi bertipe bilangan bulat (termasuk char dan enum), bukan float atau double.", 3,
					"int", "char", "enum", "float"),
				q("Apa bentuk umum operator ternary di C?",
					"Operator ternary: kondisi ? nilai_jika_benar : nilai_jika_salah.", 2,
					"kondisi : nilai1 ? nilai2",
					"if kondisi then nilai1 else nilai2",
					"kondisi ? nilai_jika_benar : nilai_jika_salah",
					"kondisi ?? nilai"),
				q(`Apa output kode berikut?

int x = 8;
int hasil = (x > 5) ? 100 : 200;
printf("%d", hasil);`,
					"x > 5 benar, sehingga operator ternary mengambil nilai pertama, yaitu 100.", 0,
					"100", "200", "8", "1"),
				q(`Apa output kode berikut?

int a = 3, b = 9;
int maks = (a > b) ? a : b;
printf("%d", maks);`,
					"a > b salah, sehingga yang diambil adalah b, yaitu 9.", 1,
					"3", "9", "12", "0"),
				q("Apa yang terjadi jika dua case dalam satu switch memiliki nilai yang sama?",
					"Setiap label case harus unik. Nilai yang sama menyebabkan compile error (duplicate case value).", 2,
					"Case pertama saja yang dijalankan",
					"Kedua-duanya dijalankan",
					"Compile error karena nilai case duplikat",
					"Program berhenti saat dijalankan"),
			},
		},

		// ------------------------------------------------------------ Lesson 3
		{
			Title:       "Perulangan for",
			Description: "Mengulang kode dengan jumlah iterasi yang jelas.",
			Questions: []seedQuestion{
				q("Bagian apa saja yang ada di dalam tanda kurung perulangan for?",
					"Tiga bagian dipisah titik koma: inisialisasi, kondisi, dan pembaruan (update).", 0,
					"Inisialisasi; kondisi; pembaruan",
					"Tipe; nama; nilai",
					"Nama fungsi; parameter; return",
					"Awal; akhir"),
				q(`Apa output kode berikut?

for (int i = 0; i < 3; i++) printf("*");`,
					"i bernilai 0, 1, 2 sehingga printf dijalankan 3 kali dan tercetak ***.", 2,
					"*", "**", "***", "****"),
				q(`Apa output kode berikut?

for (int i = 2; i <= 10; i += 2) printf("%d ", i);`,
					"i bertambah 2 setiap iterasi dan berhenti setelah 10: 2 4 6 8 10.", 1,
					"2 4 6 8", "2 4 6 8 10", "1 3 5 7 9", "2 3 4 5 6"),
				q(`Apa output kode berikut?

for (int i = 5; i >= 1; i--) printf("%d", i);`,
					"i dimulai dari 5 dan berkurang 1 sampai 1, sehingga tercetak 54321.", 3,
					"12345", "5432", "4321", "54321"),
				q(`Apa output kode berikut?

int total = 0;
for (int i = 1; i <= 5; i++) total += i;
printf("%d", total);`,
					"total = 1 + 2 + 3 + 4 + 5 = 15.", 0,
					"15", "10", "5", "120"),
				q("Berapa kali badan perulangan ini dijalankan?\n\nfor (int i = 1; i < 10; i += 3)",
					"Nilai i: 1, 4, 7, lalu 10 (berhenti karena 10 < 10 salah). Jadi dijalankan 3 kali.", 2,
					"4", "9", "3", "10"),
				q("Apa yang terjadi jika semua bagian dikosongkan, yaitu for (;;) ?",
					"Kondisi yang kosong dianggap selalu benar, sehingga perulangan berjalan tanpa henti kecuali ada break.", 3,
					"Compile error",
					"Dijalankan satu kali",
					"Tidak dijalankan sama sekali",
					"Perulangan tak terbatas"),
				q(`Apa output kode berikut?

for (int i = 0; i < 3; i++) {
    printf("%d", i * i);
}`,
					"i * i untuk i = 0, 1, 2 menghasilkan 0, 1, 4 sehingga tercetak 014.", 1,
					"123", "014", "149", "036"),
				q("Variabel i dideklarasikan di bagian inisialisasi: for (int i = 0; ...). Di mana i dapat dipakai?",
					"Sejak C99, variabel yang dideklarasikan di bagian inisialisasi for hanya berlaku di dalam perulangan itu.", 0,
					"Hanya di dalam perulangan for itu",
					"Di seluruh program",
					"Di seluruh fungsi main",
					"Di semua file"),
				q(`Apa output kode berikut?

int n = 5, f = 1;
for (int i = 2; i <= n; i++) f *= i;
printf("%d", f);`,
					"f = 2 * 3 * 4 * 5 = 120, yaitu faktorial dari 5.", 2,
					"24", "60", "120", "720"),
			},
		},

		// ------------------------------------------------------------ Lesson 4
		{
			Title:       "while dan do-while",
			Description: "Perulangan yang berjalan berdasarkan kondisi.",
			Questions: []seedQuestion{
				q("Kapan kondisi diperiksa pada perulangan while?",
					"while memeriksa kondisi sebelum setiap iterasi. Jika kondisi sudah salah di awal, badan tidak dijalankan sama sekali.", 3,
					"Setelah badan perulangan dijalankan",
					"Tidak pernah diperiksa",
					"Hanya sekali di akhir program",
					"Sebelum setiap iterasi badan perulangan"),
				q("Apa perbedaan utama do-while dibandingkan while?",
					"do-while memeriksa kondisi setelah badan dijalankan, sehingga badan minimal dijalankan satu kali.", 1,
					"do-while tidak boleh memakai kondisi",
					"do-while menjalankan badan minimal satu kali",
					"while selalu berjalan satu kali",
					"Tidak ada bedanya"),
				q(`Apa output kode berikut?

int i = 1;
while (i <= 3) {
    printf("%d", i);
    i++;
}`,
					"i bernilai 1, 2, 3 lalu berhenti saat i = 4, sehingga tercetak 123.", 0,
					"123", "12", "321", "Tak terbatas"),
				q(`Apa output kode berikut?

int i = 10;
do {
    printf("%d", i);
    i++;
} while (i < 5);`,
					"do-while menjalankan badan sekali sebelum memeriksa kondisi. Setelah mencetak 10, kondisi i < 5 salah sehingga berhenti.", 2,
					"Tidak ada output", "Tak terbatas", "10", "1011"),
				q(`Mengapa kode berikut menjadi perulangan tak terbatas?

int i = 0;
while (i < 5) {
    printf("%d", i);
}`,
					"Nilai i tidak pernah berubah sehingga kondisi i < 5 selalu benar. Perlu ada i++ di dalam badan perulangan.", 1,
					"Kondisi memakai operator <",
					"Nilai i tidak pernah berubah sehingga kondisi i < 5 selalu benar",
					"Fungsi printf tidak boleh di dalam while",
					"Variabel i bertipe int"),
				q(`Apa output kode berikut?

int n = 12345, jumlah = 0;
while (n > 0) {
    jumlah += n % 10;
    n /= 10;
}
printf("%d", jumlah);`,
					"Kode menjumlahkan setiap digit: 5 + 4 + 3 + 2 + 1 = 15.", 3,
					"5", "12345", "10", "15"),
				q(`Apa output kode berikut?

int n = 3;
while (n--) printf("%d", n);`,
					"n-- memakai nilai n sebelum dikurangi. Iterasi: n 3 lalu 2 dicetak, n 2 lalu 1 dicetak, n 1 lalu 0 dicetak, lalu n 0 salah dan berhenti. Hasilnya 210.", 0,
					"210", "321", "3210", "10"),
				q("Manakah penulisan perulangan do-while yang benar?",
					"do-while diakhiri dengan titik koma setelah kondisi.", 2,
					`do { ... } while (kondisi)`,
					`do while (kondisi) { ... }`,
					`do { ... } while (kondisi);`,
					`while { ... } do (kondisi);`),
				q("Perulangan mana yang paling tepat untuk menampilkan menu minimal satu kali, lalu mengulanginya sampai pengguna memilih keluar?",
					"Menu harus tampil minimal sekali, dan do-while menjamin badan dijalankan sebelum kondisi diperiksa.", 3,
					"for", "while", "if", "do-while"),
				q(`Apa output kode berikut?

int x = 0;
while (x < 10) {
    x += 4;
}
printf("%d", x);`,
					"Nilai x: 0 lalu 4 lalu 8 lalu 12. Saat x = 12 kondisi x < 10 salah, sehingga tercetak 12.", 1,
					"10", "12", "8", "16"),
			},
		},

		// ------------------------------------------------------------ Lesson 5
		{
			Title:       "break, continue, dan Perulangan Bersarang",
			Description: "Mengontrol alur perulangan dan perulangan di dalam perulangan.",
			Questions: []seedQuestion{
				q("Apa efek break di dalam perulangan?",
					"break menghentikan perulangan dan program lanjut ke kode setelah perulangan.", 2,
					"Melompat ke iterasi berikutnya",
					"Mengulang iterasi saat ini",
					"Menghentikan perulangan dan keluar darinya",
					"Menghentikan seluruh program"),
				q("Apa efek continue di dalam perulangan?",
					"continue melewati sisa kode pada iterasi saat ini dan langsung lanjut ke iterasi berikutnya.", 0,
					"Melewati sisa iterasi saat ini dan lanjut ke iterasi berikutnya",
					"Keluar dari perulangan",
					"Menghentikan program",
					"Mengulang program dari awal"),
				q(`Apa output kode berikut?

for (int i = 1; i <= 5; i++) {
    if (i == 3) break;
    printf("%d", i);
}`,
					"Saat i = 3, break menghentikan perulangan. Hanya 1 dan 2 yang sempat dicetak: 12.", 3,
					"123", "1245", "12345", "12"),
				q(`Apa output kode berikut?

for (int i = 1; i <= 5; i++) {
    if (i == 3) continue;
    printf("%d", i);
}`,
					"Saat i = 3, continue melewati printf. Semua angka lain tercetak: 1245.", 1,
					"12", "1245", "12345", "45"),
				q(`Apa output kode berikut?

int total = 0;
for (int i = 1; i <= 10; i++) {
    if (i % 2 == 0) continue;
    total += i;
}
printf("%d", total);`,
					"Bilangan genap dilewati, sehingga yang dijumlahkan hanya 1 + 3 + 5 + 7 + 9 = 25.", 2,
					"55", "30", "25", "20"),
				q(`Berapa kali printf dijalankan?

for (int i = 1; i <= 2; i++)
    for (int j = 1; j <= 3; j++)
        printf("*");`,
					"Perulangan luar berjalan 2 kali dan setiap kali perulangan dalam berjalan 3 kali: 2 * 3 = 6.", 0,
					"6", "5", "2", "3"),
				q(`Pola apa yang dicetak kode berikut?

for (int i = 1; i <= 3; i++) {
    for (int j = 1; j <= i; j++) printf("*");
    printf("\n");
}`,
					"Perulangan dalam berjalan sebanyak nilai i, sehingga baris pertama 1 bintang, kedua 2 bintang, dan ketiga 3 bintang.", 1,
					"Tiga baris masing-masing 3 bintang",
					"Tiga baris berisi 1, 2, dan 3 bintang",
					"Satu baris berisi 6 bintang",
					"Tiga baris berisi 3, 2, dan 1 bintang"),
				q("Pada perulangan bersarang, apa efek break yang ada di perulangan dalam?",
					"break hanya menghentikan perulangan terdalam yang langsung membungkusnya.", 3,
					"Menghentikan semua perulangan",
					"Menghentikan program",
					"Mengulang perulangan luar dari awal",
					"Hanya menghentikan perulangan dalam tempat break berada"),
				q(`Apa output kode berikut?

int hitung = 0;
for (int i = 1; i <= 3; i++)
    for (int j = 1; j <= 4; j++)
        hitung++;
printf("%d", hitung);`,
					"hitung bertambah sebanyak 3 * 4 = 12 kali.", 0,
					"12", "7", "3", "4"),
				q(`Berapa nilai i setelah perulangan selesai?

int i = 0;
while (1) {
    if (i == 4) break;
    i++;
}`,
					"i bertambah sampai 4. Saat i == 4, break menghentikan perulangan sehingga i berakhir di 4.", 2,
					"3", "5", "4", "0"),
			},
		},
	},
}