# To-Do List API (Go & Clean Architecture)

Sebuah RESTful API sederhana namun kokoh untuk mengelola daftar tugas (To-Do List), dibangun dengan Go (Golang). Proyek ini dirancang dengan fokus pada kualitas kode, skalabilitas, dan kemudahan pemeliharaan dengan mengimplementasikan prinsip-prinsip **Clean Architecture**.

Proyek ini adalah portofolio saya untuk menunjukkan pemahaman dalam membangun layanan backend yang terstruktur, siap untuk lingkungan produksi, dan mengikuti praktik terbaik dalam industri.

## ✨ Fitur Utama

-   **CRUD Penuh untuk Tugas**: Buat, Baca, Perbarui, dan Hapus tugas.
-   **Arsitektur Bersih**: Memisahkan logika bisnis dari detail teknis (framework, database), membuat aplikasi lebih mudah diuji, dirawat, dan diskalakan.
-   **Penanganan Error yang Elegan**: Menggunakan *sentinel errors* (`ErrNotFound`, `ErrInternal`) untuk memisahkan kesalahan domain bisnis dari kesalahan teknis, menghasilkan kode yang lebih bersih dan penanganan kasus yang lebih prediktif.
-   **RESTful API**: Menyediakan *endpoint* yang jelas dan standar industri.
-   **Manajemen Konfigurasi**: Mengelola kredensial dan konfigurasi lingkungan secara aman melalui file `.env`.
-   **Validasi Request**: Memastikan data yang masuk sesuai dengan format yang diharapkan.
