package services

import (
	"github.com/pr1te/left-it-api/pkg/models"
	"github.com/pr1te/left-it-api/pkg/repositories"
)

type ProfileService struct {
	profileRepo *repositories.ProfileRepository
}

func (service *ProfileService) GetById(id uint) (*models.Profile, error) {
	return service.profileRepo.GetById(id)
}

func NewProfile(profileRepo *repositories.ProfileRepository) *ProfileService {
	return &ProfileService{
		profileRepo,
	}
}
