package repositories

import (
	"github.com/pr1te/announcify-api/pkg/database"
	"github.com/pr1te/announcify-api/pkg/models"
)

type UserProfileRepository struct {
	db *database.Database
}

func NewUserProfile(db *database.Database) *UserProfileRepository {
	db.Client.AutoMigrate(&models.UserProfile{})

	return &UserProfileRepository{db}
}
