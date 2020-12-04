package connection

import (
	"github.com/cipent1112/majoo/model"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/majoo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(model.User{})
	return db
}
