package database

import (
	"database/sql"
	"devhive/internal/domain/config"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	dbconfig, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	configStr := dbconfig.GetConnectionString()

	DB, err = sql.Open("postgres", configStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to database")

	if err := runMigrations(DB); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

func runMigrations(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create driver: %w", err)
	}

	absPath, err := filepath.Abs("internal/domain/indrastructure/migrations")
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	log.Println("Migrations absolute path:", absPath)

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("migrations directory does not exist: %s", absPath)
	}

	files, err := os.ReadDir(absPath)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}
	log.Printf("Found %d migration files", len(files))

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+absPath,
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return fmt.Errorf("failed to get migration version: %w", err)
	}

	log.Printf("Migrations applied successfully. Version: %d, dirty: %v", version, dirty)
	return nil
}
func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
