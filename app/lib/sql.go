package lib

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SqlConnect() (database *gorm.DB) {
	var db *gorm.DB
	var err error

	// 環境変数から値を取得
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	protocol := os.Getenv("MYSQL_PROTOCOL")
	dbname := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=true&loc=Asia%%2FTokyo", user, password, protocol, dbname)
	// GORMのDialectorを作成
	dialector := mysql.Open(dsn)

	// log.Default().Println(dsn)

	if db, err = gorm.Open(dialector); err != nil {
		db = connect(dialector, 10)
	}
	fmt.Println("db connected!!")

	return db
}

func connect(dialector gorm.Dialector, count uint) *gorm.DB {
	var err error
	var db *gorm.DB
	if db, err = gorm.Open(dialector); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			return connect(dialector, count)
		}
		panic(err.Error())
	}
	return (db)
}
