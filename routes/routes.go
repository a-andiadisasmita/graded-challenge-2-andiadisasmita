package routes

import (
	"graded-challenge-2-andiadisasmita/handlers"
	"graded-challenge-2-andiadisasmita/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/students/register", handlers.RegisterStudent)
	r.POST("/students/login", handlers.LoginStudent)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.JWTMiddleware())
	protected.GET("/students/me", handlers.GetStudentDetails)
	protected.GET("/courses", handlers.GetCourses)
	protected.POST("/enrollments", handlers.EnrollCourse)
	protected.DELETE("/enrollments/:id", handlers.DeleteEnrollment)

	return r
}
