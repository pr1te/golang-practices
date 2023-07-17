package models

type Profile struct {
	Model
	DisplayName string `json:"displayName" gorm:"not null,size:100"`
}
