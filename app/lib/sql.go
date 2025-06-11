package lib

import (
	"fmt"
	"kajiLabTeam/DEBUKATU/model"
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once sync.Once
	DB   *gorm.DB
)

func InitDB() *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"), // 例: localhost:3306
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DATABASE"),
		)
		db, err := gorm.Open(mysql.Open(dsn))
		if err != nil {
			log.Fatalf("DB connect error: %v", err)
		}
		db.AutoMigrate(&model.User{})
		DB = db
	})
	return DB
}
