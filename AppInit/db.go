package AppInit

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"web/models"
)

var (
	db *gorm.DB
)

func init() {
	var err error

	db, err = gorm.Open(driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, database))
	if err!=nil {
		log.Fatalln(err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	db.DB().SetConnMaxLifetime(15)

	db.LogMode(true)

	migrate()
}

func migrate() {
	db.AutoMigrate(&models.Book{})
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return db
}