package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	id       uint   `gorm:"primaryKey"`
	Name     string `json:"name" gorm:"column:name;not null"` // Pastikan kolom 'name' ada di database
	Email    string `json:"email" gorm:"unique;not null"`
	Password []byte `json:"-" gorm:"not null"`
	Role     string `json:"role"`
}
type Kategori_Soal struct {
	gorm.Model
	id          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Tingkatan struct {
	gorm.Model
	id          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Kelas struct {
	gorm.Model
	id          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Kuis struct {
	gorm.Model
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Kategori_id   uint          `json:"kategori_id"    `
	Kategori      Kategori_Soal `gorm:"foreignKey:Kategori_id;constraint:OnDelete:CASCADE;"`
	Tingkatan_id  uint          `json:"tingkatan_id"`
	Tingkatan     Tingkatan     `gorm:"foreignKey:Tingkatan_id;constraint:OnDelete:CASCADE;"`
	Kelas_id      uint          `json:"kelas_id"`
	Kelas         Kelas         `gorm:"foreignKey:Kelas_id;constraint:OnDelete:CASCADE;"`
	Pendidikan_id uint          `json:"pendidikan_id"`
	Pendidikan    Pendidikan    `gorm:"foreignKey:Pendidikan_id;constraint:OnDelete:CASCADE;"`
}

type Soal struct {
	gorm.Model
	Question       string          `json:"question"`
	Options        json.RawMessage `json:"options_json"`
	Correct_answer string          `json:"correct_answer"`
	Kuis_id        uint            `json:"kuis_id"`
	Kuis           Kuis            `gorm:"foreignKey:Kuis_id;constraint:OnDelete:CASCADE;"`
}

type Pendidikan struct {
	gorm.Model
	id          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Hasil_Kuis struct {
	gorm.Model
	Users_id       uint  `json:"users_id"`
	Users          Users `gorm:"foreignKey:Users_id;constraint:OnDelete:CASCADE;"`
	Kuis_id        uint  `json:"kuis_id"`
	Kuis           Kuis  `gorm:"foreignKey:Kuis_id;constraint:OnDelete:CASCADE;"`
	Score          uint  `json:"score"`
	Correct_Answer uint  `json:"correct_answer;constraint:OnDelete:CASCADE;"`
}
type SoalAnswer struct {
	gorm.Model
	Soal_id uint   `json:"soal_id"`
	Soal    Soal   `gorm:"foreignKey:Soal_id;constraint:OnDelete:CASCADE;"`
	Answer  string `json:"answer"`
	User_id uint   `json:"user_id"`
	User    Users  `gorm:"foreignKey:User_id;constraint:OnDelete:CASCADE;"`
}
type Kelas_Pengguna struct {
	gorm.Model
	Users_id uint  `json:"users_id"`
	Users    Users `gorm:"foreignKey:Users_id;constraint:OnDelete:CASCADE;"`
	Kelas_id uint  `json:"kelas_id"`
	Kelas    Kelas `gorm:"foreignKey:Kelas_id;constraint:OnDelete:CASCADE;"`
}
