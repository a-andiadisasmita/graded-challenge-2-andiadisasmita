package models

type Course struct {
	ID           uint   `gorm:"primaryKey" json:"course_id"`
	Name         string `gorm:"not null" json:"name"`
	Description  string `gorm:"not null" json:"description"`
	DepartmentID uint   `gorm:"not null" json:"department_id"`
	Credits      int    `gorm:"not null" json:"credits"`
}
