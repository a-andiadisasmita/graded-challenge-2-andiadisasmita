package models

type Teaching struct {
	ID          uint `gorm:"primaryKey" json:"teaching_id"`
	ProfessorID uint `gorm:"not null" json:"professor_id"`
	CourseID    uint `gorm:"not null" json:"course_id"`
}
