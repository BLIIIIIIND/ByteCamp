package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	UUID    string `gorm:"type:varchar(100); primaryKey"`
	Name    string `gorm:"type:varchar(50)"`
	Age     uint
	Version uint
}

func main() {
	db, err := gorm.Open(mysql.Open("root:zxc05020519@tcp(192.168.200.128:23306)/course?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
}
