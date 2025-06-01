package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null;index"`
	Email    string `json:"email" gorm:"unique;not null;index"`
	Password []byte `json:"-" gorm:"not null"` // Hide password in JSON responses
	Role     string `json:"role" gorm:"not null;default:'student';index"`
}

type Kategori_Soal struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null;index"`
	Description string `json:"description"`
}

type Tingkatan struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null;index"`
	Description string `json:"description"`
}

type Kelas struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null;index"`
	Description string `json:"description"`
}
type Kuis struct {
	gorm.Model
	Title        string        `json:"title" gorm:"not null;index"`
	Description  string        `json:"description"`
	KategoriID   uint          `json:"kategori_id" gorm:"not null;index"`
	Kategori     Kategori_Soal `gorm:"foreignKey:KategoriID;constraint:OnDelete:CASCADE"`
	TingkatanID  uint          `json:"tingkatan_id" gorm:"not null;index"`
	Tingkatan    Tingkatan     `gorm:"foreignKey:TingkatanID;constraint:OnDelete:CASCADE"`
	KelasID      uint          `json:"kelas_id" gorm:"not null;index"`
	Kelas        Kelas         `gorm:"foreignKey:KelasID;constraint:OnDelete:CASCADE"`
	PendidikanID uint          `json:"pendidikan_id" gorm:"not null;index"`
	Pendidikan   Pendidikan    `gorm:"foreignKey:PendidikanID;constraint:OnDelete:CASCADE"`
}

type Soal struct {
	gorm.Model
	Question      string          `json:"question" gorm:"not null"`
	Options       json.RawMessage `json:"options" gorm:"type:jsonb"`
	CorrectAnswer string          `json:"correct_answer" gorm:"not null"`
	KuisID        uint            `json:"kuis_id" gorm:"not null;index"`
	Kuis          Kuis            `gorm:"foreignKey:KuisID;constraint:OnDelete:CASCADE"`
}

type Pendidikan struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null;index"`
	Description string `json:"description"`
}

type HasilKuis struct {
	gorm.Model
	UserID        uint  `json:"user_id" gorm:"not null;index"`
	User          Users `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	KuisID        uint  `json:"kuis_id" gorm:"not null;index"`
	Kuis          Kuis  `gorm:"foreignKey:KuisID;constraint:OnDelete:CASCADE"`
	Score         uint  `json:"score" gorm:"not null"`
	CorrectAnswer uint  `json:"correct_answer" gorm:"not null"`
}

type SoalAnswer struct {
	gorm.Model
	SoalID uint   `json:"soal_id" gorm:"not null;index"`
	Soal   Soal   `gorm:"foreignKey:SoalID;constraint:OnDelete:CASCADE"`
	Answer string `json:"answer" gorm:"not null"`
	UserID uint   `json:"user_id" gorm:"not null;index"`
	User   Users  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type KelasPengguna struct {
	gorm.Model
	UserID  uint  `json:"user_id" gorm:"not null;index"`
	User    Users `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	KelasID uint  `json:"kelas_id" gorm:"not null;index"`
	Kelas   Kelas `gorm:"foreignKey:KelasID;constraint:OnDelete:CASCADE"`
}
