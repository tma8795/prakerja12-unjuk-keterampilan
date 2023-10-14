package configs

import (
	"ecommerce/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/prakerja-golang?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	InitMigration()
}

func InitMigration() {
	DB.AutoMigrate(&models.Products{})
}
