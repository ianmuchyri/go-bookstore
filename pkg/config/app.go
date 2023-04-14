package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//this comment illustrates the format of func open in gorm which initializes a new db connection
	/* import _ "github.com/go-sql-driver/mysql"
	func main() {
	  db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	}*/
	d, err := gorm.Open("mysql", "muchyri:12345678@/bookstore?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
