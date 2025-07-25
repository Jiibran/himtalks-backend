# Manual Admin Setup Guide

## 🔧 Setup Database Tables

Jalankan SQL berikut di database PostgreSQL Anda:

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

-- Add default config
INSERT INTO configs (key, value) VALUES ('songfess_days', '7') 
ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value;
```

## 👥 Add First Admin User

Jalankan SQL berikut untuk menambahkan admin pertama:

```sql
-- Replace 'your-email@gmail.com' with your actual Google email
INSERT INTO admins (email) VALUES ('your-email@gmail.com');
```

**⚠️ Important**: Gunakan email yang sama dengan yang Anda gunakan untuk login Google OAuth.

## 🔍 Verify Setup

Check jika setup berhasil:

```sql
-- Check if tables exist
\dt

-- Check admin users
SELECT * FROM admins;

-- Check configs
SELECT * FROM configs;
```

## 🚀 After Setup

1. **Login**: Go to `https://api.teknohive.me/auth/google/login`
2. **Test Auth**: Check `https://api.teknohive.me/api/protected`
3. **Test Admin**: Check `https://api.teknohive.me/api/admin/list`

## 📋 Admin Management Endpoints

Once you're logged in as admin, you can manage other admins:

### Get Current Admins
```http
GET /api/admin/list
```

### Add New Admin
```http
POST /api/admin/addAdmin
Content-Type: application/json

{
  "Email": "newadmin@example.com"
}
```

### Remove Admin
```http
POST /api/admin/removeAdmin
Content-Type: application/json

{
  "Email": "admin@example.com"
}
```

**Note**: Admin cannot remove themselves.
