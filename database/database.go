package database

import (
	"fmt"
	"log"

	"github.com/Joko206/UAS_PWEB1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "metro.proxy.rlwy.net"
	port     = 11951
	user     = "postgres"
	password = "VxYgKiPnPDgILDlzcYAxXOzEdOTUQxwh"
	dbname   = "railway"
)

var Dsn = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, user, password, dbname)

var DB *gorm.DB

func DBconn() {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	DB = db

	db.AutoMigrate(&models.Users{}, &models.Kategori_Soal{}, &models.Tingkatan{}, models.Kelas{}, models.Kuis{}, models.Soal{}, models.Pendidikan{})
}
