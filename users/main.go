package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("hello users")

	app := gin.Default()

	app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	app.GET("/check", func(c *gin.Context) {
		client := http.Client{}
		_, err := client.Get("http://orders/health")
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(500)
			return
		}
		c.Status(200)
	})

	srv := http.Server{
		Addr:              ":80",
		Handler:           app,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	srv.ListenAndServe()
}
