package entity

import "time"

type User struct {
	ID          int
	PhoneNumber int `gorm:"unique"`
	Password    string
	Name        string
	Surname     string
	DateOfBirth string
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	URL         string    `gorm:"default:''"`
	Messages    []Message
}
