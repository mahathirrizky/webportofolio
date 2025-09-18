# Go-Vue Portfolio Web

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Vue.js](https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vue.js&logoColor=4FC08D)
![Vite](https://img.shields.io/badge/Vite-B73BFE?style=for-the-badge&logo=vite&logoColor=FFD62E)
![SQLite](https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white)

## 📸 Screenshot

![Screenshot](screenshots/ss.jpeg)


Aplikasi web portofolio full-stack yang dibangun dengan backend Go (Gin) dan frontend Vue.js (Vite). Proyek ini mencakup panel admin untuk mengelola konten portofolio secara dinamis.

## 🌟 Fitur

- **Backend Go (Gin):** Backend yang cepat dan andal.
- **Frontend Vue.js (Vite):** UI yang modern dan reaktif.
- **Database SQLite:** Ringan dan mudah di-setup.
- **Otentikasi JWT:** Mengamankan rute admin.
- **Panel Admin:**
  - Kelola konten (CRUD - Create, Read, Update, Delete).
  - Lihat dan hapus pesan dari pengunjung.
  - Unggah file/gambar.
- **Penyajian Tunggal:** Server Go menyajikan build file frontend, menjadikannya aplikasi mandiri.

## 🛠️ Tumpukan Teknologi

- **Backend:** Go, Gin
- **Frontend:** Vue.js, Vue Router, Vite
- **Database:** SQLite
- **Styling:** CSS Standar

## 🚀 Memulai

Untuk menjalankan proyek ini secara lokal, ikuti langkah-langkah berikut.

### Prasyarat

- [Go](https://golang.org/doc/install) (versi 1.18 atau lebih baru)
- [Node.js](https://nodejs.org/en/download/) (versi 16 atau lebih baru)
- `npm`

### 1. Backend

Kloning repositori:
```bash
git clone https://github.com/mahathirrizky/webportofolio.git
cd webportofolio
```

Buat file `.env` di direktori root dan tambahkan variabel berikut. Ganti `your-secret-key` dengan kunci rahasia Anda.
```
SECRET_KEY=your-secret-key
PORT=8080
```

Jalankan server backend:
```bash
go run main.go
```
Server akan berjalan di `http://localhost:8080`.

### 2. Frontend (untuk pengembangan)

Buka terminal baru, masuk ke direktori `ui`:
```bash
cd ui
```

Instal dependensi:
```bash
npm install
```

Jalankan server pengembangan frontend:
```bash
npm run dev
```
Aplikasi frontend akan dapat diakses di `http://localhost:5173` dan akan terhubung ke backend Go.

## 📦 Build untuk Produksi

Untuk membuat build produksi dari aplikasi Vue dan menjalankannya dengan server Go:

1.  **Build aplikasi Vue:**
    ```bash
    cd ui
    npm run build
    ```
    Ini akan membuat direktori `dist` di dalam `ui`.

2.  **Jalankan server Go:**
    Dari direktori root, jalankan:
    ```bash
    go run main.go
    ```
    Server Go sekarang akan menyajikan file yang telah di-build dari `ui/dist`. Buka `http://localhost:8080` di browser Anda.

## Endpoints API

### Rute Publik
- `GET /api/content`: Mengambil semua konten publik.
- `POST /api/messages`: Mengirim pesan baru (form kontak).

### Rute Admin (memerlukan token JWT)
- `POST /api/auth/login`: Login untuk mendapatkan token.
- `GET /api/admin/content`: Mengambil semua konten.
- `POST /api/admin/content`: Membuat konten baru.
- `PUT /api/admin/content/:id`: Memperbarui konten.
- `DELETE /api/admin/content/:id`: Menghapus konten.
- `POST /api/admin/upload`: Mengunggah file.
- `GET /api/admin/messages`: Mengambil semua pesan.
- `DELETE /api/admin/messages/:id`: Menghapus pesan.

## ❤️ Dukungan

Jika proyek ini berguna bagi Anda, pertimbangkan untuk mentraktir saya kopi!

[![Buy Me a Coffee](https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png)](https://www.buymeacoffee.com/mahathirrizky)

## 📄 Lisensi

Proyek ini dilisensikan di bawah Lisensi MIT. Lihat file `LICENSE` untuk detailnya.
