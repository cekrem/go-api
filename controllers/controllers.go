package controllers

import (  
  "net/http"
  "gopkg.in/gin-gonic/gin.v1"
)

type User struct {
	Name		string	`json: "name"`
	Password	string	`json: "password"`
}

func Index (c *gin.Context) {
	c.String(http.StatusOK, "Hello world");
}

func Name (c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello " + name);
}

func Json (c *gin.Context) {
	var json User
	if c.BindJSON(&json) == nil {
		c.JSON(http.StatusOK, json)
	} else {
		c.String(500, "NOT OK!")
	}
}