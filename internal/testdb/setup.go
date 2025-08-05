package testdb

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/joho/godotenv"
	"github.com/oogway93/golangArchitecture/internal/adapter/config"
	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres"
	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	testDB     *gorm.DB
	testDBOnce sync.Once
)

func WithTestDB(t *testing.T) *gorm.DB {
	// Инициализируем соединение с БД один раз
	testDBOnce.Do(func() {
		cfg := loadTestDBConfig(t)

		var err error
		testDB, err = repositoryPostgres.DatabaseConnection(repositoryPostgres.Config{
			Username: cfg.User,
			Password: cfg.Password,
			Host:     cfg.Host,
			Port:     cfg.Port,
			DBName:   cfg.Name,
			SSLMode:  cfg.SSLMode,
		})

		if err != nil {
			t.Fatalf("Failed to connect to test database: %v", err)
		}

		// Применяем миграции
		if err := RunMigrations(testDB); err != nil {
			t.Fatalf("Failed to run migrations: %v", err)
		}
	})

	// Включаем логирование SQL для дебаггинга
	testDB.Logger = logger.Default.LogMode(logger.Info)

	// Возвращаем оригинальное соединение, НЕ транзакцию
	return testDB
}

func loadTestDBConfig(t *testing.T) config.TEST_DB {
	// Определяем возможные пути к .env файлам
	envFile := ".env.test"

	// Пытаемся загрузить каждый файл
	envPath := getAbsolutePath(envFile)
	if _, err := os.Stat(envPath); err == nil {
		if err := godotenv.Load(envPath); err != nil {
			t.Logf("Warning: failed to load %s: %v", envPath, err)
		} else {
			t.Logf("Loaded environment variables from %s", envPath)
		}
	}

	// Возвращаем конфигурацию с значениями по умолчанию
	return config.TEST_DB{
		Host:     getEnv("TEST_DB_HOST", "localhost"),
		Name:     getEnv("TEST_DB_NAME", "testdb"),
		Password: getEnv("TEST_DB_PASSWORD", "postgres"),
		Port:     getEnv("TEST_DB_PORT", "5432"),
		User:     getEnv("TEST_DB_USERNAME", "postgres"),
		SSLMode:  getEnv("TEST_DB_SSLMODE", "disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getAbsolutePath(relativePath string) string {
	// Получаем абсолютный путь к проекту
	projectRoot := os.Getenv("PROJECT_ROOT")
	if projectRoot == "" {
		// Пытаемся определить корень проекта автоматически
		cwd, err := os.Getwd()
		if err != nil {
			return relativePath
		}

		// Поднимаемся на 3 уровня вверх от internal/testutils
		projectRoot = filepath.Join(cwd, "..", "..", "..")
	}

	return filepath.Join(projectRoot, relativePath)
}

// RunMigrations применяет миграции
func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		// Перечислите все ваши модели
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
		&models.Delivery{},
	)
}
