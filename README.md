# To-Do List API (Go & Clean Architecture)

Sebuah RESTful API sederhana namun kokoh untuk mengelola daftar tugas (To-Do List), dibangun dengan Go (Golang). Proyek ini dirancang dengan fokus pada kualitas kode, skalabilitas, dan kemudahan pemeliharaan dengan mengimplementasikan prinsip-prinsip **Clean Architecture**.

Proyek ini adalah portofolio untuk menunjukkan pemahaman dalam membangun layanan backend yang terstruktur, siap untuk lingkungan produksi, dan mengikuti praktik terbaik dalam industri.

## ✨ Fitur Utama

-   **Arsitektur Bersih**: Memisahkan logika bisnis dari detail teknis (framework, database), membuat aplikasi lebih mudah diuji, dirawat, dan diskalakan.
-   **RESTful API**: Menyediakan endpoint yang jelas dan standar industri untuk operasi CRUD (Create, Read, Update, Delete) tugas.
-   **Penanganan Error yang Elegan**: Menggunakan *sentinel errors* (`ErrNotFound`, `ErrInternal`) untuk memisahkan kesalahan domain bisnis dari kesalahan teknis, menghasilkan kode yang lebih bersih dan penanganan kasus yang lebih prediktif.
-   **Interaksi Database dengan ORM**: Menggunakan GORM untuk abstraksi dan interaksi dengan database PostgreSQL.
-   **Manajemen Konfigurasi**: Mengelola kredensial dan konfigurasi lingkungan secara aman melalui file `.env`.

## 🏛️ Arsitektur

Aplikasi ini mengadopsi prinsip **Clean Architecture** untuk memastikan pemisahan tanggung jawab (*separation of concerns*) yang jelas. Ketergantungan kode hanya mengarah ke dalam, dari lapisan luar yang konkret ke lapisan dalam yang abstrak.

-   **Domain**: Lapisan inti yang berisi model entitas (`Task`) dan definisi *error* aplikasi. Lapisan ini tidak bergantung pada lapisan lain.
-   **Repository**: Mendefinisikan *interface* untuk abstraksi penyimpanan data dan implementasinya yang berinteraksi langsung dengan database menggunakan GORM.
-   **Service**: Berisi logika bisnis utama aplikasi. Lapisan ini bergantung pada *interface repository* untuk mengakses data.
-   **Handler**: Bertindak sebagai *controller* yang menerima permintaan HTTP dari Gin, memanggil metode yang sesuai di lapisan *service*, dan mengembalikan respons dalam format JSON.
-   **Main & Routes**: Titik masuk aplikasi yang menginisialisasi semua komponen (DI), melakukan migrasi database, dan mendefinisikan rute-rute API.

## 🛠️ Teknologi yang Digunakan

-   **Bahasa**: Go (Golang)
-   **Web Framework**: Gin
-   **ORM**: GORM
-   **Database**: PostgreSQL
-   **Manajemen Konfigurasi**: godotenv

## 🚀 Memulai Proyek

Untuk menjalankan proyek ini di lingkungan lokal, ikuti langkah-langkah berikut.

### Prasyarat

-   Go (versi yang digunakan dalam pengembangan adalah `go1.24.3`)
-   PostgreSQL sedang berjalan
-   Git

### Instalasi & Menjalankan

1.  **Clone repositori ini:**
    ```sh
    git clone [https://github.com/USERNAME/NAMA_REPO.git](https://github.com/USERNAME/NAMA_REPO.git)
    cd NAMA_REPO
    ```

2.  **Install dependensi:**
    ```sh
    go mod tidy
    ```

3.  **Buat database** di PostgreSQL dengan nama yang Anda inginkan.

4.  **Siapkan file `.env`**:
    Proyek ini menggunakan file `.env` untuk konfigurasi. Buat file bernama `.env` di root proyek dan isi dengan konfigurasi database Anda. Contoh:
    ```env
    DB_HOST=localhost
    DB_USER=postgres
    DB_PASS=password_anda
    DB_NAME=nama_database_anda
    DB_PORT=5432
    DB_SSLMODE=disable
    ```

5.  **Jalankan aplikasi:**
    Aplikasi akan melakukan migrasi skema database (`Task`) secara otomatis saat pertama kali dijalankan.
    ```sh
    go run main.go
    ```
    Aplikasi akan berjalan di `http://localhost:8080`.

## 🔌 Dokumentasi API

Berikut adalah daftar *endpoint* yang tersedia.

---

### 1. Membuat Tugas Baru

-   **Endpoint**: `POST /task`
-   **Deskripsi**: Menambahkan tugas baru ke dalam daftar.
-   **Request Body** (`CreateTaskRequest`):
    ```json
    {
      "title": "Selesaikan Laporan Proyek",
      "description": "Tulis bagian kesimpulan dan analisis akhir."
    }
    ```
-   **Contoh Respons Sukses (201 Created)**:
    ```json
    {
      "status": "success",
      "message": "Task created successfully",
      "data": {
        "ID": 1,
        "Title": "Selesaikan Laporan Proyek",
        "Description": "Tulis bagian kesimpulan dan analisis akhir.",
        "CreatedAt": "2025-06-30T15:10:00Z",
        "UpdatedAt": "2025-06-30T15:10:00Z"
      }
    }
    ```

---

### 2. Mendapatkan Semua Tugas

-   **Endpoint**: `GET /task`
-   **Deskripsi**: Mengambil semua tugas yang ada.
-   **Contoh Respons Sukses (200 OK)**:
    ```json
    {
      "status": "success",
      "message": "Tasks fetched successfully",
      "data": [
        {
          "ID": 1,
          "Title": "Selesaikan Laporan Proyek",
          "Description": "Tulis bagian kesimpulan dan analisis akhir.",
          "CreatedAt": "2025-06-30T15:10:00Z",
          "UpdatedAt": "2025-06-30T15:10:00Z"
        }
      ]
    }
    ```

---

### 3. Mendapatkan Tugas Berdasarkan ID

-   **Endpoint**: `GET /task/:id`
-   **Deskripsi**: Mengambil detail satu tugas spesifik.
-   **Contoh Respons Sukses (200 OK)**:
    ```json
    {
      "status": "success",
      "message": "Task fetched successfully",
      "data": {
        "ID": 1,
        "Title": "Selesaikan Laporan Proyek",
        "Description": "Tulis bagian kesimpulan dan analisis akhir.",
        "CreatedAt": "2025-06-30T15:10:00Z",
        "UpdatedAt": "2025-06-30T15:10:00Z"
      }
    }
    ```

---

### 4. Memperbarui Tugas

-   **Endpoint**: `PUT /task/:id`
-   **Deskripsi**: Memperbarui judul atau deskripsi tugas yang sudah ada.
-   **Request Body** (`UpdateTaskRequest`):
    ```json
    {
      "title": "Selesaikan Laporan Proyek Akhir",
      "description": "Revisi bagian kesimpulan."
    }
    ```
-   **Contoh Respons Sukses (200 OK)**:
    ```json
    {
      "status": "success",
      "message": "Task updated successfully",
      "data": {
        "ID": 1,
        "Title": "Selesaikan Laporan Proyek Akhir",
        "Description": "Revisi bagian kesimpulan.",
        "CreatedAt": "2025-06-30T15:10:00Z",
        "UpdatedAt": "2025-06-30T15:15:00Z"
      }
    }
    ```

---

### 5. Menghapus Tugas

-   **Endpoint**: `DELETE /task/:id`
-   **Deskripsi**: Menghapus tugas dari daftar.
-   **Contoh Respons Sukses (200 OK)**:
    ```json
    {
      "status": "success",
      "message": "Task deleted successfully"
    }
    ```
