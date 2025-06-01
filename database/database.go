package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Joko206/UAS_PWEB1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the global database connection
var DB *gorm.DB

// GetDSN returns the database connection string
func GetDSN() string {
	// Try to get from environment variables first
	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		return dsn
	}

	// Fallback to hardcoded values (for development only)
	host := getEnv("DB_HOST", "metro.proxy.rlwy.net")
	port := getEnv("DB_PORT", "11951")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "VxYgKiPnPDgILDlzcYAxXOzEdOTUQxwh")
	dbname := getEnv("DB_NAME", "railway")
	sslmode := getEnv("DB_SSLMODE", "disable")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		host, port, user, password, dbname, sslmode)
}

// getEnv gets environment variable with fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// InitDB initializes the database connection
func InitDB() (*gorm.DB, error) {
	dsn := GetDSN()

	// Configure GORM logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Enable color
		},
	)

	// Open database connection with configuration
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Run AutoMigrate to ensure the database schema is up to date
	if err := db.AutoMigrate(
		&models.Users{},
		&models.Kategori_Soal{},
		&models.Tingkatan{},
		&models.Kelas{},
		&models.Kuis{},
		&models.Soal{},
		&models.Pendidikan{},
		&models.HasilKuis{},
		&models.SoalAnswer{},
		&models.KelasPengguna{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("✅ Database connected and migrated successfully")
	return db, nil
}

// GetDBConnection returns the database connection
func GetDBConnection() (*gorm.DB, error) {
	if DB == nil {
		db, err := InitDB()
		if err != nil {
			return nil, err
		}
		DB = db
	}
	return DB, nil
}

// MustGetDB returns the database connection or panics
func MustGetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database not initialized. Call GetDBConnection() first.")
	}
	return DB
}
