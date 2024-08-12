package config

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

// Global variable untuk menyimpan koneksi database
var db *sql.DB

// OpenConnection membuka koneksi ke database dengan konfigurasi
func OpenConnection(cfg Config) error {
	var err error
	db, err = setupConnection(cfg)
	if err != nil {
		return err
	}

	// Cek koneksi
	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

// setupConnection menyiapkan koneksi ke database
func setupConnection(cfg Config) (*sql.DB, error) {
	connection := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBHost, cfg.DBPort, cfg.SSLMode)
	fmt.Println("Connection Info:", cfg.DBDriver, connection)

	db, err := sql.Open(cfg.DBDriver, connection)
	if err != nil {
		return nil, errors.New("failed to create the database connection")
	}

	return db, nil
}

// CloseConnectionDB menutup koneksi ke database
func CloseConnectionDB() {
	if db != nil {
		db.Close()
	}
}

// DBConnection mengembalikan koneksi database
func DBConnection() *sql.DB {
	return db
}
