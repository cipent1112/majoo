package main

import (
	"github.com/cipent1112/majoo/connection"
	"github.com/cipent1112/majoo/handler"
	"github.com/cipent1112/majoo/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	session := connection.DBConnect()
	db := &handler.DB{DB: session}
	router.Route(db)
}
