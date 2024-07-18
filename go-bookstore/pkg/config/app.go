package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)
//to connect directly to mysql by cmd; use command: mysql -u anhdev -p anhDB
//then type your password
func Connect() {
	d, err := gorm.Open("mysql", "anhdev:30122004AAbb@@@/anhDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
