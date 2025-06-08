# 🎓 BrainQuiz - Backend API

> **Aplikasi Quiz Online untuk Platform Pembelajaran Digital**

[![Go Version](https://img.shields.io/badge/Go-1.24.2-blue.svg)](https://golang.org/)
[![Fiber](https://img.shields.io/badge/Fiber-v2-00ADD8.svg)](https://gofiber.io/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13+-336791.svg)](https://postgresql.org/)

## 📋 Deskripsi

**BrainQuiz** adalah platform pembelajaran digital yang memungkinkan siswa untuk mengikuti kuis online dengan berbagai kategori mata pelajaran. Aplikasi ini dibangun menggunakan **Go (Golang)** dengan framework **Fiber** dan **PostgreSQL** sebagai database, menyediakan API yang robust dan scalable untuk frontend aplikasi quiz.

## ✨ Fitur Utama

### 🔐 **Sistem Autentikasi & Otorisasi**
- Registrasi dan login pengguna dengan JWT
- Role-based access control (Admin, Teacher, Student)
- Password hashing dengan bcrypt
- Session management dengan cookies

### 👥 **Manajemen Pengguna**
- Multi-role system (Admin, Teacher, Student)
- Profile management
- User authentication middleware

### 📚 **Manajemen Konten Pembelajaran**
- **Kategori Soal**: Organisasi berdasarkan mata pelajaran
- **Tingkatan Kesulitan**: Klasifikasi soal (Mudah, Sedang, Sulit, Sangat Sulit)
- **Jenjang Pendidikan**: SD, SMP, SMA, SMK, Universitas
- **Manajemen Kelas**: Sistem kelas untuk mengorganisir siswa

### 🎯 **Sistem Kuis & Evaluasi**
- Pembuatan dan pengelolaan kuis
- Soal pilihan ganda dengan multiple options
- Real-time quiz submission
- Automatic scoring system
- Detailed answer tracking

### 📊 **Analytics & Reporting**
- Hasil kuis siswa dengan scoring
- Detail jawaban untuk analisis pembelajaran
- Progress tracking per siswa
- Performance analytics

## 🛠 Teknologi yang Digunakan

| Kategori | Teknologi | Versi |
|----------|-----------|-------|
| **Backend** | Go (Golang) | 1.24.2 |
| **Web Framework** | Fiber | v2.52.6 |
| **Database** | PostgreSQL | 13+ |
| **ORM** | GORM | v1.26.1 |
| **Authentication** | JWT | v3.2.2 |
| **Security** | bcrypt | - |
| **CORS** | Fiber CORS | - |

## 🗄 Struktur Database

Aplikasi menggunakan **10 tabel utama** dengan relasi yang terstruktur:

### 🏗 **Tabel Database**

| No | Tabel | Deskripsi | Relasi |
|----|-------|-----------|--------|
| 1 | **Users** | Data pengguna (admin, teacher, student) | One-to-Many dengan Hasil_Kuis, SoalAnswer, Kelas_Pengguna |
| 2 | **Kategori_Soal** | Kategori mata pelajaran | One-to-Many dengan Kuis |
| 3 | **Tingkatan** | Level kesulitan soal | One-to-Many dengan Kuis |
| 4 | **Kelas** | Kelas untuk mengorganisir siswa | Many-to-Many dengan Users |
| 5 | **Pendidikan** | Jenjang pendidikan (SD, SMP, SMA, dll) | One-to-Many dengan Kuis |
| 6 | **Kuis** | Data kuis/ujian | One-to-Many dengan Soal, Hasil_Kuis |
| 7 | **Soal** | Pertanyaan dalam kuis | One-to-Many dengan SoalAnswer |
| 8 | **Kelas_Pengguna** | Relasi many-to-many antara user dan kelas | Junction table |
| 9 | **Hasil_Kuis** | Hasil kuis yang dikerjakan siswa | - |
| 10 | **SoalAnswer** | Detail jawaban siswa untuk setiap soal | - |

### 📊 **Gambaran Umum Struktur Database**

![Database Schema](https://github.com/user-attachments/assets/5af6d057-ed50-4346-89a5-dddd4a2ec9ec)

**Skema database aplikasi mendukung entitas berikut:**

1. **Pengguna**: Menyimpan data pengguna seperti nama pengguna, email, peran, dan informasi kata sandi
2. **Kelas**: Mendefinisikan struktur kelas yang dapat diikuti oleh siswa dan guru
3. **Kategori_soal**: Kategori yang mengorganisir soal-soal kuis berdasarkan topik
4. **Tingkatan**: Mewakili berbagai tingkat kesulitan kuis
5. **Kuis**: Kuis yang sebenarnya, tempat pengguna berpartisipasi, dengan kaitan ke kategori dan tingkat kesulitan
6. **Soal**: Soal yang merupakan bagian dari setiap kuis, termasuk pilihan jawaban dan jawaban yang benar
7. **Kelas_pengguna**: Menghubungkan pengguna ke kelas yang mereka ikuti
8. **Hasil_kuis**: Menyimpan hasil untuk setiap pengguna setelah menyelesaikan kuis, termasuk skor dan waktu penyelesaian

## 🚀 Instalasi dan Setup

### 📋 Prerequisites

- **Go** 1.24.2 atau lebih baru
- **PostgreSQL** 13+
- **Git**

### 🔧 Langkah Instalasi

#### 1. **Clone Repository**
```bash
git clone https://github.com/Joko206/UAS_PWEB1.git
cd UAS_PWEB1
```

#### 2. **Install Dependencies**
```bash
go mod tidy
```

#### 3. **Setup Database**
Buat database PostgreSQL dan update konfigurasi di `database/database.go`:

```go
const (
    host     = "your-host"           // Default: ballast.proxy.rlwy.net
    port     = your-port             // Default: 46530
    user     = "your-username"       // Default: postgres
    password = "your-password"       // Update sesuai kebutuhan
    dbname   = "your-database-name"  // Default: railway
)
```

#### 4. **Build Aplikasi**
```bash
go build -o main .
```

#### 5. **Database Migration & Seeding**
```bash
# Jalankan migration dan seeding
./main seed

# Atau gunakan seeding sederhana
go run cmd/simple-seed/main.go

# Tambah data lebih banyak
go run cmd/add-more-data/main.go
```

#### 6. **Jalankan Server**
```bash
./main
```

🎉 **Server akan berjalan di** `http://localhost:8000`

## 📡 API Endpoints

### 🔐 **Authentication**
| Method | Endpoint | Deskripsi | Auth Required |
|--------|----------|-----------|---------------|
| `POST` | `/user/register` | Registrasi pengguna baru | ❌ |
| `POST` | `/user/login` | Login pengguna | ❌ |
| `GET` | `/user/logout` | Logout pengguna | ✅ |
| `GET` | `/user/get-user` | Get data pengguna yang sedang login | ✅ |

### 📚 **Kategori Soal** (Admin Only)
| Method | Endpoint | Deskripsi | Role |
|--------|----------|-----------|------|
| `GET` | `/kategori/get-kategori` | Get semua kategori | All |
| `POST` | `/kategori/add-kategori` | Tambah kategori baru | Admin |
| `PATCH` | `/kategori/update-kategori/:id` | Update kategori | Admin |
| `DELETE` | `/kategori/delete-kategori/:id` | Hapus kategori | Admin |

### 📊 **Tingkatan** (Admin Only)
| Method | Endpoint | Deskripsi | Role |
|--------|----------|-----------|------|
| `GET` | `/tingkatan/get-tingkatan` | Get semua tingkatan | All |
| `POST` | `/tingkatan/add-tingkatan` | Tambah tingkatan baru | Admin |
| `PATCH` | `/tingkatan/update-tingkatan/:id` | Update tingkatan | Admin |
| `DELETE` | `/tingkatan/delete-tingkatan/:id` | Hapus tingkatan | Admin |

### 🏫 **Kelas** (Admin & Teacher)
| Method | Endpoint | Deskripsi | Role |
|--------|----------|-----------|------|
| `GET` | `/kelas/get-kelas` | Get semua kelas | All |
| `POST` | `/kelas/add-kelas` | Tambah kelas baru | Admin, Teacher |
| `PATCH` | `/kelas/update-kelas/:id` | Update kelas | Admin, Teacher |
| `DELETE` | `/kelas/delete-kelas/:id` | Hapus kelas | Admin, Teacher |
| `POST` | `/kelas/join-kelas` | Join kelas | Student |
| `GET` | `/kelas/get-kelas-by-user` | Get kelas berdasarkan user | All |

### 🎓 **Pendidikan** (Admin Only)
| Method | Endpoint | Deskripsi | Role |
|--------|----------|-----------|------|
| `GET` | `/pendidikan/get-pendidikan` | Get semua jenjang pendidikan | All |
| `POST` | `/pendidikan/add-pendidikan` | Tambah jenjang pendidikan | Admin |
| `PATCH` | `/pendidikan/update-pendidikan/:id` | Update jenjang pendidikan | Admin |
| `DELETE` | `/pendidikan/delete-pendidikan/:id` | Hapus jenjang pendidikan | Admin |

### 🎯 **Kuis** (Admin & Teacher)
| Method | Endpoint | Deskripsi | Role |
|--------|----------|-----------|------|
| `GET` | `/kuis/get-kuis` | Get semua kuis | All |
| `POST` | `/kuis/add-kuis` | Tambah kuis baru | Admin, Teacher |
| `PATCH` | `/kuis/update-kuis/:id` | Update kuis | Admin, Teacher |
| `DELETE` | `/kuis/delete-kuis/:id` | Hapus kuis | Admin, Teacher |
| `GET` | `/kuis/filter-kuis` | Filter kuis berdasarkan kriteria | All |

### ❓ **Soal** (Admin & Teacher)
| Method | Endpoint | Deskripsi | Role |
|--------|----------|-----------|------|
| `GET` | `/soal/get-soal` | Get semua soal | All |
| `GET` | `/soal/get-soal/:kuis_id` | Get soal berdasarkan kuis | All |
| `POST` | `/soal/add-soal` | Tambah soal baru | Admin, Teacher |
| `PATCH` | `/soal/update-soal/:id` | Update soal | Admin, Teacher |
| `DELETE` | `/soal/delete-soal/:id` | Hapus soal | Admin, Teacher |

### 📈 **Hasil Kuis**
| Method | Endpoint | Deskripsi | Role |
|--------|----------|-----------|------|
| `GET` | `/hasil-kuis/:user_id/:kuis_id` | Get hasil kuis spesifik | All |
| `POST` | `/hasil-kuis/submit-jawaban` | Submit jawaban kuis | Student |

## 📁 Struktur Project

```
UAS_PWEB1/
├── 📁 cmd/                    # Command line utilities
│   ├── 📁 seed/              # Database seeding
│   ├── 📁 check/             # Database checking utility
│   ├── 📁 simple-seed/       # Simple seeding script
│   └── 📁 add-more-data/     # Additional data seeding
├── 📁 controllers/           # HTTP handlers
│   ├── 📄 user.go           # User authentication & management
│   ├── 📄 kategori.go       # Category management
│   ├── 📄 tingkatan.go      # Difficulty level management
│   ├── 📄 kelas.go          # Class management
│   ├── 📄 pendidikan.go     # Education level management
│   ├── 📄 Kuis.go           # Quiz management
│   ├── 📄 soal.go           # Question management
│   ├── 📄 HasilKuis.go      # Quiz results
│   ├── 📄 Kelas_Pengguna.go # User-class relationships
│   └── 📄 response.go       # Response helpers
├── 📁 database/             # Database layer
│   ├── 📄 database.go       # Database connection & migration
│   ├── 📄 seed.go           # Database seeding functions
│   ├── 📄 kategori.go       # Category database operations
│   ├── 📄 tingkatan.go      # Difficulty level database operations
│   ├── 📄 kelas.go          # Class database operations
│   ├── 📄 pendidikan.go     # Education level database operations
│   ├── 📄 kuis.go           # Quiz database operations
│   └── 📄 soal.go           # Question database operations
├── 📁 models/               # Data models
│   └── 📄 models.go         # GORM models
├── 📁 routes/               # Route definitions
│   └── 📄 routes.go         # API routes setup
├── 📄 main.go               # Application entry point
├── 📄 go.mod                # Go module dependencies
└── 📄 README.md             # Project documentation
```

## 👥 Sistem Role & Permission

### 🔑 **Admin**
- ✅ Full access ke semua fitur
- ✅ Dapat mengelola kategori, tingkatan, dan pendidikan
- ✅ Dapat mengelola kuis dan soal
- ✅ Dapat melihat semua hasil kuis
- ✅ Dapat mengelola semua pengguna

### 👨‍🏫 **Teacher**
- ✅ Dapat mengelola kelas
- ✅ Dapat membuat dan mengelola kuis
- ✅ Dapat membuat dan mengelola soal
- ✅ Dapat melihat hasil kuis
- ❌ Tidak dapat mengelola kategori dan tingkatan

### 👨‍🎓 **Student**
- ✅ Dapat join kelas
- ✅ Dapat mengikuti kuis
- ✅ Dapat melihat hasil kuis sendiri
- ❌ Tidak dapat membuat kuis atau soal

## 🔒 Authentication & Security

- **JWT Token**: Digunakan untuk autentikasi dengan expiry 24 jam
- **Password Hashing**: Menggunakan bcrypt dengan cost 14
- **Role-based Access Control**: Middleware untuk mengontrol akses berdasarkan role
- **CORS**: Dikonfigurasi untuk frontend yang diizinkan
- **Cookie Support**: JWT dapat disimpan dalam cookie HTTPOnly

## 🌱 Database Seeding

Aplikasi menyediakan beberapa cara untuk melakukan seeding database:

### 🚀 **Full Seeding**
```bash
./main seed
```
Menjalankan seeding lengkap untuk semua tabel dengan data komprehensif.

### ⚡ **Simple Seeding**
```bash
go run cmd/simple-seed/main.go
```
Seeding sederhana dengan data minimal yang diperlukan.

### 📊 **Add More Data**
```bash
go run cmd/add-more-data/main.go
```
Menambahkan lebih banyak data ke database yang sudah ada.

### 🔍 **Check Database**
```bash
go run cmd/check/main.go
```
Memeriksa isi database dan menampilkan statistik data.

## 📊 Sample Data

Setelah seeding, database akan berisi:

| Tabel | Jumlah Data | Deskripsi |
|-------|-------------|-----------|
| **Users** | 21 | 1 admin, 4 teachers, 16 students |
| **Kategori_Soal** | 10 | Matematika, Bahasa Indonesia, IPA, dll |
| **Tingkatan** | 4 | Mudah, Sedang, Sulit, Sangat Sulit |
| **Kelas** | 12 | Kelas 1-12 |
| **Pendidikan** | 5 | SD, SMP, SMA, SMK, Universitas |
| **Kuis** | 15 | Berbagai topik dan tingkatan |
| **Soal** | 19+ | Soal pilihan ganda dengan 4 opsi |
| **Kelas_Pengguna** | 18+ | Assignment siswa ke kelas |
| **Hasil_Kuis** | 21+ | Hasil kuis dengan scoring |
| **SoalAnswer** | 29+ | Detail jawaban siswa |

### 👤 **Sample Users**
```json
{
  "admin": {
    "email": "admin@example.com",
    "password": "password123",
    "role": "admin"
  },
  "teacher": {
    "email": "sarah.johnson@example.com",
    "password": "password123",
    "role": "teacher"
  },
  "student": {
    "email": "alice.smith@example.com",
    "password": "password123",
    "role": "student"
  }
}
```

## 🧪 Development & Testing

### 🔧 **Menambah Endpoint Baru**

1. **Buat function handler** di folder `controllers/`
2. **Tambahkan database operation** di folder `database/`
3. **Daftarkan route** di `routes/routes.go`
4. **Tambahkan middleware** authentication/authorization jika diperlukan

### 🧪 **Testing API**

Gunakan tools seperti **Postman**, **Insomnia**, atau **curl** untuk testing:

#### **Login**
```bash
curl -X POST http://localhost:8000/user/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "password123"
  }'
```

#### **Get Categories (dengan token)**
```bash
curl -X GET http://localhost:8000/kategori/get-kategori \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

#### **Create Quiz**
```bash
curl -X POST http://localhost:8000/kuis/add-kuis \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Quiz Matematika Baru",
    "description": "Quiz untuk menguji kemampuan matematika",
    "kategori_id": 1,
    "tingkatan_id": 2,
    "kelas_id": 7,
    "pendidikan_id": 2
  }'
```

## 🚀 Deployment

### 🌐 **Production Setup**

1. **Environment Variables**: Pindahkan konfigurasi database ke environment variables
2. **HTTPS**: Setup SSL/TLS certificate
3. **Database**: Gunakan managed PostgreSQL service (Railway, Supabase, AWS RDS)
4. **Monitoring**: Setup logging dan monitoring (Prometheus, Grafana)
5. **Load Balancer**: Untuk high availability

### 🐳 **Docker Deployment**

Buat `Dockerfile`:
```dockerfile
FROM golang:1.24.2-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
EXPOSE 8000

CMD ["./main"]
```

**Build dan Run:**
```bash
docker build -t brainquiz-backend .
docker run -p 8000:8000 brainquiz-backend
```

### ☁️ **Cloud Deployment**

#### **Railway**
```bash
# Install Railway CLI
npm install -g @railway/cli

# Login dan deploy
railway login
railway init
railway up
```

#### **Heroku**
```bash
# Install Heroku CLI dan deploy
heroku create brainquiz-backend
git push heroku main
```

## 🤝 Contributing

Kami menyambut kontribusi dari developer lain! Ikuti langkah berikut:

1. **Fork** repository ini
2. **Buat feature branch** (`git checkout -b feature/AmazingFeature`)
3. **Commit changes** (`git commit -m 'Add some AmazingFeature'`)
4. **Push ke branch** (`git push origin feature/AmazingFeature`)
5. **Buat Pull Request**

### 📝 **Contribution Guidelines**

- Ikuti Go coding standards
- Tambahkan tests untuk fitur baru
- Update dokumentasi jika diperlukan
- Pastikan semua tests pass

## 📄 License

Distributed under the **MIT License**. See `LICENSE` for more information.

## 📞 Contact & Support

- **Developer**: Joko - [@Joko206](https://github.com/Joko206)
- **Project Link**: [https://github.com/Joko206/UAS_PWEB1](https://github.com/Joko206/UAS_PWEB1)
- **Issues**: [GitHub Issues](https://github.com/Joko206/UAS_PWEB1/issues)

## 🙏 Acknowledgments

- [**Fiber**](https://gofiber.io/) - Express-inspired web framework for Go
- [**GORM**](https://gorm.io/) - The fantastic ORM library for Golang
- [**JWT-Go**](https://github.com/golang-jwt/jwt) - JWT implementation for Go
- [**bcrypt**](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Password hashing library
- [**PostgreSQL**](https://postgresql.org/) - Advanced open source database

---

<div align="center">

**⭐ Jika project ini membantu, berikan star di GitHub! ⭐**

Made with ❤️ by [Joko](https://github.com/Joko206)

</div>
