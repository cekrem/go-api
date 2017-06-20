package controllers

import (
	"github.com/cekrem/go-api/models"
	"gopkg.in/gin-gonic/gin.v1"
)

func Index(c *gin.Context) {
	c.String(200, "Hello world")
}

func Name(c *gin.Context) {
	name := c.Param("name")
	c.String(200, "Hello "+name)
}

func Json(c *gin.Context) {
	var json models.User
	if err := c.Bind(&json); err == nil {
		c.JSON(200, json)
	} else {
		c.JSON(400, err)
	}
}
