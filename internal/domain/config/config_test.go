package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Создаем временный .env файл
	content := `
DB_HOST=test_host
DB_PORT=5433
DB_USER=test_user
DB_PASSWORD=test_pass
DB_NAME=test_db
DB_SSLMode=require
`
	// Создаем временную директорию
	tempDir := t.TempDir()
	envPath := filepath.Join(tempDir, ".env")

	// Записываем .env файл
	err := os.WriteFile(envPath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create .env file: %v", err)
	}

	// Меняем текущую директорию на временную
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(oldWd)

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	tests := []struct {
		name    string
		want    *ConfigDB
		wantErr bool
	}{
		{
			name: "successful load",
			want: &ConfigDB{
				Host:     "test_host",
				Port:     "5433",
				User:     "test_user",
				Password: "test_pass",
				Name:     "test_db",
				SSLMode:  "require",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && !tt.wantErr {
				t.Error("LoadConfig() returned nil config without error")
				return
			}
			if got != nil {
				if got.Host != tt.want.Host ||
					got.Port != tt.want.Port ||
					got.User != tt.want.User ||
					got.Password != tt.want.Password ||
					got.Name != tt.want.Name ||
					got.SSLMode != tt.want.SSLMode {
					t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
