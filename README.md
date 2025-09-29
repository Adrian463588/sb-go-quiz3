
# Brief Document — Quiz Bootcamp Intensif Golang

**Project:** Golang  
**Document No.:** N/A  
**Date:** N/A  
**Revision No.:** N/A  
**Revision Date:** N/A  
**Reference No.:** N/A  
**Assignee:** N/A  
**Mandays:** 24 Hours  
**Target Release:** N/A

---

## Brief Description
Mini Project: Membuat sebuah project untuk input data buku dan kategorinya.  
Database yang digunakan adalah relational database (MySQL atau PostgreSQL).  
Deploy production project di Railway.  
Mohon mengikuti instruksi yang diberikan.

---

## Setup Project
- Kerjakan dalam repository Git (GitHub).  
  - Boleh menggunakan repository sebelumnya (tugas harian) atau membuat repository baru untuk quiz.
- Struktur folder: wajib menerapkan package dan import dengan baik (struktur modular sesuai best practice Go).
- Paket yang direkomendasikan:
  - [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
  - [github.com/lib/pq](https://github.com/lib/pq) (jika pakai PostgreSQL)
  - [github.com/rubenv/sql-migrate](https://github.com/rubenv/sql-migrate)
- Inisialisasi project dengan Go modules (`go mod init ...`).

---

## Database Design

### Tabel `books`
| Column       | Data Type  | Notes |
|--------------|------------|-------|
| id           | integer    | primary key, auto increment |
| title        | varchar    | required |
| description  | varchar    | optional |
| image_url    | varchar    | optional |
| release_year | integer    | validasi: 1980 - 2024 |
| price        | integer    | optional |
| total_page   | integer    | required |
| thickness    | varchar    | derived: `"tebal"` / `"tipis"` berdasarkan `total_page` |
| category_id  | integer    | foreign key -> categories.id |
| created_at   | timestamp  | |
| created_by   | varchar    | |
| modified_at  | timestamp  | |
| modified_by  | varchar    | |

### Tabel `categories`
| Column       | Data Type  | Notes |
|--------------|------------|-------|
| id           | integer    | primary key, auto increment |
| name         | varchar    | required, unique |
| created_at   | timestamp  | |
| created_by   | varchar    | |
| modified_at  | timestamp  | |
| modified_by  | varchar    | |

### Tabel `users`
| Column       | Data Type  | Notes |
|--------------|------------|-------|
| id           | integer    | primary key, auto increment |
| username     | varchar    | required, unique |
| password     | varchar    | hashed password |
| created_at   | timestamp  | |
| created_by   | varchar    | |
| modified_at  | timestamp  | |
| modified_by  | varchar    | |

**Catatan konversi `thickness`:**
- Jika `total_page` > 100 → `thickness = "tebal"`.
- Jika `total_page` ≤ 100 → `thickness = "tipis"`.

---

## Detail Feature — Authentication
Pilih salah satu mekanisme authentication untuk mengakses API:
- **Basic Auth** (middleware)
- **JSON Web Token (JWT)** (middleware)

Jika memilih **JWT**, buat endpoint login:
- `POST /api/users/login` — menerima credential, mengembalikan token JWT.

Semua endpoint yang sensitif (create/update/delete) harus dilindungi oleh middleware (Basic Auth atau JWT).

---

## Detail Feature — Kategori (Categories)

**Base path:** `/api/categories`

| Path | Method | Deskripsi | Auth |
|------|--------|----------:|:----:|
| `/api/categories` | `GET` | Menampilkan seluruh kategori | Required |
| `/api/categories` | `POST` | Menambahkan kategori | Required |
| `/api/categories/:id` | `GET` | Menampilkan detail kategori | Required |
| `/api/categories/:id` | `DELETE` | Menghapus kategori | Required |
| `/api/categories/:id/books` | `GET` | Menampilkan buku berdasarkan kategori | Required |

**Validasi & Error handling:**
- Gunakan middleware Basic Auth / JWT untuk akses.
- Jika menghapus kategori yang tidak tersedia → kembalikan error `404` dengan message informatif.
- Jika mengupdate kategori yang tidak tersedia → kembalikan error `404`.

**Contoh request (JSON) untuk POST /api/categories**
```json
{
  "name": "Programming"
}
````

---

## Detail Feature — Buku (Books)

**Base path:** `/api/books`

| Path             | Method   |                Deskripsi |   Auth   |
| ---------------- | -------- | -----------------------: | :------: |
| `/api/books`     | `GET`    | Menampilkan seluruh buku | Required |
| `/api/books`     | `POST`   |         Menambahkan buku | Required |
| `/api/books/:id` | `GET`    |  Menampilkan detail buku | Required |
| `/api/books/:id` | `DELETE` |           Menghapus buku | Required |

**Validasi & Business rules:**

* Gunakan middleware Basic Auth / JWT untuk akses.
* Jika menghapus / update buku yang tidak tersedia → kembalikan error `404`.
* `release_year` harus berada di rentang **1980 — 2024**; jika tidak, kembalikan error `400` dengan message yang jelas.
* `thickness` tidak diberikan langsung oleh client — dihitung berdasarkan `total_page`:

  * `total_page > 100` → `thickness = "tebal"`
  * `total_page <= 100` → `thickness = "tipis"`

**Contoh request (JSON) untuk POST /api/books**

```json
{
  "title": "Contoh Buku",
  "description": "Deskripsi singkat",
  "image_url": "https://...",
  "release_year": 2020,
  "price": 50000,
  "total_page": 150,
  "category_id": 1
}
```

(Server akan menyimpan `thickness: "tebal"` berdasarkan `total_page`.)

---

## Deployment

* Deploy ke **Railway** (production).
* Sertakan environment variables pada Railway untuk koneksi DB, JWT secret, dsb.

---

## Documentation

* Perbarui `README.md` pada repository dengan:

  * Cara setup (local & production)
  * List environment variables
  * Struktur folder singkat
  * List path (endpoints) + contoh request/response
* Opsional: tambahkan **Swagger** (OpenAPI) untuk dokumentasi API interaktif.

---

## Acceptance Criteria / Checklist

* [ ] Repo GitHub tersedia (link).
* [ ] Project menggunakan Go modules.
* [ ] Struktur package jelas (controller, service, repository, model, middleware, config).
* [ ] Migrasi DB tersedia (sql-migrate atau tool lain).
* [ ] Authentication (Basic Auth / JWT) berjalan.
* [ ] Semua endpoint categories & books tersedia dan ter-protect.
* [ ] Validasi release_year diterapkan.
* [ ] Konversi `thickness` berdasarkan `total_page` diterapkan.
* [ ] Deploy ke Railway berhasil.
* [ ] README / (opsional) Swagger tersedia.

---


