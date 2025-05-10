package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	id       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}
type Pendidikan struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Kategori_Soal struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Tingkatan struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Kelas struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Kuis struct {
	gorm.Model
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Kategori_id  uint          `json:"kategori_id"    `                                     // Gunakan tipe uint untuk Kategori_id
	Kategori     Kategori_Soal `gorm:"foreignKey:Kategori_id;constraint:OnDelete:CASCADE;"` //Menambahkan constraint foreign key
	Tingkatan_id uint          `json:"tingkatan_id"`
	Tingkatan    Tingkatan     `gorm:"foreignKey:Tingkatan_id;constraint:OnDelete:CASCADE;"`
	Kelas_id     uint          `json:"kelas_id"`
	Kelas        Kelas         `gorm:"foreignKey:Kelas_id;constraint:OnDelete:CASCADE;"`
}

type Soal struct {
	gorm.Model
	Question       string          `json:"question"`
	Options        json.RawMessage `json:"options_json"`
	Correct_answer string          `json:"correct_answer"`
	Kuis_id        uint            `json:"kuis_id"`
	Kuis           Kuis            `gorm:"foreignKey:Kuis_id"`
}
