package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {

	d, err := gorm.Open(mysql.Open("vedant:123@tcp(localhost:3306)/simplerest?charset=utf8&parseTime=True&loc=Asia%2FKolkata"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d

}

func GetDB() *gorm.DB {

	if db == nil {
		Connect()
	}
	return db
}
