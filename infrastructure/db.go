package infrastructure

import (
	"fmt"
	"log"
	"os"
	"ziyaremtestcase/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// db bağalntısı için
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" {
		user = "username"
	}
	if pass == "" {
		pass = "password"
	}
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "3306"
	}
	if name == "" {
		name = "ziyarem"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("veri tabanına bağlanamadı:", err)
	}

	// objectleri dbye yazmak için
	if err := db.AutoMigrate(
		&domain.SensorType{},
		&domain.SensorDevice{},
		&domain.SensorAPI{},
		&domain.SensorData{},
	); err != nil {
		log.Fatal("migrasyon başarısız:", err)
	}

	log.Println(" veri tabanına bağlandı, migrasyon başarılı")
	return db
}
