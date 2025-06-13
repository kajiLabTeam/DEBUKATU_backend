package lib

import (
	"fmt"
	"kajiLabTeam/DEBUKATU/model"
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	once sync.Once
	DB   *gorm.DB
)

func InitDB() *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"), // 例: localhost:3306
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DATABASE"),
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // ← 追加
		})
		if err != nil {
			log.Fatalf("DB connect error: %v", err)
		}
		db.AutoMigrate(&model.User{}, &model.Model{})
		DB = db
	})
	return DB
}
