package models

type Department struct {
	ID          uint   `gorm:"primaryKey" json:"department_id"`
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
}
