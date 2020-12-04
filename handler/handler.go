package handler

import (
	"fmt"
	"net/http"

	"github.com/cipent1112/majoo/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type DB struct {
	*gorm.DB
}

func (db *DB) GetUser(c *gin.Context) {
	client := model.User{}
	result := gin.H{}

	id := c.Param("id")
	err := db.DB.Raw("SELECT * FROM user WHERE id = ? LIMIT 1", id).Find(&client).Error
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error : %s", err.Error()))
		return
	}
	result = gin.H{"result": client}

	c.JSON(http.StatusOK, result)
}

func (db *DB) GetUsers(c *gin.Context) {
	clients := []model.User{}
	total := 0
	db.DB.Table(model.TableName).Count(&total)
	if total <= 0 {
		c.String(http.StatusNotFound, "not found")
		return
	}

	if err := db.DB.Raw("SELECT * FROM user").Find(&clients).Error; err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error : %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, clients)
}

func (db *DB) CreateUser(c *gin.Context) {
	file, err := c.FormFile("foto")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error : %s", err.Error()))
		return
	}
	path := "foto/" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error : %s", err.Error()))
		return
	}

	client := model.User{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
		Nama:     c.PostForm("nama"),
		Foto:     path}

	err = model.ValidateUser(client)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error : %s", err.Error()))
		return
	}

	result := gin.H{}
	db.DB.Table(model.TableName).Create(&client)
	result = gin.H{
		"result": client,
	}
	c.JSON(http.StatusOK, result)
}

func (db *DB) UpdateUser(c *gin.Context) {
	client := model.User{}
	result := gin.H{}

	id := c.Query("id")
	err := db.DB.Table(model.TableName).First(&client, id).Error
	if err != nil {
		result = gin.H{"result": "data not found"}
	}

	file, err := c.FormFile("foto")
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error : %s", err.Error()))
		return
	}

	path := "foto/" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error : %s", err.Error()))
		return
	}

	newClient := model.User{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
		Nama:     c.PostForm("nama"),
		Foto:     path}

	err = model.ValidateUser(client)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error : %s", err.Error()))
		return
	}

	err = db.DB.Table(model.TableName).Model(&client).Updates(newClient).Error
	if err != nil {
		result = gin.H{"result": "update failed"}
		return
	}
	result = gin.H{"result": "successfully updated data"}

	c.JSON(http.StatusOK, result)
}

func (db *DB) DeleteUser(c *gin.Context) {
	client := model.User{}
	result := gin.H{}

	id := c.Param("id")
	if err := db.DB.Table(model.TableName).First(&client, id).Error; err != nil {
		result = gin.H{"result": "data not found"}
	}
	if err := db.DB.Table(model.TableName).Delete(&client).Error; err != nil {
		result = gin.H{"result": "delete failed"}
	} else {
		result = gin.H{"result": "Data deleted successfully"}
	}

	c.JSON(http.StatusOK, result)
}
