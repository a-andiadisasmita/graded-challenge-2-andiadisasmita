package models

type Professor struct {
	ID        uint   `gorm:"primaryKey" json:"professor_id"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"unique;not null" json:"email"`
	Address   string `gorm:"not null" json:"address"`
}
