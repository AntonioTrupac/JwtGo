package models

type User struct {
	ID       int    `json:"id" gorm:"type:varchar(255);primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);not null"`
	Password string `json:"password" gorm:"type:varchar(100);not null"`
}
