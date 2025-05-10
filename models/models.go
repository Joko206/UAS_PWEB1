package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
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
	Title        string `json:"title"`
	Description  string `json:"description"`
	Kategori_id  int    `json:"kategori_id"`
	Tingkatan_id int    `json:"tingkatan_id"`
	Kelas_id     int    `json:"kelas_id"`
}
type Soal struct {
	gorm.Model
	Question       string          `json:"question"`
	Options        json.RawMessage `json:"options_json"`
	Correct_answer string          `json:"correct_answer"`
	Kuis_id        int             `json:"kuis_id"`
}
