package models

import "time"

type Enrollment struct {
	ID             uint      `gorm:"primaryKey" json:"enrollment_id"`
	StudentID      uint      `gorm:"not null" json:"student_id"`
	CourseID       uint      `gorm:"not null" json:"course_id"`
	EnrollmentDate time.Time `gorm:"not null" json:"enrollment_date"`
}
