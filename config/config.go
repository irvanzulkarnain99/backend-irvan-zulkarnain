package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

// LoadEnv akan membaca .env file saat aplikasi start
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Warning: .env file not found, using system ENV instead")
	}
}

// ConnectDB melakukan koneksi ke PostgreSQL menggunakan GORM
func ConnectDB() {
	once.Do(func() {
		var err error
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			GetEnv("DB_HOST", "localhost"),
			GetEnv("DB_USER", "postgres"),
			GetEnv("DB_PASSWORD", "123"),
			GetEnv("DB_NAME", "marketplace"),
			GetEnv("DB_PORT", "5432"),
		)

		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("❌ Failed to connect to database:", err)
		} else {
			log.Println("✅ Connected to PostgreSQL database")
		}
	})
}

// GetEnv membantu mengambil env variable dengan fallback default
func GetEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
