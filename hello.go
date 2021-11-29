package main

import "github.com/gin-gonic/gin"

func main() {
 r := gin.Default()
 r.GET("/ping", func(c *gin.Context) {
     c.JSON(200, gin.H{
         "message": "hello huoshan 20211022 and test 55555 站内信通知44443333 !!!",
     })
 })
 r.Run() // listen and serve on 0.0.0.0:8000 (for windows "localhost:8080")
}
