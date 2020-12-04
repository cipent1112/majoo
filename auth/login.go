package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cipent1112/majoo/connection"
	"github.com/cipent1112/majoo/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Login(c *gin.Context) {
	var db *gorm.DB = connection.DBConnect()
	client := model.User{
		Username: c.PostForm("username"),
		Password: c.PostForm("password")}

	fmt.Println("username ", client.Username)
	result := gin.H{}

	if err := db.Raw("SELECT * FROM user WHERE username = ? AND password = ?", client.Username, client.Password).First(&client); err != nil {
		result = gin.H{"result": "username / password salah"}
	}

	token, err := generateToken(strToken)
	if err != nil {
		log.Print(err)
	}
	result = gin.H{"token": token}

	c.JSON(http.StatusOK, result)
}
