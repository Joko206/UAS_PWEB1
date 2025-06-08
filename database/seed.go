package database

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Joko206/UAS_PWEB1/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedDatabase populates all tables with sample data
func SeedDatabase() error {
	log.Println("Attempting to get database connection...")
	db, err := GetDBConnection()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %w", err)
	}
	log.Println("Database connection successful!")

	log.Println("Starting database seeding...")

	// 1. Seed Users
	if err := seedUsers(db); err != nil {
		return fmt.Errorf("failed to seed users: %w", err)
	}

	// 2. Seed Kategori_Soal
	if err := seedKategoriSoal(db); err != nil {
		return fmt.Errorf("failed to seed kategori_soal: %w", err)
	}

	// 3. Seed Tingkatan
	if err := seedTingkatan(db); err != nil {
		return fmt.Errorf("failed to seed tingkatan: %w", err)
	}

	// 4. Seed Kelas
	if err := seedKelas(db); err != nil {
		return fmt.Errorf("failed to seed kelas: %w", err)
	}

	// 5. Seed Pendidikan
	if err := seedPendidikan(db); err != nil {
		return fmt.Errorf("failed to seed pendidikan: %w", err)
	}

	// 6. Seed Kuis
	if err := seedKuis(db); err != nil {
		return fmt.Errorf("failed to seed kuis: %w", err)
	}

	// 7. Seed Soal
	if err := seedSoal(db); err != nil {
		return fmt.Errorf("failed to seed soal: %w", err)
	}

	// 8. Seed Kelas_Pengguna
	if err := seedKelasPengguna(db); err != nil {
		return fmt.Errorf("failed to seed kelas_pengguna: %w", err)
	}

	// 9. Seed Hasil_Kuis
	if err := seedHasilKuis(db); err != nil {
		return fmt.Errorf("failed to seed hasil_kuis: %w", err)
	}

	// 10. Seed SoalAnswer
	if err := seedSoalAnswer(db); err != nil {
		return fmt.Errorf("failed to seed soal_answer: %w", err)
	}

	log.Println("Database seeding completed successfully!")
	return nil
}

func seedUsers(db *gorm.DB) error {
	log.Println("Seeding Users...")

	// Check if we have enough users (we need at least 20 for our seeding)
	var count int64
	db.Model(&models.Users{}).Count(&count)
	if count >= 20 {
		log.Printf("Already have %d users, skipping...", count)
		return nil
	}

	// Hash password for all users
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), 14)
	if err != nil {
		return err
	}

	users := []models.Users{
		{Name: "Admin User", Email: "admin@example.com", Password: hashedPassword, Role: "admin"},
		{Name: "Dr. Sarah Johnson", Email: "sarah.johnson@example.com", Password: hashedPassword, Role: "teacher"},
		{Name: "Prof. Michael Chen", Email: "michael.chen@example.com", Password: hashedPassword, Role: "teacher"},
		{Name: "Ms. Emily Davis", Email: "emily.davis@example.com", Password: hashedPassword, Role: "teacher"},
		{Name: "Mr. David Wilson", Email: "david.wilson@example.com", Password: hashedPassword, Role: "teacher"},
		{Name: "Alice Smith", Email: "alice.smith@example.com", Password: hashedPassword, Role: "student"},
		{Name: "Bob Brown", Email: "bob.brown@example.com", Password: hashedPassword, Role: "student"},
		{Name: "Charlie Green", Email: "charlie.green@example.com", Password: hashedPassword, Role: "student"},
		{Name: "Diana White", Email: "diana.white@example.com", Password: hashedPassword, Role: "student"},
		{Name: "Eva Black", Email: "eva.black@example.com", Password: hashedPassword, Role: "student"},
		{Name: "Frank Miller", Email: "frank.miller@example.com", Password: hashedPassword, Role: "student"},
		{Name: "Grace Taylor", Email: "grace.taylor@example.com", Password: hashedPassword, Role: "student"},
		{Name: "Henry Anderson", Email: "henry.anderson@example.com", Password: hashedPassword, Role: "student"},
		{Name: "Ivy Thomas", Email: "ivy.thomas@example.com", Password: hashedPassword, Role: "student"},
		{Name: "Jack Martinez", Email: "jack.martinez@example.com", Password: hashedPassword, Role: "student"},
		{Name: "Kelly Garcia", Email: "kelly.garcia@example.com", Password: hashedPassword, Role: "student"},
		{Name: "Leo Rodriguez", Email: "leo.rodriguez@example.com", Password: hashedPassword, Role: "student"},
		{Name: "Mia Lopez", Email: "mia.lopez@example.com", Password: hashedPassword, Role: "student"},
		{Name: "Noah Gonzalez", Email: "noah.gonzalez@example.com", Password: hashedPassword, Role: "student"},
		{Name: "Olivia Hernandez", Email: "olivia.hernandez@example.com", Password: hashedPassword, Role: "student"},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	log.Printf("Created %d users", len(users))
	return nil
}

func seedKategoriSoal(db *gorm.DB) error {
	log.Println("Seeding Kategori_Soal...")

	var count int64
	db.Model(&models.Kategori_Soal{}).Count(&count)
	if count > 0 {
		log.Println("Kategori_Soal already exist, skipping...")
		return nil
	}

	categories := []models.Kategori_Soal{
		{Name: "Matematika", Description: "Soal-soal matematika dasar hingga lanjutan"},
		{Name: "Bahasa Indonesia", Description: "Soal-soal bahasa Indonesia dan sastra"},
		{Name: "Bahasa Inggris", Description: "Soal-soal bahasa Inggris dan grammar"},
		{Name: "IPA", Description: "Ilmu Pengetahuan Alam - Fisika, Kimia, Biologi"},
		{Name: "IPS", Description: "Ilmu Pengetahuan Sosial - Sejarah, Geografi, Ekonomi"},
		{Name: "Komputer", Description: "Teknologi Informasi dan Komunikasi"},
		{Name: "Seni Budaya", Description: "Seni rupa, musik, dan budaya Indonesia"},
		{Name: "Olahraga", Description: "Pendidikan jasmani dan kesehatan"},
		{Name: "Agama", Description: "Pendidikan agama dan moral"},
		{Name: "PKN", Description: "Pendidikan Kewarganegaraan"},
	}

	for _, category := range categories {
		if err := db.Create(&category).Error; err != nil {
			return err
		}
	}

	log.Printf("Created %d categories", len(categories))
	return nil
}

func seedTingkatan(db *gorm.DB) error {
	log.Println("Seeding Tingkatan...")

	var count int64
	db.Model(&models.Tingkatan{}).Count(&count)
	if count > 0 {
		log.Println("Tingkatan already exist, skipping...")
		return nil
	}

	levels := []models.Tingkatan{
		{Name: "Mudah", Description: "Tingkat kesulitan mudah untuk pemula"},
		{Name: "Sedang", Description: "Tingkat kesulitan menengah"},
		{Name: "Sulit", Description: "Tingkat kesulitan tinggi untuk yang berpengalaman"},
		{Name: "Sangat Sulit", Description: "Tingkat kesulitan tertinggi untuk ahli"},
	}

	for _, level := range levels {
		if err := db.Create(&level).Error; err != nil {
			return err
		}
	}

	log.Printf("Created %d levels", len(levels))
	return nil
}

func seedKelas(db *gorm.DB) error {
	log.Println("Seeding Kelas...")

	var count int64
	db.Model(&models.Kelas{}).Count(&count)
	if count > 0 {
		log.Println("Kelas already exist, skipping...")
		return nil
	}

	classes := []models.Kelas{
		{Name: "Kelas 1", Description: "Kelas untuk siswa tingkat 1"},
		{Name: "Kelas 2", Description: "Kelas untuk siswa tingkat 2"},
		{Name: "Kelas 3", Description: "Kelas untuk siswa tingkat 3"},
		{Name: "Kelas 4", Description: "Kelas untuk siswa tingkat 4"},
		{Name: "Kelas 5", Description: "Kelas untuk siswa tingkat 5"},
		{Name: "Kelas 6", Description: "Kelas untuk siswa tingkat 6"},
		{Name: "Kelas 7", Description: "Kelas untuk siswa tingkat 7 (SMP)"},
		{Name: "Kelas 8", Description: "Kelas untuk siswa tingkat 8 (SMP)"},
		{Name: "Kelas 9", Description: "Kelas untuk siswa tingkat 9 (SMP)"},
		{Name: "Kelas 10", Description: "Kelas untuk siswa tingkat 10 (SMA)"},
		{Name: "Kelas 11", Description: "Kelas untuk siswa tingkat 11 (SMA)"},
		{Name: "Kelas 12", Description: "Kelas untuk siswa tingkat 12 (SMA)"},
	}

	for _, class := range classes {
		if err := db.Create(&class).Error; err != nil {
			return err
		}
	}

	log.Printf("Created %d classes", len(classes))
	return nil
}

func seedPendidikan(db *gorm.DB) error {
	log.Println("Seeding Pendidikan...")

	var count int64
	db.Model(&models.Pendidikan{}).Count(&count)
	if count > 0 {
		log.Println("Pendidikan already exist, skipping...")
		return nil
	}

	educations := []models.Pendidikan{
		{Name: "SD", Description: "Sekolah Dasar (Kelas 1-6)"},
		{Name: "SMP", Description: "Sekolah Menengah Pertama (Kelas 7-9)"},
		{Name: "SMA", Description: "Sekolah Menengah Atas (Kelas 10-12)"},
		{Name: "SMK", Description: "Sekolah Menengah Kejuruan (Kelas 10-12)"},
		{Name: "Universitas", Description: "Pendidikan tinggi"},
	}

	for _, education := range educations {
		if err := db.Create(&education).Error; err != nil {
			return err
		}
	}

	log.Printf("Created %d education levels", len(educations))
	return nil
}

func seedKuis(db *gorm.DB) error {
	log.Println("Seeding Kuis...")

	var count int64
	db.Model(&models.Kuis{}).Count(&count)
	if count > 0 {
		log.Println("Kuis already exist, skipping...")
		return nil
	}

	quizzes := []models.Kuis{
		{Title: "Matematika Dasar SD", Description: "Kuis matematika untuk siswa SD", Kategori_id: 1, Tingkatan_id: 1, Kelas_id: 1, Pendidikan_id: 1},
		{Title: "Bahasa Indonesia Kelas 2", Description: "Kuis bahasa Indonesia untuk kelas 2", Kategori_id: 2, Tingkatan_id: 1, Kelas_id: 2, Pendidikan_id: 1},
		{Title: "IPA Kelas 5", Description: "Kuis IPA untuk kelas 5 SD", Kategori_id: 4, Tingkatan_id: 2, Kelas_id: 5, Pendidikan_id: 1},
		{Title: "Matematika SMP", Description: "Kuis matematika untuk siswa SMP", Kategori_id: 1, Tingkatan_id: 2, Kelas_id: 7, Pendidikan_id: 2},
		{Title: "Bahasa Inggris SMP", Description: "Kuis bahasa Inggris untuk SMP", Kategori_id: 3, Tingkatan_id: 2, Kelas_id: 8, Pendidikan_id: 2},
		{Title: "IPS Kelas 9", Description: "Kuis IPS untuk kelas 9 SMP", Kategori_id: 5, Tingkatan_id: 3, Kelas_id: 9, Pendidikan_id: 2},
		{Title: "Fisika SMA", Description: "Kuis fisika untuk siswa SMA", Kategori_id: 4, Tingkatan_id: 3, Kelas_id: 10, Pendidikan_id: 3},
		{Title: "Kimia Lanjutan", Description: "Kuis kimia tingkat lanjutan", Kategori_id: 4, Tingkatan_id: 4, Kelas_id: 11, Pendidikan_id: 3},
		{Title: "Komputer Dasar", Description: "Kuis komputer untuk pemula", Kategori_id: 6, Tingkatan_id: 1, Kelas_id: 7, Pendidikan_id: 2},
		{Title: "Seni Budaya", Description: "Kuis seni dan budaya Indonesia", Kategori_id: 7, Tingkatan_id: 2, Kelas_id: 8, Pendidikan_id: 2},
		{Title: "PKN SMA", Description: "Kuis Pendidikan Kewarganegaraan SMA", Kategori_id: 10, Tingkatan_id: 2, Kelas_id: 10, Pendidikan_id: 3},
		{Title: "Olahraga dan Kesehatan", Description: "Kuis tentang olahraga dan kesehatan", Kategori_id: 8, Tingkatan_id: 1, Kelas_id: 6, Pendidikan_id: 1},
		{Title: "Agama Islam", Description: "Kuis pendidikan agama Islam", Kategori_id: 9, Tingkatan_id: 2, Kelas_id: 9, Pendidikan_id: 2},
		{Title: "Matematika Lanjutan SMA", Description: "Kuis matematika tingkat lanjutan untuk SMA", Kategori_id: 1, Tingkatan_id: 4, Kelas_id: 12, Pendidikan_id: 3},
		{Title: "Biologi SMA", Description: "Kuis biologi untuk siswa SMA", Kategori_id: 4, Tingkatan_id: 3, Kelas_id: 11, Pendidikan_id: 3},
	}

	for _, quiz := range quizzes {
		if err := db.Create(&quiz).Error; err != nil {
			return err
		}
	}

	log.Printf("Created %d quizzes", len(quizzes))
	return nil
}

func seedSoal(db *gorm.DB) error {
	log.Println("Seeding Soal...")

	var count int64
	db.Model(&models.Soal{}).Count(&count)
	if count > 0 {
		log.Println("Soal already exist, skipping...")
		return nil
	}

	// Sample questions for different quizzes
	questions := []struct {
		Question      string
		Options       map[string]string
		CorrectAnswer string
		KuisID        uint
	}{
		// Matematika Dasar SD (Kuis ID: 1)
		{
			Question:      "Berapa hasil dari 5 + 3?",
			Options:       map[string]string{"A": "6", "B": "7", "C": "8", "D": "9"},
			CorrectAnswer: "C",
			KuisID:        1,
		},
		{
			Question:      "Berapa hasil dari 10 - 4?",
			Options:       map[string]string{"A": "5", "B": "6", "C": "7", "D": "8"},
			CorrectAnswer: "B",
			KuisID:        1,
		},
		{
			Question:      "Berapa hasil dari 3 × 4?",
			Options:       map[string]string{"A": "10", "B": "11", "C": "12", "D": "13"},
			CorrectAnswer: "C",
			KuisID:        1,
		},
		// Bahasa Indonesia Kelas 2 (Kuis ID: 2)
		{
			Question:      "Apa sinonim dari kata 'besar'?",
			Options:       map[string]string{"A": "kecil", "B": "raksasa", "C": "sedang", "D": "tipis"},
			CorrectAnswer: "B",
			KuisID:        2,
		},
		{
			Question:      "Manakah yang merupakan kata benda?",
			Options:       map[string]string{"A": "lari", "B": "meja", "C": "cantik", "D": "cepat"},
			CorrectAnswer: "B",
			KuisID:        2,
		},
		// IPA Kelas 5 (Kuis ID: 3)
		{
			Question:      "Planet terdekat dengan matahari adalah?",
			Options:       map[string]string{"A": "Venus", "B": "Mars", "C": "Merkurius", "D": "Bumi"},
			CorrectAnswer: "C",
			KuisID:        3,
		},
		{
			Question:      "Proses perubahan air menjadi uap disebut?",
			Options:       map[string]string{"A": "kondensasi", "B": "evaporasi", "C": "sublimasi", "D": "kristalisasi"},
			CorrectAnswer: "B",
			KuisID:        3,
		},
		// Matematika SMP (Kuis ID: 4)
		{
			Question:      "Berapa hasil dari 2x + 3 = 11, nilai x adalah?",
			Options:       map[string]string{"A": "3", "B": "4", "C": "5", "D": "6"},
			CorrectAnswer: "B",
			KuisID:        4,
		},
		{
			Question:      "Luas lingkaran dengan jari-jari 7 cm adalah? (π = 22/7)",
			Options:       map[string]string{"A": "154 cm²", "B": "144 cm²", "C": "164 cm²", "D": "174 cm²"},
			CorrectAnswer: "A",
			KuisID:        4,
		},
		// Bahasa Inggris SMP (Kuis ID: 5)
		{
			Question:      "What is the past tense of 'go'?",
			Options:       map[string]string{"A": "goes", "B": "going", "C": "went", "D": "gone"},
			CorrectAnswer: "C",
			KuisID:        5,
		},
		{
			Question:      "Choose the correct sentence:",
			Options:       map[string]string{"A": "She are happy", "B": "She is happy", "C": "She am happy", "D": "She be happy"},
			CorrectAnswer: "B",
			KuisID:        5,
		},
	}

	for _, q := range questions {
		optionsJSON, err := json.Marshal(q.Options)
		if err != nil {
			return err
		}

		soal := models.Soal{
			Question:       q.Question,
			Options:        optionsJSON,
			Correct_answer: q.CorrectAnswer,
			Kuis_id:        q.KuisID,
		}

		if err := db.Create(&soal).Error; err != nil {
			return err
		}
	}

	log.Printf("Created %d questions", len(questions))
	return nil
}

func seedKelasPengguna(db *gorm.DB) error {
	log.Println("Seeding Kelas_Pengguna...")

	var count int64
	db.Model(&models.Kelas_Pengguna{}).Count(&count)
	if count > 0 {
		log.Println("Kelas_Pengguna already exist, skipping...")
		return nil
	}

	// Get all student users from database
	var students []models.Users
	if err := db.Where("role = ?", "student").Find(&students).Error; err != nil {
		return err
	}

	if len(students) < 6 {
		log.Printf("Not enough students found (%d), need at least 6 for class assignments", len(students))
		return nil
	}

	// Assign students to classes dynamically based on actual user IDs
	var classAssignments []models.Kelas_Pengguna

	// Assign first 6 students to elementary classes (SD)
	for i := 0; i < 6 && i < len(students); i++ {
		classAssignments = append(classAssignments, models.Kelas_Pengguna{
			Users_id: students[i].ID,
			Kelas_id: uint(i + 1), // Kelas 1-6
		})
	}

	// Assign next 3 students to middle school classes (SMP)
	for i := 6; i < 9 && i < len(students); i++ {
		classAssignments = append(classAssignments, models.Kelas_Pengguna{
			Users_id: students[i].ID,
			Kelas_id: uint(i + 1), // Kelas 7-9
		})
	}

	// Assign next 3 students to high school classes (SMA)
	for i := 9; i < 12 && i < len(students); i++ {
		classAssignments = append(classAssignments, models.Kelas_Pengguna{
			Users_id: students[i].ID,
			Kelas_id: uint(i + 1), // Kelas 10-12
		})
	}

	// Add some students to multiple classes if we have enough
	if len(students) > 12 {
		// Student 13 in multiple classes
		classAssignments = append(classAssignments,
			models.Kelas_Pengguna{Users_id: students[12].ID, Kelas_id: 7},
			models.Kelas_Pengguna{Users_id: students[12].ID, Kelas_id: 8},
		)
	}
	if len(students) > 13 {
		// Student 14 in multiple classes
		classAssignments = append(classAssignments,
			models.Kelas_Pengguna{Users_id: students[13].ID, Kelas_id: 9},
			models.Kelas_Pengguna{Users_id: students[13].ID, Kelas_id: 10},
		)
	}
	if len(students) > 14 {
		// Student 15 in multiple classes
		classAssignments = append(classAssignments,
			models.Kelas_Pengguna{Users_id: students[14].ID, Kelas_id: 11},
			models.Kelas_Pengguna{Users_id: students[14].ID, Kelas_id: 12},
		)
	}

	for _, assignment := range classAssignments {
		if err := db.Create(&assignment).Error; err != nil {
			return err
		}
	}

	log.Printf("Created %d class assignments", len(classAssignments))
	return nil
}

func seedHasilKuis(db *gorm.DB) error {
	log.Println("Seeding Hasil_Kuis...")

	var count int64
	db.Model(&models.Hasil_Kuis{}).Count(&count)
	if count > 0 {
		log.Println("Hasil_Kuis already exist, skipping...")
		return nil
	}

	// Get all student users from database
	var students []models.Users
	if err := db.Where("role = ?", "student").Find(&students).Error; err != nil {
		return err
	}

	if len(students) < 5 {
		log.Printf("Not enough students found (%d), need at least 5 for quiz results", len(students))
		return nil
	}

	// Create sample quiz results using actual student IDs
	var results []models.Hasil_Kuis

	// Basic results for first few students
	if len(students) > 0 {
		results = append(results, models.Hasil_Kuis{Users_id: students[0].ID, Kuis_id: 1, Score: 85, Correct_Answer: 3})
	}
	if len(students) > 1 {
		results = append(results, models.Hasil_Kuis{Users_id: students[1].ID, Kuis_id: 2, Score: 90, Correct_Answer: 2})
	}
	if len(students) > 2 {
		results = append(results, models.Hasil_Kuis{Users_id: students[2].ID, Kuis_id: 3, Score: 75, Correct_Answer: 2})
	}
	if len(students) > 3 {
		results = append(results, models.Hasil_Kuis{Users_id: students[3].ID, Kuis_id: 4, Score: 80, Correct_Answer: 2})
	}
	if len(students) > 4 {
		results = append(results, models.Hasil_Kuis{Users_id: students[4].ID, Kuis_id: 5, Score: 95, Correct_Answer: 2})
	}

	// Add more results if we have more students
	if len(students) > 5 {
		results = append(results,
			models.Hasil_Kuis{Users_id: students[5].ID, Kuis_id: 6, Score: 70, Correct_Answer: 1},
			models.Hasil_Kuis{Users_id: students[5].ID, Kuis_id: 7, Score: 88, Correct_Answer: 1},
		)
	}
	if len(students) > 6 {
		results = append(results,
			models.Hasil_Kuis{Users_id: students[6].ID, Kuis_id: 8, Score: 92, Correct_Answer: 1},
			models.Hasil_Kuis{Users_id: students[6].ID, Kuis_id: 9, Score: 85, Correct_Answer: 1},
		)
	}
	if len(students) > 7 {
		results = append(results,
			models.Hasil_Kuis{Users_id: students[7].ID, Kuis_id: 10, Score: 78, Correct_Answer: 1},
			models.Hasil_Kuis{Users_id: students[7].ID, Kuis_id: 11, Score: 82, Correct_Answer: 1},
		)
	}

	for _, result := range results {
		if err := db.Create(&result).Error; err != nil {
			return err
		}
	}

	log.Printf("Created %d quiz results", len(results))
	return nil
}

func seedSoalAnswer(db *gorm.DB) error {
	log.Println("Seeding SoalAnswer...")

	var count int64
	db.Model(&models.SoalAnswer{}).Count(&count)
	if count > 0 {
		log.Println("SoalAnswer already exist, skipping...")
		return nil
	}

	// Get all student users from database
	var students []models.Users
	if err := db.Where("role = ?", "student").Find(&students).Error; err != nil {
		return err
	}

	if len(students) < 3 {
		log.Printf("Not enough students found (%d), need at least 3 for answers", len(students))
		return nil
	}

	// Create sample answers using actual student IDs
	var answers []models.SoalAnswer

	// Student 1 answers for Matematika Dasar SD (Soal IDs 1-3)
	if len(students) > 0 {
		answers = append(answers,
			models.SoalAnswer{Soal_id: 1, Answer: "C", User_id: students[0].ID}, // Correct answer
			models.SoalAnswer{Soal_id: 2, Answer: "B", User_id: students[0].ID}, // Correct answer
			models.SoalAnswer{Soal_id: 3, Answer: "C", User_id: students[0].ID}, // Correct answer
		)
	}

	// Student 2 answers for Bahasa Indonesia Kelas 2 (Soal IDs 4-5)
	if len(students) > 1 {
		answers = append(answers,
			models.SoalAnswer{Soal_id: 4, Answer: "B", User_id: students[1].ID}, // Correct answer
			models.SoalAnswer{Soal_id: 5, Answer: "B", User_id: students[1].ID}, // Correct answer
		)
	}

	// Student 3 answers for IPA Kelas 5 (Soal IDs 6-7)
	if len(students) > 2 {
		answers = append(answers,
			models.SoalAnswer{Soal_id: 6, Answer: "C", User_id: students[2].ID}, // Correct answer
			models.SoalAnswer{Soal_id: 7, Answer: "A", User_id: students[2].ID}, // Wrong answer (correct is B)
		)
	}

	// Student 4 answers for Matematika SMP (Soal IDs 8-9)
	if len(students) > 3 {
		answers = append(answers,
			models.SoalAnswer{Soal_id: 8, Answer: "B", User_id: students[3].ID}, // Correct answer
			models.SoalAnswer{Soal_id: 9, Answer: "A", User_id: students[3].ID}, // Correct answer
		)
	}

	// Student 5 answers for Bahasa Inggris SMP (Soal IDs 10-11)
	if len(students) > 4 {
		answers = append(answers,
			models.SoalAnswer{Soal_id: 10, Answer: "C", User_id: students[4].ID}, // Correct answer
			models.SoalAnswer{Soal_id: 11, Answer: "B", User_id: students[4].ID}, // Correct answer
		)
	}

	// Additional answers from other students if available
	if len(students) > 5 {
		answers = append(answers,
			models.SoalAnswer{Soal_id: 1, Answer: "B", User_id: students[5].ID}, // Wrong answer
			models.SoalAnswer{Soal_id: 2, Answer: "B", User_id: students[5].ID}, // Correct answer
			models.SoalAnswer{Soal_id: 8, Answer: "B", User_id: students[5].ID}, // Correct answer
			models.SoalAnswer{Soal_id: 9, Answer: "B", User_id: students[5].ID}, // Wrong answer
		)
	}

	for _, answer := range answers {
		if err := db.Create(&answer).Error; err != nil {
			return err
		}
	}

	log.Printf("Created %d student answers", len(answers))
	return nil
}
