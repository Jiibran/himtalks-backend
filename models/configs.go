package models

import (
	"database/sql"
	"strconv"
)

func CreateTableConfigs(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS configs (
        key TEXT PRIMARY KEY,
        value TEXT NOT NULL
    );`
	_, err := db.Exec(query)
	return err
}

// Simpan config misalnya "songfess_days"
func SetConfig(db *sql.DB, key, value string) error {
	_, err := db.Exec(`
        INSERT INTO configs (key, value) VALUES ($1, $2)
        ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value
    `, key, value)
	return err
}

func GetSongfessDays(db *sql.DB) (int, error) {
	var stored string
	err := db.QueryRow("SELECT value FROM configs WHERE key='songfess_days'").Scan(&stored)
	if err != nil {
		return 7, nil // Default 7 apabila belum diatur
	}
	days, convErr := strconv.Atoi(stored)
	if convErr != nil {
		return 7, nil
	}
	return days, nil
}
