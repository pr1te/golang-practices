package models

type UserProfile struct {
	Model
	LinkedLocalUser int  `json:"linkedLocalUser" gorm:"index:,unique,composite:linked_local_user_deleted"`
	Deleted         bool `json:"deleted" gorm:"index:,unique,composite:linked_local_user_deleted;default:false"`
}
