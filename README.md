# Go Gin Employees Training Enrollment System

Nextjs Frontend: [https://github.com/yehezkiel1086/nextjs-employees-training-enrollment-system](https://github.com/yehezkiel1086/nextjs-employees-training-enrollment-system)

Sebuah aplikasi web *full-stack* sederhana untuk mengelola program pelatihan karyawan. Admin dapat membuat dan mengelola sesi pelatihan, sementara karyawan dapat menelusuri dan mendaftar untuk pelatihan.

![Flowchart Sistem Pelatihan Karyawan](/assets/employee_training_system_flowchart.png)

![ERD Sistem Pelatihan Karyawan](/assets/employee_training_system_erd.png)

## Fitur Utama

### Manajemen Pengguna

- Terdapat peran **Karyawan** dan **Admin** yang datanya disimpan di dalam database.
- Akses berbasis peran:
    - **Admin:** Dapat melakukan operasi CRUD (Create, Read, Update, Delete) pada data pelatihan dan melihat semua data pendaftaran.
    - **Karyawan:** Dapat melihat daftar pelatihan yang tersedia, mendaftar, dan membatalkan pendaftaran.

### Manajemen Pelatihan

- Operasi CRUD (Create, Read, Update, Delete) untuk data pelatihan.
- Setiap data pelatihan memiliki informasi seperti: judul, deskripsi, tanggal, dan instruktur.

### Pendaftaran

- Karyawan dapat mendaftar atau membatalkan pendaftaran pada sebuah pelatihan.
- Setiap pendaftaran akan menghubungkan data pengguna (karyawan) dengan sesi pelatihan yang dipilih.

### Autentikasi

- Proses registrasi dan login menggunakan **autentikasi JWT** yang ditangani oleh backend (Golang).
- Rute yang terproteksi hanya bisa diakses dengan menyertakan *bearer token*.
- Token disimpan di sisi frontend (Next.js) menggunakan *cookie* atau *localStorage*.

## Teknologi yang Digunakan

- **Backend:** Go (Gin Framework)
- **Database:** PostgreSQL
- **Caching:** Redis
- **Containerization:** Docker

## Cara Instalasi dan Menjalankan Proyek

Proyek ini dijalankan menggunakan Docker. Pastikan Anda sudah menginstal Docker dan Docker Compose di sistem Anda.

1.  **Clone repository ini:**
    ```bash
    git clone https://github.com/username/go-gin-employees-training-enrollment-system.git
    cd go-gin-employees-training-enrollment-system
    ```

2.  **Buat file `.env`:**
    Salin file `.env.example` menjadi `.env` dan sesuaikan nilainya jika diperlukan.
    ```bash
    cp .env.example .env
    ```
    File `.env` Anda akan terlihat seperti ini:
    ```env
    # PostgreSQL
    DB_USER=admin
    DB_PASS=secret
    DB_NAME=employee_training
    DB_PORT=5432

    # Redis
    REDIS_PASS=supersecret
    REDIS_PORT=6379
    ```

3.  **Jalankan dengan Docker Compose:**
    Buka terminal di root direktori proyek dan jalankan perintah berikut untuk membangun dan menjalankan *container* database (PostgreSQL) dan Redis.
    ```bash
    docker-compose up -d
    ```

4.  **Jalankan Aplikasi Go:**
    (Asumsi aplikasi Go belum di-containerize) Jalankan aplikasi Go Anda. Aplikasi akan terhubung ke database dan Redis yang berjalan di dalam Docker.
    ```bash
    go run cmd/main.go
    ```

> Anda juga dapat memanfaatkan [Makefile](Makefile)

## Contoh API Endpoint

| Method | Endpoint | Deskripsi | Akses |
|:---|:---|:---|:---|
| `POST` | `/api/v1/register` | Registrasi pengguna baru (karyawan). | Publik |
| `POST` | `/api/v1/login` | Login untuk mendapatkan token JWT. | Publik |
| `GET` | `/api/v1/trainings` | Mendapatkan semua daftar pelatihan. | Karyawan, Admin |
| `GET` | `/api/v1/trainings/:id` | Mendapatkan detail satu pelatihan. | Karyawan, Admin |
| `POST` | `/api/v1/trainings` | Membuat pelatihan baru. | Admin |
| `PUT` | `/api/v1/trainings/:id` | Memperbarui data pelatihan. | Admin |
| `DELETE` | `/api/v1/trainings/:id` | Menghapus data pelatihan. | Admin |
| `POST` | `/api/v1/enrollments` | Mendaftarkan karyawan ke pelatihan. | Karyawan |
| `DELETE` | `/api/v1/enrollments/:id` | Membatalkan pendaftaran pelatihan. | Karyawan |
| `GET` | `/api/enrollments/:email` | Mendapatkan semua data pendaftaran untuk user yang terotentikasi. | Karyawan |
| `GET` | `/api/v1/enrollments` | Mendapatkan semua data pendaftaran. | Admin |

