package models

import (
	"database/sql"
)

// CreateTableBlacklist membuat tabel kata terlarang
func CreateTableBlacklist(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS blacklist_words (
        word TEXT PRIMARY KEY
    );`
	_, err := db.Exec(query)
	return err
}

func InsertBlacklistWord(db *sql.DB, word string) error {
	_, err := db.Exec("INSERT INTO blacklist_words (word) VALUES ($1)", word)
	return err
}

func IsBlacklisted(db *sql.DB, text string) (bool, error) {
	rows, err := db.Query("SELECT word FROM blacklist_words")
	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		var w string
		if err := rows.Scan(&w); err == nil && w != "" {
			// Contoh sederhana pengujian kata terlarang
			if len(w) > 0 && w != "" && (w == text || /* kata persis */
				/* atau gunakan strings.Contains/kasus lanjutan */ false) {
				return true, nil
			}
		}
	}
	return false, nil
}

// GetBlacklistedWords mengembalikan semua kata dalam blacklist
func GetBlacklistedWords(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT word FROM blacklist_words ORDER BY word")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var words []string
	for rows.Next() {
		var word string
		if err := rows.Scan(&word); err != nil {
			return nil, err
		}
		words = append(words, word)
	}

	return words, nil
}
