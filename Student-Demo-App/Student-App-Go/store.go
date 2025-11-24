package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	DB *sql.DB
}

func (s *Store) migrate() {
	query := `
        CREATE TABLE IF NOT EXISTS students (
            id INT AUTO_INCREMENT PRIMARY KEY,
            first_name VARCHAR(100) NOT NULL,
            last_name  VARCHAR(100) NOT NULL,
            email      VARCHAR(150) NOT NULL UNIQUE,
            age        INT NOT NULL,
            created_at DATETIME NOT NULL,
            updated_at DATETIME NOT NULL
        );
    `
	_, err := s.DB.Exec(query)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migrated successfully!")
}

func NewStore(cfg *Config) *Store {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	log.Println("Connected to MySQL successfully!")

	store := &Store{DB: db}
	store.migrate()

	return store

}
