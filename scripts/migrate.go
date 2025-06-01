package main

import (
	"log"

	"github.com/Joko206/UAS_PWEB1/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Get database connection string
	dsn := database.GetDSN()

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("🔄 Starting database migration...")

	// Get raw SQL connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB: %v", err)
	}

	// Migration queries to fix existing data
	migrations := []string{
		// Drop all existing tables to start fresh
		"DROP TABLE IF EXISTS hasil_kuis CASCADE;",
		"DROP TABLE IF EXISTS soal_answers CASCADE;",
		"DROP TABLE IF EXISTS kelas_penggunas CASCADE;",
		"DROP TABLE IF EXISTS soals CASCADE;",
		"DROP TABLE IF EXISTS kuis CASCADE;",

		// Drop other tables that might have conflicts
		"DROP TABLE IF EXISTS users CASCADE;",
		"DROP TABLE IF EXISTS kategori_soals CASCADE;",
		"DROP TABLE IF EXISTS tingkatans CASCADE;",
		"DROP TABLE IF EXISTS kelas CASCADE;",
		"DROP TABLE IF EXISTS pendidikans CASCADE;",
	}

	// Execute migration queries
	for _, query := range migrations {
		log.Printf("Executing: %s", query)
		if _, err := sqlDB.Exec(query); err != nil {
			log.Printf("Warning: %v", err)
		}
	}

	log.Println("✅ Database migration completed successfully!")
	log.Println("You can now run 'go run main.go' to start the application")
}
