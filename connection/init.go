package connection

import (
	"github.com/jinzhu/gorm"
)

func DBConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/majoo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}
	return db
}
