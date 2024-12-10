package handlers

import (
	"graded-challenge-2-andiadisasmita/database"
	"graded-challenge-2-andiadisasmita/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Enroll in a course
func EnrollCourse(c *gin.Context) {
	db := database.GetDB()
	studentID := uint(c.GetInt("student_id"))

	var input struct {
		CourseID int `json:"course_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// Check if the student is already enrolled
	var existingEnrollment models.Enrollment
	if err := db.Where("student_id = ? AND course_id = ?", studentID, uint(input.CourseID)).First(&existingEnrollment).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Already enrolled in this course"})
		return
	}

	// Create a new enrollment
	enrollment := models.Enrollment{
		StudentID:      studentID,
		CourseID:       uint(input.CourseID),
		EnrollmentDate: time.Now(),
	}

	if err := db.Create(&enrollment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to enroll in course"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Enrollment successful",
		"data":    enrollment,
	})
}

// Delete an enrollment
func DeleteEnrollment(c *gin.Context) {
	db := database.GetDB()
	studentID := c.GetInt("student_id")
	enrollmentID := c.Param("id")

	var enrollment models.Enrollment
	if err := db.Where("id = ? AND student_id = ?", enrollmentID, studentID).First(&enrollment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Enrollment not found"})
		return
	}

	if err := db.Delete(&enrollment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete enrollment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Enrollment deleted successfully",
		"data":    enrollment,
	})
}
