package database

import (
	"fmt"
	"log"

	"belajar-via-dev.to/models"
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

var dsn string = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, user, password, dbname)

var DB *gorm.DB

func DBconn() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	db.AutoMigrate(&models.Users{}, &models.Kategori_Soal{}, &models.Tingkatan{}, models.Kelas{}, models.Kuis{}, models.Soal{})
}
