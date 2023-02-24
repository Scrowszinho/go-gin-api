package routes

import (
	"github.com/Scrowszinho/go-gin-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowStudents)
	r.Run()
}
