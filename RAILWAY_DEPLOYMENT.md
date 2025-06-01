# Railway Deployment Guide

## ✅ MASALAH SUDAH DIPERBAIKI

### Masalah Go Version Compatibility
**FIXED**: Error `go.mod requires go >= 1.23.0 (running go 1.21.13)` sudah diperbaiki dengan:
- Update `go.mod` ke versi 1.23.0
- Update `Dockerfile` menggunakan `golang:1.23-alpine`
- Tambah `nixpacks.toml` untuk Railway configuration
- Update `railway.toml` dengan GO_VERSION = "1.23.0"

### Masalah Binary Not Found
**FIXED**: Error `./main: No such file or directory` sudah diperbaiki dengan:
- Tambah `start.sh` script yang robust
- Auto-build binary jika tidak ditemukan
- Proper chmod +x untuk executable permissions
- Debug logging untuk troubleshooting
- Update Procfile, nixpacks.toml, dan Dockerfile

### 502 Error - Connection Refused
Error ini biasanya terjadi karena:
1. ❌ ~~Go version compatibility~~ ✅ **FIXED**
2. Database connection gagal
3. Port configuration salah
4. Environment variables tidak diset

## Langkah Deployment

### 1. Persiapan Environment Variables di Railway
Pastikan environment variables berikut sudah diset di Railway dashboard:

```
DATABASE_URL=postgresql://username:password@host:port/database?sslmode=require
PORT=8000 (biasanya auto-set oleh Railway)
CORS_ORIGINS=https://your-frontend-domain.com
```

### 2. Deploy ke Railway
Ada beberapa cara:

#### A. Menggunakan Git (Recommended)
```bash
git add .
git commit -m "Deploy to Railway"
git push origin main
```

#### B. Menggunakan Railway CLI
```bash
railway login
railway link
railway up
```

### 3. Monitoring Deployment
1. Buka Railway dashboard
2. Pilih project Anda
3. Lihat tab "Deployments" untuk status
4. Lihat tab "Logs" untuk debug

### 4. Testing
Setelah deployment berhasil:
```bash
curl https://your-app.railway.app/health
```

## Troubleshooting

### Jika masih 502 Error:
1. **Cek Logs Railway**: Lihat error message di Railway dashboard
2. **Cek Database Connection**: Pastikan DATABASE_URL benar
3. **Cek Port**: Railway biasanya auto-set PORT variable
4. **Cek SSL Mode**: Gunakan `sslmode=require` untuk Railway PostgreSQL

### Format DATABASE_URL yang Benar:
```
postgresql://username:password@host:port/database?sslmode=require
```

### Contoh Environment Variables:
```
DATABASE_URL=postgresql://postgres:password@containers-us-west-123.railway.app:5432/railway?sslmode=require
PORT=8000
CORS_ORIGINS=https://brainquiz-psi.vercel.app
```

## Files yang Ditambahkan untuk Railway:
- `railway.toml` - Railway configuration
- `Dockerfile` - Container configuration
- `Procfile` - Process configuration
- `.dockerignore` - Docker ignore file
- `deploy.sh` - Deployment script

## Health Check
Aplikasi memiliki health check endpoint di `/health` yang akan mengembalikan:
```json
{
  "status": "ok",
  "message": "BrainQuiz API is running",
  "version": "1.0.0"
}
```
