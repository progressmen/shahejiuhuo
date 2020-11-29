package main

import (
    "github.com/gin-gonic/gin"
    "shahejiuhuo/handle"
)

func main() {
    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    router.POST("/shjh/getList", handle.GetList)
    router.Run(":8877") // listen and serve on 0.0.0.0:8080
}