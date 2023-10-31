package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type UserType string

const (
	LOCAL_USER_TYPE  UserType = "local"
	GOOGLE_USER_TYPE UserType = "google"
)

type UserProfileLink struct {
	UserID    uint      `json:"userId" gorm:"primaryKey;autoIncrement:false;not null"`
	ProfileID uint      `json:"profileId" gorm:"index:,unique,composite:profile_id_type;not null"`
	Type      UserType  `json:"type" gorm:"primaryKey;autoIncrement:false;index:,unique,composite:profile_id_type;not null"`
	CreatedAt time.Time `json:"createdAt"`
}

// Value validate enum when set to database
func (t UserType) Value() (driver.Value, error) {
	switch t {
	case LOCAL_USER_TYPE, GOOGLE_USER_TYPE:
		return string(t), nil
	}

	return nil, fmt.Errorf("invalid type. got '%s'", t)
}

// Scan validate enum on read from data base
func (t *UserType) Scan(value interface{}) error {
	if value == nil {
		*t = ""
		return nil
	}

	str, ok := value.(string)

	if !ok {
		return fmt.Errorf("invalid type. got '%s'", value)
	}

	// convert type from string to ProductType
	at := UserType(str)

	switch at {
	case LOCAL_USER_TYPE, GOOGLE_USER_TYPE:
		*t = at

		return nil
	}

	return fmt.Errorf("invalid type. got '%s'", at)
}
