package controllers

import (
	"gopkg.in/gin-gonic/gin.v1"
	"gitlab/christianek/go-api/models"
)

func Index(c *gin.Context) {
	c.String(200, "Hello world");
}

func Name(c *gin.Context) {
	name := c.Param("name")
	c.String(200, "Hello "+name);
}

func Json(c *gin.Context) {
	var json models.User
	if c.Bind(&json) == nil {
		c.JSON(200, json)
	} else {
		c.String(500, "NOT OK!")
	}
}
