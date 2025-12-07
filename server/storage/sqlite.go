package storage

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func InitDB(dbPath string) (*DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS secrets (
		id TEXT PRIMARY KEY,
		encrypted_data TEXT NOT NULL,
		created_at INTEGER NOT NULL
	);
	
	CREATE INDEX IF NOT EXISTS idx_created_at ON secrets(created_at);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	log.Println("Database initialized successfully")

	return &DB{db}, nil
}

func (db *DB) SaveSecret(id, encryptedData string) error {
	query := `INSERT INTO secrets (id, encrypted_data, created_at) VALUES (?, ?, ?)`
	_, err := db.Exec(query, id, encryptedData, time.Now().Unix())
	return err
}

func (db *DB) GetAndDeleteSecret(id string) (string, error) {
	var encryptedData string
	
	err := db.QueryRow("SELECT encrypted_data FROM secrets WHERE id = ?", id).Scan(&encryptedData)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", sql.ErrNoRows
		}
		return "", err
	}

	_, err = db.Exec("DELETE FROM secrets WHERE id = ?", id)
	if err != nil {
		return "", err
	}

	return encryptedData, nil
}

func (db *DB) CleanupOldSecrets(daysOld int) error {
	cutoffTime := time.Now().AddDate(0, 0, -daysOld).Unix()
	_, err := db.Exec("DELETE FROM secrets WHERE created_at < ?", cutoffTime)
	return err
}



