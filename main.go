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

    // 获取分类
    router.POST("/shjh/getItemList", handle.GetItemList)

    // 获取图片
    router.POST("/shjh/getPicList", handle.GetPicList)

    // 初始化数据库
    handle.InitDb()

    // 启动服务
    router.Run(":8877") // listen and serve on 0.0.0.0:8080
}