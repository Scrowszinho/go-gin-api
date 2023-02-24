package main

import (
	"github.com/gin-gonic/gin"
)

func showStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Gustavo",
	})
}

func main() {
	r := gin.Default()
	r.GET("/students", showStudents)
	r.Run()
}
