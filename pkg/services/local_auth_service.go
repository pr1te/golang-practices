package services

import (
	"github.com/pr1te/announcify-api/pkg/errors"
	"github.com/pr1te/announcify-api/pkg/logger"
	"github.com/pr1te/announcify-api/pkg/models"
	"github.com/pr1te/announcify-api/pkg/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LocalAuthService struct {
	logger *logger.Logger

	// repositories
	helperRepo      *repositories.HelperRepository
	localUserRepo   *repositories.LocalUserRepository
	userProfileRepo *repositories.UserProfileRepository
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err
}

func comparePassword(hashedPassword, comparedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(comparedPassword)); err != nil {
		return false
	}

	return true
}

func (service *LocalAuthService) CreateAccount(email string, password string) (*models.LocalUser, error) {
	var result *models.LocalUser

	if err := service.helperRepo.RunInTransaction(func(tx *gorm.DB) error {
		existingUser, err := service.localUserRepo.GetByEmail(email, repositories.GetOptions{
			Tx: tx,
		})

		if existingUser != nil {
			return errors.NewLocalUserDuplicated("the email has already existed", []interface{}{
				map[string]string{"email": email},
			})
		}

		if err != nil {
			return err
		}

		hashPassword, _ := hashPassword(password)
		user := service.localUserRepo.Create(models.LocalUser{Email: email, Password: hashPassword})

		result = &user

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (service *LocalAuthService) Login(email string, password string) (*models.LocalUser, error) {
	user, err := service.localUserRepo.GetByEmail(email)

	if err != nil {
		return nil, err
	}

	if verifiedPasswordResult := comparePassword(user.Password, password); verifiedPasswordResult {
		return user, nil
	}

	return nil, nil
}

func NewLocalAuth(
	logger *logger.Logger,

	// repositories
	helperRepo *repositories.HelperRepository,
	localUserRepo *repositories.LocalUserRepository,
	userProfileRepo *repositories.UserProfileRepository,
) *LocalAuthService {
	return &LocalAuthService{
		logger,

		helperRepo,
		localUserRepo,
		userProfileRepo,
	}
}
