# Panduan tim: database Supabase bersama

Semua developer menjalankan backend lokal masing-masing, tetapi terhubung ke
**satu database Supabase yang sama**. Perubahan data langsung terlihat oleh
semua orang, tanpa perlu menarik data.

## Koneksi

-   Pakai connection string **Session pooler** (port 5432). Koneksi langsung
    (`db.<ref>.supabase.co`) hanya IPv6 kecuali project punya add-on IPv4.
-   Simpan di `backend/.env` (tidak di-commit). Kirim kredensial lewat kanal
    aman, bukan lewat repository atau grup publik. Ganti password database
    bila ada anggota yang keluar.
-   Pool koneksi dibuat kecil (default 5) karena banyak developer berbagi
    batas koneksi yang sama. Jangan menaikkannya tanpa alasan.
-   `APP_TIMEZONE` harus sama untuk semua developer (default `Asia/Jakarta`),
    karena menentukan batas hari/pekan misi, quest, streak, dan leaderboard.

## Aturan yang melindungi data bersama

1.  **Migrasi skema (`go run ./cmd/migrate`)** dijalankan oleh satu orang yang
    ditunjuk, setelah perubahan model di-merge. Jangan menjalankannya dari
    branch fitur yang belum di-merge: skema berubah untuk semua orang
    sementara kode developer lain belum siap.
2.  Perubahan skema harus **kompatibel mundur** (tambah kolom/tabel, jangan
    hapus atau ganti nama) agar backend developer lain tidak rusak.
3.  **Jangan menjalankan test yang menulis data ke database bersama.** Test
    integrasi memakai PostgreSQL lokal terpisah.
4.  Jangan menjalankan skrip yang menghapus atau mereset data di database
    bersama tanpa kesepakatan tim dan backup lebih dulu.
5.  Akun uji memakai prefix `dev_` pada username agar mudah dikenali dan
    dibersihkan. Ingat: data uji ikut tampil di leaderboard semua orang.
6.  Jangan commit `.env`. Jika secret pernah masuk repository, anggap bocor
    dan ganti.

## Backup

Sebelum migrasi skema atau operasi berisiko, ambil dump:

``` powershell
pg_dump "<connection-string>" --no-owner --format=custom --file=codelingo-YYYYMMDD.dump
```

Simpan di luar repository. Pulihkan dengan `pg_restore`. Jangan hanya
mengandalkan backup otomatis Supabase; periksa dulu apa yang disediakan paket
project Anda.

## Keamanan

-   `go run ./cmd/migrate` mengaktifkan Row Level Security (RLS) pada semua
    tabel tanpa policy, sehingga tabel tidak terbuka lewat Data API Supabase
    (anon key). Backend Go memakai role `postgres` yang melewati RLS.
-   Jika memakai role database lain yang tidak melewati RLS, query backend
    akan mengembalikan data kosong.

## Akun admin pertama

``` powershell
$env:ADMIN_USERNAME = "admin"
$env:ADMIN_EMAIL = "admin@example.com"
$env:ADMIN_PASSWORD = "minimal-8-karakter"
go run ./cmd/create-admin
```

Perintah ini idempoten dan tidak pernah menimpa password akun yang sudah ada.
Admin berikutnya dipromosikan lewat Admin Dashboard.
