# Parking Lot

## Deskripsi Proyek
Proyek ini adalah sistem otomatis untuk mengelola parkir mobil di sebuah tempat parkir. Sistem ini memungkinkan pengguna untuk membuat tempat parkir, memarkir mobil, mengeluarkan mobil dari tempat parkir, dan melihat status tempat parkir saat ini.

## Fitur
- Membuat tempat parkir dengan jumlah slot yang ditentukan.
- Memarkir mobil di slot yang tersedia terdekat.
- Mengeluarkan mobil dari slot parkir dan menghitung biaya parkir berdasarkan waktu parkir.
- Menampilkan status saat ini dari semua slot parkir.

## Persyaratan
- Go 1.16 atau lebih baru
- Sistem operasi Unix (Linux, macOS)

## Struktur Direktori
parking_lot/
│
├── bin/
│   ├── setup
│   └── parking_lot
│
├── cmd/
│   └── parking_lot/
│       └── main.go
│
├── functional_spec/
│   └── README.md
│
├── internal/
│   ├── app/
│   │   └── parking_lot.go
│   ├── domain/
│   │   └── models.go
│   ├── infra/
│   │   └── repository.go
│   └── usecases/
│       ├── park.go
│       ├── leave.go
│       ├── status.go
│       └── create_parking_lot.go
│
├── test/
│   └── parking_lot_test.go
│
└── go.mod


## Cara Menjalankan Proyek
- Setup
1. Clone repository ini atau ekstrak zip yang diberikan.
2. Navigasikan ke direktori proyek `parking_lot`.
3. Jalankan perintah berikut di terminal untuk mengatur proyek dan menjalankan pengujian:
   $ bin/setup

- Menjalankan Aplikasi
Untuk menjalankan aplikasi, gunakan perintah berikut:
   $ bin/parking_lot <nama_file_input>

Contoh:
   $ bin/parking_lot file_inputs.txt

- Format Perintah
Aplikasi ini menerima perintah berikut melalui file input:
1. Membuat tempat parkir:
    create_parking_lot {capacity}

2. Memarkir mobil:
    park {registration_number} {color}

3. Mengeluarkan mobil:
    leave {slot_number} {hours}

4. Menampilkan status:
    status
