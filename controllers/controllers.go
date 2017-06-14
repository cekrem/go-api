package controllers

import (  
  "gopkg.in/gin-gonic/gin.v1"
)

type User struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Index (c *gin.Context) {
	c.String(200, "Hello world");
}

func Name (c *gin.Context) {
	name := c.Param("name")
	c.String(200, "Hello " + name);
}

func Json (c *gin.Context) {
	var json User
	if c.Bind(&json) == nil {
		c.JSON(200, json)
	} else {
		c.String(500, "NOT OK!")
	}
}