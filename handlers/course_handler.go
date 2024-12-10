package handlers

import (
	"graded-challenge-2-andiadisasmita/database"
	"graded-challenge-2-andiadisasmita/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all courses
func GetCourses(c *gin.Context) {
	db := database.GetDB()

	var courses []models.Course
	if err := db.Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve courses"})
		return
	}

	c.JSON(http.StatusOK, courses)
}
