package models

import (
	"database/sql"
	"strings"
)

// CreateTableAdmins membuat tabel admin
func CreateTableAdmins(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS admins (
        email TEXT PRIMARY KEY
    );`
	_, err := db.Exec(query)
	return err
}

// Tambahkan data admin (convert to lowercase)
func InsertAdmin(db *sql.DB, email string) error {
	email = strings.ToLower(email)
	_, err := db.Exec("INSERT INTO admins (email) VALUES ($1)", email)
	return err
}

// Cek apakah email admin (case insensitive)
func IsAdmin(db *sql.DB, email string) (bool, error) {
	var count int
	// Convert to lowercase for case-insensitive comparison
	email = strings.ToLower(email)
	err := db.QueryRow("SELECT COUNT(*) FROM admins WHERE LOWER(email)=$1", email).Scan(&count)
	return (count > 0), err
}

// DeleteAdmin menghapus admin berdasarkan email
func DeleteAdmin(db *sql.DB, email string) error {
	_, err := db.Exec("DELETE FROM admins WHERE email=$1", email)
	return err
}
