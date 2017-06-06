package main

import (  
  "gopkg.in/gin-gonic/gin.v1"
  "gitlab/christianek/go-api/controllers"
)

func main() {
  router:= gin.Default();

  v1 := router.Group("api/v1")
  {
    v1.GET("/", controllers.Index)
    v1.GET("/name/:name", controllers.Name)
    v1.POST("/json", controllers.Json)
  }
  router.Run(":3200");
}

