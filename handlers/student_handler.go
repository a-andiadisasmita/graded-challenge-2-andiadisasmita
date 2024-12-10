package handlers

import (
	"graded-challenge-2-andiadisasmita/database"
	"graded-challenge-2-andiadisasmita/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Register a new student
func RegisterStudent(c *gin.Context) {
	var input struct {
		FirstName   string `json:"first_name" binding:"required"`
		LastName    string `json:"last_name" binding:"required"`
		Email       string `json:"email" binding:"required,email"`
		Address     string `json:"address" binding:"required"`
		Password    string `json:"password" binding:"required"`
		DateOfBirth string `json:"date_of_birth" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// Parse DateOfBirth to time.Time
	dateOfBirth, err := time.Parse("2006-01-02", input.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid date format"})
		return
	}

	db := database.GetDB()

	// Check if email already exists
	var existingStudent models.Student
	if err := db.Where("email = ?", input.Email).First(&existingStudent).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
		return
	}

	// Create a new student
	student := models.Student{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		Address:     input.Address,
		Password:    input.Password, // Ideally, hash the password before saving
		DateOfBirth: dateOfBirth,
	}

	if err := db.Create(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register student"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Student registered successfully",
		"data": gin.H{
			"first_name":    student.FirstName,
			"last_name":     student.LastName,
			"email":         student.Email,
			"address":       student.Address,
			"date_of_birth": student.DateOfBirth.Format("2006-01-02"),
		},
	})
}

// Login a student and generate JWT token
func LoginStudent(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	db := database.GetDB()

	var student models.Student
	if err := db.Where("email = ?", input.Email).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
		return
	}

	if student.Password != input.Password { // Ideally, verify hashed passwords
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid password"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"student_id": student.ID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Get currently logged-in student details
func GetStudentDetails(c *gin.Context) {
	db := database.GetDB()
	studentID := uint(c.GetInt("student_id"))

	var student models.Student
	if err := db.Preload("Enrollments.Course").First(&student, studentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"first_name":    student.FirstName,
		"last_name":     student.LastName,
		"email":         student.Email,
		"address":       student.Address,
		"date_of_birth": student.DateOfBirth.Format("2006-01-02"),
		"enrollments":   student.Enrollments,
	})
}
