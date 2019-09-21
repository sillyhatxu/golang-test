package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func main() {
	fmt.Println("-----------------------------------")
	fmt.Println(time.Now().Nanosecond())
	fmt.Println("-----------------------------------")
	for _, e := range os.Environ() {
		fmt.Println(e)
		//pair := strings.Split(e, "=")
		//fmt.Println(pair[0])
	}
	fmt.Println("-----------------------------------")
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "UP", "message": "OK"}) })
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
