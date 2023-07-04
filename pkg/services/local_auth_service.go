package services

import (
	"github.com/pr1te/announcify-api/pkg/exceptions"
	"github.com/pr1te/announcify-api/pkg/models"
	"github.com/pr1te/announcify-api/pkg/repositories"
	"gorm.io/gorm"
)

type LocalAuthService struct {
	helperRepo    *repositories.HelperRepository
	localUserRepo *repositories.LocalUserRepository
}

func (service *LocalAuthService) CreateAccount(email string) (*models.LocalUser, error) {
	var result *models.LocalUser

	if err := service.helperRepo.RunInTransaction(func(tx *gorm.DB) error {
		user, err := service.localUserRepo.GetByEmail(email, repositories.GetOptions{
			Tx: tx,
		})

		if user != nil {
			return exceptions.NewDuplicateLocalUserException("the email has already existed", []interface{}{
				map[string]string{"email": email},
			})
		}

		if err != nil {
			return err
		}

		result = user

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func NewLocalAuth(localUserRepo *repositories.LocalUserRepository, helperRepo *repositories.HelperRepository) *LocalAuthService {
	return &LocalAuthService{
		helperRepo,
		localUserRepo,
	}
}
