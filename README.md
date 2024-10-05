# Bank Backend API

API ini merupakan backend untuk aplikasi perbankan sederhana yang memfasilitasi interaksi antara nasabah dan merchant dengan bank. API ini dibangun menggunakan Golang dan framework Fiber, serta menggunakan database Postgres untuk menyimpan data.

## Fitur

*   **Login**: Autentikasi nasabah dengan username dan password.
*   **Payment**: Memungkinkan nasabah yang telah login untuk melakukan pembayaran ke merchant.
*   **Logout**:  Memungkinkan nasabah untuk logout.

## Teknologi yang Digunakan

*   Golang 1.21
*   Fiber Framework
*   Postgres Database
*   JWT (JSON Web Token)
*   Docker

## Menjalankan Aplikasi

1.  Clone repositori ini.
2.  Buat database Postgres dan update variabel environment `DATABASE_URL` di file `.env` dengan kredensial database Anda.
3.  Update variabel environment `JWT_SECRET` di file `.env` dengan kunci rahasia yang aman.
4.  Jalankan `make build` untuk membangun image Docker.
5.  Jalankan `make up` untuk menjalankan aplikasi.
6.  Jalankan `make seed` untuk mengisi database dengan data awal.

## Endpoint API

Dokumentasi API dapat diakses melalui Swagger di `/swagger/index.html` setelah aplikasi dijalankan.

## Menjalankan Test

*   Jalankan `make test` untuk menjalankan unit test.
