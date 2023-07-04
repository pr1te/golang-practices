package models

type LocalUser struct {
	Model
	Email    string `json:"email" gorm:"index:,unique,composite:email_deleted;not null"`
	Password string `json:"password" gorm:"not null"`
	Deleted  bool   `json:"deleted" gorm:"index:,unique,composite:email_deleted;default:false"`
}
