package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lessongin/lessongin/src/models/mysql"
	"log"
	"net/http"
)

func creatUser(c *gin.Context) {

	var i mysql.UserInfo
	if err := c.BindJSON(&i); err != nil {
		return
	}
	u := mysql.UserInfoMode{
		Model: gorm.Model{},
		Name:  i.Name,
		Sex:   i.Sex,
		Phone: i.Phone,
		City:  i.City,
	}
	err := mysql.CreateTable()
	if err != nil {
		log.Printf("create user fail: %v", err)
	}
	c.JSON(http.StatusCreated, i)
}

func main() {
	_, err := mysql.ConnectMySQL()
	if err != nil {
		log.Fatalf("connect mysql err: %v", err)
	}

	log.Println("connect mysql success")

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	v1 := router.Group("/api/v1/users")
	{
		v1.POST("/", creatUser)
	}

}
