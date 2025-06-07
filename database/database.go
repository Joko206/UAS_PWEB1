package database

import (
	"fmt"
	"github.com/Joko206/UAS_PWEB1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "ballast.proxy.rlwy.net"
	port     = 46530
	user     = "postgres"
	password = "rQeJuOyOlkKxlVKfHYvhUkAoboRYCrWG"
	dbname   = "railway"
)

// Dsn contains the Data Source Name for PostgreSQL connection
var Dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, user, password, dbname)
var DB *gorm.DB

// Fungsi untuk menginisialisasi koneksi database
func InitDB() (*gorm.DB, error) {
	// Open a new connection to the database
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
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
		&models.Hasil_Kuis{},
		&models.SoalAnswer{},
		&models.Kelas_Pengguna{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}

// Helper function to get DB connection
func GetDBConnection() (*gorm.DB, error) {
	// If DB is already initialized, reuse it, otherwise initialize a new connection
	if DB == nil {
		db, err := InitDB()
		if err != nil {
			return nil, err
		}
		DB = db
	}
	return DB, nil
}
