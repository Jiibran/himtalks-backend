# Admin Setup & Troubleshooting Guide

## 🚨 Masalah: Admin Endpoints Returns 500 Error

Jika semua endpoint admin mengembalikan error 500, kemungkinan penyebabnya adalah:

### 1. **Tabel Database Belum Dibuat**

Backend memerlukan tabel-tabel berikut:
- `admins` - Untuk menyimpan daftar admin users
- `configs` - Untuk konfigurasi sistem
- `blacklist_words` - Untuk kata-kata terlarang

**Solusi**: Restart aplikasi untuk membuat tabel otomatis, atau buat manual:

```sql
-- Create admin table
CREATE TABLE IF NOT EXISTS admins (
    email TEXT PRIMARY KEY
);

-- Create configs table
CREATE TABLE IF NOT EXISTS configs (
    key TEXT PRIMARY KEY,
    value TEXT NOT NULL
);

-- Create blacklist table
CREATE TABLE IF NOT EXISTS blacklist_words (
    word TEXT PRIMARY KEY
);
```

### 2. **Tidak Ada Admin User**

Setelah tabel dibuat, tidak ada admin user yang terdaftar.

**Solusi**: Tambahkan admin pertama menggunakan script:

```bash
# Windows PowerShell
go run add_admin.go your-email@example.com

# Linux/Mac
go run add_admin.go your-email@example.com
```

### 3. **Problem dengan JWT Token**

Admin endpoints memerlukan JWT token yang valid dari Google OAuth.

**Debug Authentication**:
1. Login dulu via `/auth/google/login`
2. Pastikan cookie `jwt` ter-set
3. Cek status auth dengan `/api/protected`

## 🔧 Debug Tools

### Check Database Status
```bash
go run debug_db.go
```

Output yang diharapkan:
```
=== Database Debug Information ===
Table messages exists: true
Table songfess exists: true
Table admins exists: true
Table configs exists: true
Table blacklist_words exists: true

=== Admin Users ===
Admin: your-email@example.com
Total admins: 1

Total messages: X
Total songfess: Y
```

### Test Admin Endpoints
```bash
# Test dengan cookie authentication
curl -b cookies.txt https://api.teknohive.me/api/admin/messages

# Check response
# - 401: Not authenticated (need to login first)
# - 403: Not admin (add user to admin table)
# - 500: Database/server error (check logs)
# - 200: Success
```

## 📋 Step-by-Step Fix

### Step 1: Restart Backend
Restart aplikasi untuk membuat tabel yang kurang:
```bash
# Stop current process
# Restart with: go run main.go
```

### Step 2: Add Admin User
```bash
go run add_admin.go your-google-email@gmail.com
```

### Step 3: Test Authentication
1. Login: `https://api.teknohive.me/auth/google/login`
2. Check: `https://api.teknohive.me/api/protected`
3. Should return your email

### Step 4: Test Admin Endpoint
```bash
curl -b cookies.txt https://api.teknohive.me/api/admin/messages
```

## 🔍 Common Issues

### Issue: "Table doesn't exist"
**Fix**: Restart aplikasi atau buat tabel manual

### Issue: "Not an admin" (403)
**Fix**: Tambahkan email ke tabel admin
```bash
go run add_admin.go your-email@example.com
```

### Issue: "Unauthorized" (401)
**Fix**: Login dulu via Google OAuth
- Go to: `https://api.teknohive.me/auth/google/login`
- Complete OAuth flow
- Try admin endpoint again

### Issue: JWT Cookie tidak ter-set
**Fix**: Check CORS dan cookie domain configuration di `main.go`

## 📊 Admin Endpoint Flow

```
Request → CORS Middleware → AuthMiddlewareAdmin → CheckIsAdmin → Controller
                                ↓                      ↓
                         Check JWT Cookie        Check if email in
                         Extract email           admin table
```

Setiap step bisa gagal:
1. **CORS**: Domain tidak diizinkan
2. **AuthMiddlewareAdmin**: JWT tidak valid/tidak ada
3. **CheckIsAdmin**: Email tidak ada di tabel admin
4. **Controller**: Error di business logic

## 🚀 Quick Fix Commands

```bash
# 1. Debug database
go run debug_db.go

# 2. Add admin (replace with your Google email)
go run add_admin.go your-email@gmail.com

# 3. Test endpoint
curl -b cookies.txt https://api.teknohive.me/api/protected

# 4. If successful, test admin endpoint
curl -b cookies.txt https://api.teknohive.me/api/admin/messages
```

## 📝 Logs to Check

Jika masih error, cek logs untuk:
```
AuthMiddlewareAdmin: processing request for /api/admin/songfess
AuthMiddlewareAdmin: JWT token found in cookie
AuthMiddlewareAdmin: valid JWT token for email: user@example.com
CheckIsAdmin: checking admin status for email: user@example.com
CheckIsAdmin: user user@example.com is admin, allowing access to /api/admin/songfess
```

Jika ada yang missing, itu adalah source masalahnya.
