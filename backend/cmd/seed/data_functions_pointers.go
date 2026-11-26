package main

// Materi 4: Functions & Pointers (5 lesson x 10 soal = 50 soal)
var functionsPointers = seedPath{
	Name:        "Functions & Pointers",
	Description: "Membuat fungsi, memahami scope dan rekursi, lalu menguasai pointer dan memori dinamis di C.",
	OrderIndex:  4,
	Lessons: []seedLesson{
		// ------------------------------------------------------------ Lesson 1
		{
			Title:       "Dasar Fungsi",
			Description: "Membuat, memanggil, dan mengembalikan nilai dari fungsi.",
			Questions: []seedQuestion{
				q("Apa keuntungan utama memakai fungsi?",
					"Fungsi mengelompokkan kode agar bisa dipakai ulang sehingga program lebih terstruktur dan mudah dirawat.", 0,
					"Kode bisa dipakai ulang dan program lebih terstruktur",
					"Program otomatis berjalan lebih cepat",
					"Variabel tidak perlu dideklarasikan",
					"Program tidak akan pernah error"),
				q("Tipe kembalian apa yang dipakai untuk fungsi yang tidak mengembalikan nilai?",
					"void berarti fungsi tidak mengembalikan nilai apa pun.", 2,
					"int", "null", "void", "none"),
				q(`Apa output kode berikut?

int tambah(int a, int b) { return a + b; }

printf("%d", tambah(4, 5));`,
					"tambah(4, 5) mengembalikan 4 + 5 = 9.", 1,
					"45", "9", "20", "Error"),
				q("Apa fungsi return di dalam sebuah fungsi?",
					"return mengembalikan nilai ke pemanggil dan sekaligus mengakhiri eksekusi fungsi.", 3,
					"Mencetak nilai ke layar",
					"Mengulang fungsi",
					"Menghapus variabel",
					"Mengembalikan nilai ke pemanggil dan mengakhiri fungsi"),
				q("Apa itu prototipe fungsi?",
					"Prototipe adalah deklarasi fungsi (tipe kembalian, nama, dan parameter) yang ditulis sebelum definisinya agar compiler mengenal fungsi tersebut.", 0,
					"Deklarasi fungsi (tipe kembalian, nama, parameter) sebelum definisinya",
					"Isi lengkap sebuah fungsi",
					"Fungsi yang dipanggil otomatis",
					"Nama lain dari main()"),
				q("Pada int kali(int x, int y), apa sebutan untuk x dan y?",
					"x dan y adalah parameter (pada definisi fungsi). Nilai yang dikirim saat pemanggilan disebut argumen.", 2,
					"Argumen", "Nilai kembalian", "Parameter", "Konstanta"),
				q(`Apa output kode berikut?

void sapa() { printf("Hai"); }

sapa();
sapa();`,
					"Fungsi sapa dipanggil dua kali, sehingga Hai tercetak dua kali: HaiHai.", 3,
					"Hai", "Error", "Tidak ada output", "HaiHai"),
				q(`Apa output kode berikut?

int maks(int a, int b) { return (a > b) ? a : b; }

printf("%d", maks(3, 8));`,
					"maks mengembalikan nilai yang lebih besar, yaitu 8.", 1,
					"3", "8", "11", "0"),
				q("Apa yang terjadi jika sebuah fungsi dipanggil sebelum dideklarasikan atau didefinisikan (tanpa prototipe)?",
					"Compiler belum mengenal fungsi tersebut sehingga melaporkan error atau peringatan. Solusinya: tulis prototipe di bagian atas.", 0,
					"Compiler melaporkan error atau peringatan karena fungsi belum dikenal",
					"Program berjalan normal",
					"Fungsi dibuat otomatis",
					"Program mencetak 0"),
				q(`Apa output kode berikut?

int dua(int x) { return x * 2; }

printf("%d", dua(dua(3)));`,
					"dua(3) = 6, lalu dua(6) = 12.", 2,
					"6", "9", "12", "18"),
			},
		},

		// ------------------------------------------------------------ Lesson 2
		{
			Title:       "Scope, Parameter, dan Rekursi",
			Description: "Jangkauan variabel, pass by value, dan fungsi rekursif.",
			Questions: []seedQuestion{
				q("Di mana variabel lokal dapat diakses?",
					"Variabel lokal hanya dikenal di dalam fungsi atau blok tempat ia dideklarasikan.", 1,
					"Di seluruh program",
					"Hanya di dalam fungsi atau blok tempat ia dideklarasikan",
					"Di semua file",
					"Hanya di main()"),
				q("Di mana variabel global dideklarasikan?",
					"Variabel global dideklarasikan di luar semua fungsi sehingga bisa diakses oleh fungsi-fungsi di bawahnya.", 3,
					"Di dalam main()",
					"Di dalam if",
					"Di dalam for",
					"Di luar semua fungsi"),
				q(`Apa output kode berikut?

void ubah(int x) { x = 99; }

int a = 5;
ubah(a);
printf("%d", a);`,
					"C memakai pass by value: fungsi menerima salinan nilai. Mengubah x di dalam fungsi tidak memengaruhi a.", 0,
					"5", "99", "0", "Error"),
				q("Apa itu rekursi?",
					"Rekursi adalah teknik ketika sebuah fungsi memanggil dirinya sendiri.", 2,
					"Fungsi tanpa parameter",
					"Fungsi yang hanya dipanggil dari main()",
					"Fungsi yang memanggil dirinya sendiri",
					"Fungsi yang tidak punya return"),
				q("Apa yang wajib ada pada fungsi rekursif agar tidak berjalan selamanya?",
					"Fungsi rekursif harus punya kondisi dasar (base case) yang menghentikan pemanggilan berikutnya.", 1,
					"Variabel global",
					"Kondisi dasar (base case) yang menghentikan pemanggilan",
					"Kata kunci perulangan",
					"Parameter bertipe float"),
				q(`Apa output kode berikut?

int fakt(int n) {
    if (n <= 1) return 1;
    return n * fakt(n - 1);
}

printf("%d", fakt(4));`,
					"fakt(4) = 4 * 3 * 2 * 1 = 24.", 0,
					"24", "10", "16", "4"),
				q(`Apa output kode berikut?

int jumlah(int n) {
    if (n == 0) return 0;
    return n + jumlah(n - 1);
}

printf("%d", jumlah(4));`,
					"jumlah(4) = 4 + 3 + 2 + 1 + 0 = 10.", 3,
					"24", "4", "6", "10"),
				q(`Apa output kode berikut?

void hitung() {
    static int c = 0;
    c++;
    printf("%d", c);
}

hitung();
hitung();
hitung();`,
					"Variabel static mempertahankan nilainya antar pemanggilan dan tidak di-reset, sehingga tercetak 123.", 2,
					"111", "321", "123", "000"),
				q(`Apa output kode berikut?

int x = 10;

int main() {
    int x = 20;
    printf("%d", x);
    return 0;
}`,
					"Variabel lokal x di dalam main() menyembunyikan (shadow) variabel global x, sehingga yang tercetak 20.", 1,
					"10", "20", "30", "Error"),
				q("Apa yang terjadi jika fungsi rekursif tidak punya base case?",
					"Pemanggilan berlanjut terus sampai stack penuh (stack overflow) dan program crash.", 3,
					"Berhenti otomatis setelah 100 kali",
					"Hasilnya selalu 0",
					"Compiler menambahkan base case",
					"Stack overflow dan program crash"),
			},
		},

		// ------------------------------------------------------------ Lesson 3
		{
			Title:       "Dasar Pointer",
			Description: "Alamat memori, operator & dan *, serta NULL.",
			Questions: []seedQuestion{
				q("Apa itu pointer?",
					"Pointer adalah variabel yang menyimpan alamat memori dari variabel lain.", 0,
					"Variabel yang menyimpan alamat memori",
					"Variabel yang menyimpan teks",
					"Fungsi untuk mencetak alamat",
					"Tipe data pecahan"),
				q("Operator apa yang dipakai untuk mengambil alamat sebuah variabel?",
					"Operator & (address-of) menghasilkan alamat memori variabel, misalnya &x.", 2,
					"*", "#", "&", "@"),
				q("Manakah deklarasi pointer ke int yang benar?",
					"Tanda * pada deklarasi menandakan pointer: int *p; berarti p menunjuk ke int.", 1,
					"int &p;", "int *p;", "pointer int p;", "int p*;"),
				q("Apa fungsi operator * pada ekspresi *p (dereference)?",
					"*p mengakses nilai yang tersimpan di alamat yang ditunjuk oleh p.", 3,
					"Mengambil alamat p",
					"Mengalikan p dengan dirinya sendiri",
					"Menghapus p",
					"Mengakses nilai yang tersimpan di alamat yang ditunjuk p"),
				q(`Apa output kode berikut?

int x = 10;
int *p = &x;
*p = 20;
printf("%d", x);`,
					"p menunjuk ke x, sehingga *p = 20 mengubah nilai x menjadi 20.", 2,
					"10", "Alamat x", "20", "Error"),
				q("Apa arti pointer bernilai NULL?",
					"NULL berarti pointer tidak menunjuk ke alamat memori yang valid.", 0,
					"Tidak menunjuk ke alamat memori yang valid",
					"Menunjuk ke angka 0",
					"Menunjuk ke elemen pertama array",
					"Menunjuk ke fungsi main()"),
				q(`Apa output kode berikut?

int a = 4;
int *p = &a;
printf("%d", *p + 1);`,
					"*p bernilai 4 (isi a), sehingga *p + 1 = 5.", 3,
					"4", "Alamat a ditambah 1", "6", "5"),
				q("Apa yang terjadi jika kita melakukan dereference pada pointer NULL (*p saat p == NULL)?",
					"Mengakses alamat NULL adalah perilaku tak terdefinisi dan biasanya menyebabkan program crash (segmentation fault).", 1,
					"Menghasilkan nilai 0",
					"Biasanya program crash (segmentation fault)",
					"Menghasilkan nilai -1",
					"Compiler otomatis mengisi nilainya"),
				q(`Apa output kode berikut?

int x = 5, y = 8;
int *p = &x;
p = &y;
*p = 1;
printf("%d %d", x, y);`,
					"p dipindahkan menunjuk ke y, sehingga *p = 1 mengubah y. Nilai x tetap 5, hasilnya 5 1.", 2,
					"1 8", "1 1", "5 1", "5 8"),
				q("Format specifier apa yang biasa dipakai untuk mencetak alamat pointer?",
					`%p dipakai untuk mencetak nilai pointer (alamat), contohnya printf("%p", (void *)&x);`, 0,
					"%p", "%s", "%c", "%f"),
			},
		},

		// ------------------------------------------------------------ Lesson 4
		{
			Title:       "Pointer, Array, dan Fungsi",
			Description: "Aritmatika pointer, array sebagai pointer, dan mengubah variabel lewat fungsi.",
			Questions: []seedQuestion{
				q("Pada kebanyakan ekspresi, nama array a berperilaku sebagai apa?",
					"Nama array umumnya meluruh (decay) menjadi pointer ke elemen pertamanya.", 2,
					"Jumlah elemen array",
					"Nilai elemen terakhir",
					"Alamat elemen pertama",
					"Ukuran array dalam byte"),
				q(`Apa output kode berikut?

int a[3] = {1, 2, 3};
int *p = a;
printf("%d", *(p + 1));`,
					"p menunjuk ke a[0]. p + 1 menunjuk ke a[1], sehingga *(p + 1) = 2.", 0,
					"2", "1", "3", "Alamat a[1]"),
				q(`Apa output kode berikut?

int a[4] = {10, 20, 30, 40};
int *p = a;
p += 2;
printf("%d", *p);`,
					"p += 2 menggeser pointer dua elemen ke a[2], yang bernilai 30.", 3,
					"10", "20", "40", "30"),
				q("Agar sebuah fungsi dapat mengubah nilai variabel milik pemanggilnya, apa yang harus dikirim?",
					"Karena C memakai pass by value, kirimkan alamat variabel (pointer) agar fungsi bisa mengubah nilai aslinya.", 1,
					"Salinan nilainya",
					"Pointer ke variabel tersebut",
					"Namanya sebagai string",
					"Tipe datanya"),
				q(`Apa output kode berikut?

void tukar(int *a, int *b) {
    int t = *a;
    *a = *b;
    *b = t;
}

int x = 1, y = 2;
tukar(&x, &y);
printf("%d %d", x, y);`,
					"Fungsi menukar nilai x dan y lewat pointer, sehingga hasilnya 2 1.", 2,
					"1 2", "1 1", "2 1", "2 2"),
				q("Bagaimana array dikirim ke fungsi di C?",
					"Yang dikirim adalah alamat elemen pertama (pointer), bukan salinan seluruh array. Perubahan di dalam fungsi memengaruhi array asli.", 1,
					"Selalu disalin seluruhnya",
					"Sebagai alamat elemen pertamanya (pointer)",
					"Sebagai string",
					"Tidak bisa dikirim ke fungsi"),
				q(`Apa output kode berikut?

void isi(int a[], int n) {
    for (int i = 0; i < n; i++) a[i] = 0;
}

int d[3] = {5, 6, 7};
isi(d, 3);
printf("%d", d[1]);`,
					"Array dikirim sebagai pointer, sehingga isi() mengubah array asli menjadi 0 semua.", 0,
					"0", "6", "5", "7"),
				q("Mengapa fungsi yang menerima array biasanya juga menerima parameter ukuran (n)?",
					"Array yang dikirim menjadi pointer, sehingga ukuran aslinya tidak diketahui di dalam fungsi (sizeof hanya memberi ukuran pointer).", 3,
					"Karena array selalu berukuran 10",
					"Karena n wajib ada di semua fungsi",
					"Karena C menolak array tanpa n",
					"Karena fungsi tidak bisa mengetahui ukuran array dari pointer"),
				q(`Apa output kode berikut?

int a[3] = {4, 5, 6};
int *p = &a[2];
printf("%d", *(p - 1));`,
					"p menunjuk ke a[2]. p - 1 menunjuk ke a[1], yang bernilai 5.", 2,
					"4", "6", "5", "Alamat a[1]"),
				q(`Apa output kode berikut?

int a[3] = {7, 8, 9};
int *p = a;
p++;
printf("%d", *p);`,
					"p++ menggeser pointer ke elemen berikutnya (a[1]), yang bernilai 8.", 0,
					"8", "7", "9", "Alamat a[1]"),
			},
		},

		// ------------------------------------------------------------ Lesson 5
		{
			Title:       "Memori Dinamis",
			Description: "malloc, calloc, free, dan jebakan umum pengelolaan memori.",
			Questions: []seedQuestion{
				q("Header apa yang diperlukan untuk memakai malloc dan free?",
					"malloc, calloc, realloc, dan free dideklarasikan di <stdlib.h>.", 1,
					"<stdio.h>", "<stdlib.h>", "<string.h>", "<math.h>"),
				q("Apa yang dilakukan malloc(n)?",
					"malloc mengalokasikan n byte memori di heap dan mengembalikan alamat awalnya.", 3,
					"Menghapus n byte memori",
					"Menyalin n byte data",
					"Mengubah tipe variabel",
					"Mengalokasikan n byte memori di heap dan mengembalikan alamat awalnya"),
				q("Apa yang dikembalikan malloc jika alokasi memori gagal?",
					"malloc mengembalikan NULL jika memori tidak cukup, sehingga hasilnya sebaiknya diperiksa sebelum dipakai.", 0,
					"NULL",
					"Memori sebesar 0 byte",
					"-1",
					"Program berhenti otomatis"),
				q("Manakah cara yang benar mengalokasikan memori untuk 5 bilangan bulat int?",
					"Kalikan jumlah elemen dengan ukuran tiap elemen: 5 * sizeof(int), karena ukuran int bisa berbeda antar sistem.", 2,
					`malloc(5)`,
					`malloc(int * 5)`,
					`malloc(5 * sizeof(int))`,
					`malloc(sizeof(5))`),
				q("Fungsi apa yang melepaskan memori hasil malloc?",
					"free() mengembalikan memori yang dialokasikan dinamis ke sistem.", 1,
					"delete()", "free()", "release()", "clear()"),
				q("Apa itu memory leak?",
					"Memory leak terjadi ketika memori dialokasikan tetapi tidak pernah di-free sehingga terbuang selama program berjalan.", 0,
					"Memori yang sudah dialokasikan tetapi tidak pernah di-free sehingga terbuang",
					"Program yang crash saat dijalankan",
					"Array yang terlalu besar",
					"Pointer yang bernilai NULL"),
				q("Apa perbedaan calloc dengan malloc?",
					"calloc mengalokasikan memori sekaligus mengisinya dengan nol, sedangkan malloc membiarkan isinya acak (belum diinisialisasi).", 3,
					"calloc tidak bisa dipakai untuk int",
					"calloc hanya untuk string",
					"calloc melepaskan memori",
					"calloc mengisi memori dengan nol, malloc tidak"),
				q("Apa itu dangling pointer?",
					"Dangling pointer masih menyimpan alamat memori yang sudah di-free, sehingga mengaksesnya berbahaya.", 2,
					"Pointer bernilai NULL",
					"Pointer ke fungsi main",
					"Pointer yang masih menyimpan alamat memori yang sudah di-free",
					"Pointer bertipe void"),
				q(`Apa output kode berikut?

int *p = malloc(3 * sizeof(int));
p[0] = 10;
p[1] = 20;
p[2] = 30;
printf("%d", p[1]);
free(p);`,
					"Memori hasil malloc dipakai seperti array: p[1] bernilai 20.", 1,
					"10", "20", "30", "Error"),
				q("Apa praktik yang baik setelah memanggil free(p)?",
					"Mengatur p = NULL mencegah dangling pointer terpakai secara tidak sengaja.", 3,
					"Memanggil free(p) sekali lagi",
					"Memakai *p untuk memastikan masih terisi",
					"Membiarkan p apa adanya",
					"Mengatur p = NULL"),
			},
		},
	},
}