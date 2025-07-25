package models

import (
	"database/sql"
	"log"
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
	originalEmail := email
	email = strings.ToLower(email)
	log.Printf("DEBUG InsertAdmin: Storing email '%s' as '%s'", originalEmail, email)
	_, err := db.Exec("INSERT INTO admins (email) VALUES ($1)", email)
	return err
}

// Cek apakah email admin (case insensitive)
func IsAdmin(db *sql.DB, email string) (bool, error) {
	var count int
	var storedEmails []string
	
	// Debug: cek semua email yang ada di database
	rows, err := db.Query("SELECT email FROM admins")
	if err != nil {
		return false, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var storedEmail string
		if err := rows.Scan(&storedEmail); err != nil {
			continue
		}
		storedEmails = append(storedEmails, storedEmail)
	}
	
	log.Printf("DEBUG IsAdmin: Checking email '%s' against stored emails: %v", email, storedEmails)
	
	// Convert to lowercase for case-insensitive comparison
	email = strings.ToLower(email)
	err = db.QueryRow("SELECT COUNT(*) FROM admins WHERE LOWER(email)=$1", email).Scan(&count)
	
	log.Printf("DEBUG IsAdmin: Email '%s' (lowercase) found %d matches", email, count)
	return (count > 0), err
}

// DeleteAdmin menghapus admin berdasarkan email
func DeleteAdmin(db *sql.DB, email string) error {
	_, err := db.Exec("DELETE FROM admins WHERE email=$1", email)
	return err
}
