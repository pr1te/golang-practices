package models

type Workspace struct {
	Model
	Title string `json:"title" gorm:"not null"`
}
