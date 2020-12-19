package main

import (
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl", gin.H{
			"title": "あけましておめでとうございます！",
		})
	})
	router.GET("/loading", func(c *gin.Context) {
		c.HTML(http.StatusOK, "loading.tmpl", gin.H{
			"title": "loading...",
		})
	})
	router.GET("/result", func(c *gin.Context) {
		time.Sleep(2 * time.Second)
		c.HTML(http.StatusOK, "result.tmpl", gin.H{
			"title": "大吉",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
