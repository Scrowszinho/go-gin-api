package controllers

import (
	"github.com/Scrowszinho/go-gin-api/models"
	"github.com/gin-gonic/gin"
)

func ShowStudents(c *gin.Context) {
	c.JSON(200, models.Alunos)
}
