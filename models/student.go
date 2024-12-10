package models

import "time"

type Student struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	FirstName   string       `gorm:"not null" json:"first_name"`
	LastName    string       `gorm:"not null" json:"last_name"`
	Email       string       `gorm:"not null;unique" json:"email"`
	Address     string       `gorm:"not null" json:"address"`
	Password    string       `gorm:"not null" json:"-"`
	DateOfBirth time.Time    `gorm:"not null" json:"date_of_birth"`
	Enrollments []Enrollment `gorm:"foreignKey:StudentID" json:"enrollments"`
}
