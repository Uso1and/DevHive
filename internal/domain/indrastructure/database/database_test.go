package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInitDB(t *testing.T) {
	t.Run("invalid config", func(t *testing.T) {
		// Сохраняем оригинальный DB и восстанавливаем после теста
		oldDB := DB
		defer func() { DB = oldDB }()

		err := InitDB()
		if err == nil {
			t.Error("Expected error with invalid config, got nil")
		}
	})
}

func TestCloseDB(t *testing.T) {
	t.Run("nil db", func(t *testing.T) {
		// Сохраняем оригинальный DB и восстанавливаем после теста
		oldDB := DB
		defer func() { DB = oldDB }()

		DB = nil
		if err := CloseDB(); err != nil {
			t.Errorf("CloseDB() should return nil when DB is nil, got %v", err)
		}
	})

	t.Run("valid db", func(t *testing.T) {
		// Создаем mock DB
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Failed to create sqlmock: %v", err)
		}
		defer db.Close()

		// Настраиваем ожидание закрытия соединения
		mock.ExpectClose()

		// Сохраняем оригинальный DB и восстанавливаем после теста
		oldDB := DB
		defer func() { DB = oldDB }()

		DB = db
		if err := CloseDB(); err != nil {
			t.Errorf("CloseDB() should return nil when DB is valid, got %v", err)
		}

		// Проверяем, что все ожидания выполнены
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %s", err)
		}
	})
}
