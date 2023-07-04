package models

type UserProfile struct {
	Model
	LinkedLocalUser int `json:"linkedLocalUser" gorm:"uniqueindex"`
}
