package controllers

import (
	"log"
	"net/http"		
	
	"github.com/gin-gonic/gin"
)

func TestApi(c *gin.Context) {
	type user struct {
		Name   string   `json:"name"`
		Age    int      `json:"age"`
		Email  string   `json:"email"`
	}

	nut := user{
		Name: "Nuttawut",
		Age: 23,
		Email: "nut@mail.com",
	}

	log.Println(nut)

	c.JSON(http.StatusOK, nut)
}