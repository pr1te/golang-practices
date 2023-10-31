package services

import (
	"github.com/pr1te/left-it-api/pkg/errors"
	"github.com/pr1te/left-it-api/pkg/logger"
	"github.com/pr1te/left-it-api/pkg/models"
	"github.com/pr1te/left-it-api/pkg/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LocalUserService struct {
	logger *logger.Logger

	// repositories
	helperRepo          *repositories.HelperRepository
	localUserRepo       *repositories.LocalUserRepository
	ProfileRepo         *repositories.ProfileRepository
	userProfileLinkRepo *repositories.UserProfileLinkRepository
}

type LocalUserServiceCreateAccountInfo struct {
	Email       string
	Password    string
	DisplayName string
}

type LocalUserServiceCreateAccountResult struct {
	User    *models.LocalUser `json:"user"`
	Profile *models.Profile   `json:"profile"`
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

func (service *LocalUserService) CreateAccount(info LocalUserServiceCreateAccountInfo) (*LocalUserServiceCreateAccountResult, error) {
	var result *LocalUserServiceCreateAccountResult

	if err := service.helperRepo.RunInTransaction(func(tx *gorm.DB) error {
		existingUser, err := service.localUserRepo.GetByEmail(info.Email, repositories.GetOptions{
			Tx: tx,
		})

		if existingUser != nil {
			return errors.NewLocalUserDuplicated("the email has already existed", []interface{}{
				map[string]string{"email": info.Email},
			})
		}

		if err != nil {
			return err
		}

		var (
			profile *models.Profile
			user    *models.LocalUser
			linked  *models.UserProfileLink
		)

		hashPassword, _ := hashPassword(info.Password)
		profile, _ = service.ProfileRepo.Create(&models.Profile{DisplayName: info.DisplayName}, repositories.CreateOptions{Tx: tx})
		user, _ = service.localUserRepo.Create(&models.LocalUser{Email: info.Email, Password: hashPassword}, repositories.CreateOptions{Tx: tx})
		linked, _ = service.userProfileLinkRepo.Create(&models.UserProfileLink{UserID: user.ID, ProfileID: profile.ID, Type: models.LOCAL_USER_TYPE})

		service.logger.Debugf("new account created: %+v %+v %+v", user, profile, linked)

		result = &LocalUserServiceCreateAccountResult{
			User:    user,
			Profile: profile,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (service *LocalUserService) Login(email string, password string) (*models.LocalUser, *models.Profile, error) {
	user, err := service.localUserRepo.GetByEmail(email)

	if err != nil {
		return nil, nil, err
	}

	if user == nil {
		return nil, nil, nil
	}

	if verifiedPasswordResult := comparePassword(user.Password, password); verifiedPasswordResult {
		linked, getLinkErr := service.userProfileLinkRepo.GetByUserIDAndType(user.ID, models.LOCAL_USER_TYPE)

		if getLinkErr != nil {
			return nil, nil, getLinkErr
		}

		profile, getProfileErr := service.ProfileRepo.GetById(linked.ProfileID)

		if getProfileErr != nil {
			return nil, nil, getProfileErr
		}

		return user, profile, nil
	}

	return nil, nil, nil
}

func NewLocalUser(
	logger *logger.Logger,

	// repositories
	helperRepo *repositories.HelperRepository,
	localUserRepo *repositories.LocalUserRepository,
	ProfileRepo *repositories.ProfileRepository,
	userProfileLinkRepo *repositories.UserProfileLinkRepository,
) *LocalUserService {
	return &LocalUserService{
		logger,

		helperRepo,
		localUserRepo,
		ProfileRepo,
		userProfileLinkRepo,
	}
}
