package main

// Materi 3: Arrays & Strings (5 lesson x 10 soal = 50 soal)
var arraysStrings = seedPath{
	Name:        "Arrays & Strings",
	Description: "Mengelola kumpulan data dengan array 1D dan 2D, serta mengolah teks dengan string di C.",
	OrderIndex:  3,
	Lessons: []seedLesson{
		// ------------------------------------------------------------ Lesson 1
		{
			Title:       "Dasar Array",
			Description: "Menyimpan banyak nilai bertipe sama dalam satu variabel.",
			Questions: []seedQuestion{
				q("Apa itu array?",
					"Array adalah kumpulan elemen bertipe sama yang disimpan berurutan di memori dan diakses lewat indeks.", 1,
					"Variabel yang hanya bisa menyimpan satu nilai",
					"Kumpulan elemen bertipe sama yang disimpan berurutan di memori",
					"Fungsi untuk mengurutkan data",
					"Tipe data khusus untuk teks"),
				q("Manakah deklarasi array berisi 5 bilangan bulat yang benar?",
					"Formatnya: tipe nama[ukuran]; misalnya int nilai[5];", 3,
					"int nilai(5);", "int[5] nilai;", "array nilai[5];", "int nilai[5];"),
				q("Berapa indeks elemen pertama dalam array di C?",
					"Indeks array di C dimulai dari 0.", 0,
					"0", "1", "-1", "Ditentukan oleh programmer"),
				q("Pada array int a[6], berapa indeks valid terakhir?",
					"Array dengan 6 elemen memiliki indeks 0 sampai 5.", 2,
					"6", "7", "5", "0"),
				q(`Apa output kode berikut?

int a[4] = {10, 20, 30, 40};
printf("%d", a[2]);`,
					"a[0] = 10, a[1] = 20, a[2] = 30. Indeks dimulai dari 0.", 1,
					"20", "30", "40", "10"),
				q(`Berapa nilai a[4] pada deklarasi berikut?

int a[5] = {1, 2};`,
					"Jika hanya sebagian elemen yang diinisialisasi, elemen sisanya otomatis diisi 0.", 3,
					"Nilai sampah", "1", "2", "0"),
				q(`Apa output kode berikut?

int a[3] = {5, 6, 7};
a[1] = 9;
printf("%d", a[1] + a[2]);`,
					"a[1] diubah menjadi 9, sehingga a[1] + a[2] = 9 + 7 = 16.", 0,
					"16", "13", "12", "18"),
				q("Jika int berukuran 4 byte, berapa hasil sizeof(a) untuk int a[5];?",
					"Ukuran total array = jumlah elemen * ukuran tiap elemen = 5 * 4 = 20 byte.", 2,
					"5", "4", "20", "9"),
				q("Apa risiko mengakses a[10] pada array yang hanya punya 5 elemen?",
					"C tidak memeriksa batas array, sehingga akses di luar batas menyebabkan perilaku tak terdefinisi (undefined behavior): bisa crash atau membaca data acak.", 1,
					"Compiler selalu menolaknya",
					"Perilaku tak terdefinisi (undefined behavior)",
					"Array otomatis bertambah ukurannya",
					"Selalu menghasilkan 0"),
				q(`Apa output kode berikut?

int a[] = {3, 6, 9, 12};
printf("%d", sizeof(a) / sizeof(a[0]));`,
					"sizeof(a) / sizeof(a[0]) menghitung jumlah elemen: total byte dibagi ukuran satu elemen. Hasilnya 4.", 0,
					"4", "12", "16", "3"),
			},
		},

		// ------------------------------------------------------------ Lesson 2
		{
			Title:       "Array dan Perulangan",
			Description: "Menelusuri array: jumlah, nilai maksimum, pencarian, dan pembalikan.",
			Questions: []seedQuestion{
				q("Perulangan apa yang paling umum dipakai untuk menelusuri semua elemen array?",
					"for cocok karena jumlah elemen diketahui dan indeks bertambah teratur.", 2,
					"if", "switch", "for", "goto"),
				q("Untuk menelusuri array berukuran n, kondisi perulangan yang benar adalah...",
					"Indeks valid adalah 0 sampai n - 1, sehingga kondisinya i < n. Memakai i <= n menyebabkan akses di luar batas.", 0,
					"i < n", "i <= n", "i < n + 1", "i != n - 1"),
				q(`Apa output kode berikut?

int a[4] = {2, 4, 6, 8};
int s = 0;
for (int i = 0; i < 4; i++) s += a[i];
printf("%d", s);`,
					"s = 2 + 4 + 6 + 8 = 20.", 3,
					"14", "18", "24", "20"),
				q(`Apa output kode berikut?

int a[5] = {4, 9, 2, 7, 5};
int maks = a[0];
for (int i = 1; i < 5; i++)
    if (a[i] > maks) maks = a[i];
printf("%d", maks);`,
					"Kode mencari nilai terbesar dalam array. Nilai terbesarnya adalah 9.", 1,
					"4", "9", "7", "5"),
				q(`Apa output kode berikut?

int a[5] = {1, 2, 3, 4, 5};
for (int i = 4; i >= 0; i--) printf("%d", a[i]);`,
					"Perulangan menelusuri array dari indeks terakhir ke awal, sehingga tercetak 54321.", 2,
					"12345", "5432", "54321", "43210"),
				q(`Berapa nilai hitung di akhir?

int a[5] = {5, 3, 8, 3, 1};
int hitung = 0;
for (int i = 0; i < 5; i++)
    if (a[i] == 3) hitung++;`,
					"Angka 3 muncul di indeks 1 dan 3, sehingga hitung = 2.", 0,
					"2", "1", "3", "5"),
				q(`Apa output kode berikut?

int a[6] = {1, 2, 3, 4, 5, 6};
int genap = 0;
for (int i = 0; i < 6; i++)
    if (a[i] % 2 == 0) genap++;
printf("%d", genap);`,
					"Bilangan genap dalam array adalah 2, 4, dan 6, sehingga genap = 3.", 3,
					"6", "2", "4", "3"),
				q("Kode mana yang mengisi array a[5] dengan nilai 0, 1, 2, 3, 4?",
					"Perulangan harus berjalan dari 0 sampai 4 (i < 5) dan mengisi a[i] dengan i.", 1,
					`for (int i = 0; i <= 5; i++) a[i] = i;`,
					`for (int i = 0; i < 5; i++) a[i] = i;`,
					`for (int i = 1; i < 5; i++) a[i] = i;`,
					`for (int i = 0; i < 5; i++) a = i;`),
				q(`Apa output kode berikut?

int a[2] = {1, 2};
int t = a[0];
a[0] = a[1];
a[1] = t;
printf("%d%d", a[0], a[1]);`,
					"Kode menukar isi a[0] dan a[1] dengan variabel bantu t, sehingga hasilnya 21.", 2,
					"12", "11", "21", "22"),
				q(`Apa output kode berikut?

int a[3] = {1, 2, 3};
for (int i = 0; i < 3; i++) a[i] *= 2;
printf("%d", a[2]);`,
					"Setiap elemen dikali 2, sehingga a[2] menjadi 3 * 2 = 6.", 0,
					"6", "3", "4", "2"),
			},
		},

		// ------------------------------------------------------------ Lesson 3
		{
			Title:       "Array Dua Dimensi",
			Description: "Matriks: baris, kolom, dan perulangan bersarang.",
			Questions: []seedQuestion{
				q("Manakah deklarasi array 2D dengan 3 baris dan 4 kolom?",
					"Formatnya: tipe nama[baris][kolom];", 1,
					"int m[4][3];", "int m[3][4];", "int m[3,4];", "int m(3)(4);"),
				q(`Apa output kode berikut?

int m[2][2] = {{1, 2}, {3, 4}};
printf("%d", m[0][1]);`,
					"m[0][1] adalah baris pertama, kolom kedua, yaitu 2.", 0,
					"2", "1", "3", "4"),
				q(`Apa output kode berikut?

int m[2][3] = {{1, 2, 3}, {4, 5, 6}};
printf("%d", m[1][2]);`,
					"m[1][2] adalah baris kedua, kolom ketiga, yaitu 6.", 2,
					"3", "5", "6", "4"),
				q("Berapa jumlah elemen pada int m[3][4];?",
					"Jumlah elemen = baris * kolom = 3 * 4 = 12.", 3,
					"7", "9", "10", "12"),
				q("Perulangan apa yang biasanya dipakai untuk menelusuri semua elemen array 2D?",
					"Biasanya dipakai dua perulangan bersarang: satu untuk baris dan satu untuk kolom.", 1,
					"Satu if saja",
					"Dua perulangan bersarang: satu untuk baris, satu untuk kolom",
					"Hanya do-while",
					"Array 2D tidak bisa ditelusuri"),
				q(`Apa output kode berikut?

int m[2][2] = {{1, 2}, {3, 4}};
int s = 0;
for (int i = 0; i < 2; i++)
    for (int j = 0; j < 2; j++)
        s += m[i][j];
printf("%d", s);`,
					"Semua elemen dijumlahkan: 1 + 2 + 3 + 4 = 10.", 2,
					"4", "7", "10", "6"),
				q(`Apa output kode berikut?

int m[3][3] = {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}};
int s = 0;
for (int i = 0; i < 3; i++) s += m[i][i];
printf("%d", s);`,
					"m[i][i] adalah diagonal utama: 1 + 5 + 9 = 15.", 0,
					"15", "18", "12", "45"),
				q("Pada m[i][j], apa arti i dan j?",
					"Indeks pertama (i) menunjukkan baris dan indeks kedua (j) menunjukkan kolom.", 3,
					"i adalah kolom dan j adalah baris",
					"i dan j adalah nilai elemen",
					"i adalah jumlah baris dan j adalah jumlah kolom",
					"i adalah indeks baris dan j adalah indeks kolom"),
				q(`Apa output kode berikut?

int m[2][3] = {{1, 2, 3}, {4, 5, 6}};
printf("%d", m[1][0] + m[0][2]);`,
					"m[1][0] = 4 dan m[0][2] = 3, sehingga hasilnya 7.", 1,
					"5", "7", "9", "3"),
				q(`Berapa nilai m[1][1]?

int m[2][2] = {{1}, {2}};`,
					"Elemen yang tidak ditulis diisi 0 ketika sebagian elemen diinisialisasi, sehingga m[1][1] = 0.", 2,
					"2", "1", "0", "Nilai sampah"),
			},
		},

		// ------------------------------------------------------------ Lesson 4
		{
			Title:       "Dasar String",
			Description: "String sebagai array karakter yang diakhiri karakter null.",
			Questions: []seedQuestion{
				q("Apa sebenarnya string di C?",
					"C tidak punya tipe string bawaan. String adalah array char yang diakhiri karakter null.", 3,
					"Tipe data bawaan bernama string",
					"Kumpulan bilangan bulat",
					"Tipe data khusus berbentuk objek",
					"Array char yang diakhiri karakter null"),
				q("Karakter apa yang menandai akhir sebuah string di C?",
					`Karakter null (\0) menandai akhir string sehingga fungsi seperti printf tahu kapan harus berhenti.`, 1,
					`'\n'`, `'\0'`, `' '`, `'.'`),
				q(`Berapa ukuran array s (dalam elemen) pada deklarasi berikut?

char s[] = "Halo";`,
					"Empat huruf ditambah satu karakter null di akhir menjadi 5 elemen.", 0,
					"5", "4", "3", "6"),
				q("Manakah cara yang benar mendeklarasikan string berisi Halo?",
					`String disimpan di array char dan ditulis dengan kutip ganda: char s[] = "Halo";`, 2,
					`char s = "Halo";`,
					`string s = "Halo";`,
					`char s[] = "Halo";`,
					`char s[] = 'Halo';`),
				q(`Apa output kode berikut?

char s[] = "Dunia";
printf("%c", s[1]);`,
					"Indeks dimulai dari 0: s[0] = D, s[1] = u.", 1,
					"D", "u", "n", "i"),
				q("Format specifier apa yang dipakai untuk mencetak string dengan printf?",
					"%s dipakai untuk string, sedangkan %c untuk satu karakter.", 3,
					"%d", "%c", "%f", "%s"),
				q(`Apa output kode berikut?

char s[] = "ABC";
s[1] = 'X';
printf("%s", s);`,
					"Elemen array char dapat diubah, sehingga s[1] diganti X dan hasilnya AXC.", 0,
					"AXC", "ABC", "XBC", "ABX"),
				q(`Mengapa pada scanf("%s", nama); tidak perlu memakai tanda & di depan nama?`,
					"Nama array sudah berperilaku sebagai alamat elemen pertamanya, sehingga & tidak diperlukan.", 2,
					"Karena nama adalah konstanta",
					"Karena scanf dengan %s tidak menyimpan apa-apa",
					"Karena nama array sudah menjadi alamat elemen pertamanya",
					"Karena & hanya untuk bilangan pecahan"),
				q(`Apa kelemahan scanf("%s", ...) saat membaca nama lengkap seperti Budi Santoso?`,
					"%s berhenti membaca ketika menemukan spasi, sehingga hanya kata Budi yang tersimpan.", 3,
					"Hanya bisa membaca angka",
					"Selalu membaca lebih dari 100 karakter",
					"Menghapus huruf kapital",
					"Berhenti membaca saat menemukan spasi"),
				q("Fungsi apa yang lebih aman untuk membaca satu baris teks (termasuk spasi) ke dalam array?",
					"fgets membatasi jumlah karakter yang dibaca sehingga tidak melewati ukuran array, dan bisa membaca spasi. Fungsi gets() berbahaya dan sudah dihapus dari standar C.", 1,
					`gets()`, `fgets()`, `scanf("%s")`, `getchar() saja`),
			},
		},

		// ------------------------------------------------------------ Lesson 5
		{
			Title:       "Fungsi String (string.h)",
			Description: "strlen, strcpy, strcat, strcmp, dan toupper.",
			Questions: []seedQuestion{
				q("Fungsi apa yang menghitung panjang string (tanpa karakter null)?",
					"strlen() mengembalikan jumlah karakter sebelum karakter null.", 0,
					"strlen()", "strcpy()", "strcmp()", "strcat()"),
				q(`Apa output kode berikut?

printf("%zu", strlen("Halo"));`,
					"Kata Halo terdiri dari 4 karakter, dan strlen tidak menghitung karakter null.", 2,
					"5", "3", "4", "6"),
				q("Fungsi apa yang menyalin isi satu string ke string lain?",
					"strcpy(tujuan, sumber) menyalin isi string sumber ke string tujuan.", 1,
					"strcat()", "strcpy()", "strcmp()", "strlen()"),
				q("Fungsi apa yang menyambung (menggabungkan) dua string?",
					"strcat(tujuan, sumber) menambahkan string sumber di akhir string tujuan.", 3,
					"strcpy()", "strcmp()", "strlen()", "strcat()"),
				q("Apa nilai yang dikembalikan strcmp(a, b) jika kedua string sama persis?",
					"strcmp mengembalikan 0 jika kedua string sama, nilai negatif jika a lebih kecil, dan positif jika a lebih besar.", 2,
					"1", "-1", "0", "NULL"),
				q("Mengapa a == b tidak tepat untuk membandingkan isi dua string?",
					"Operator == pada array membandingkan alamat memori, bukan isi karakternya. Gunakan strcmp() untuk membandingkan isi.", 0,
					"Karena membandingkan alamat memori, bukan isi string",
					"Karena == hanya untuk bilangan pecahan",
					"Karena == selalu bernilai 1",
					"Karena == tidak ada di C"),
				q(`Apa output kode berikut?

char s[20] = "Hai ";
strcat(s, "Dunia");
printf("%s", s);`,
					"strcat menambahkan Dunia di akhir Hai , sehingga hasilnya Hai Dunia.", 1,
					"Dunia", "Hai Dunia", "HaiDunia", "Hai"),
				q("Header apa yang harus di-include agar strlen, strcpy, dan strcat bisa dipakai?",
					"Fungsi-fungsi string standar dideklarasikan di <string.h>.", 3,
					"<stdio.h>", "<stdlib.h>", "<math.h>", "<string.h>"),
				q(`Apa output kode berikut?

char a[10] = "abc";
char b[10];
strcpy(b, a);
b[0] = 'X';
printf("%s %s", a, b);`,
					"strcpy membuat salinan terpisah. Mengubah b tidak memengaruhi a, sehingga hasilnya abc Xbc.", 0,
					"abc Xbc", "Xbc Xbc", "abc abc", "Xbc abc"),
				q("Fungsi apa dari <ctype.h> yang mengubah huruf kecil menjadi huruf kapital?",
					"toupper() mengubah satu karakter huruf kecil menjadi huruf kapital.", 2,
					"upper()", "strupr()", "toupper()", "capital()"),
			},
		},
	},
}