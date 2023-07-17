package models

type Profile struct {
	Model
	FirstName string `json:"firstName" gorm:"not null,size:100"`
	LastName  string `json:"lastName" gorm:"not null,size:100"`
}
