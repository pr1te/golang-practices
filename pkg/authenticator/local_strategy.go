package authenticator

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/pr1te/announcify-api/pkg/errors"
	"github.com/pr1te/announcify-api/pkg/libs/validator"
	"github.com/pr1te/announcify-api/pkg/services"
)

type credential struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LocalStrategy struct {
	storage          *session.Store
	validator        *validator.Validator
	localUserService *services.LocalUserService
}

func (strategy *LocalStrategy) Authenticate(c *fiber.Ctx) (*AuthUser, error) {
	credential := &credential{}

	if err := c.BodyParser(credential); err != nil {
		return nil, errors.NewUnauthorized("unauthorized")
	}

	if err := strategy.validator.ValidateStruct(credential); err != nil {
		details := make([]interface{}, len(err))
		for index, value := range err {
			details[index] = value
		}

		return nil, errors.NewValidationFailed("validation error", details)
	}

	user, profile, err := strategy.localUserService.Login(credential.Email, credential.Password)

	if err != nil || user == nil || profile == nil {
		return nil, errors.NewUnauthorized("unauthorized")
	}

	authUser := AuthUser{UserID: user.ID, ProfileID: profile.ID}

	if err := strategy.Append(c, authUser); err != nil {
		return nil, err
	}

	return &authUser, nil
}

func (strategy *LocalStrategy) Append(c *fiber.Ctx, authUser AuthUser) error {
	session, err := strategy.storage.Get(c)

	if err != nil {
		return err
	}

	session.Set("user-id", authUser.UserID)
	session.Set("profile-id", authUser.ProfileID)
	session.SetExpiry(1 * time.Hour)

	if saveErr := session.Save(); saveErr != nil {
		return saveErr
	}

	return nil
}

func NewLocalStrategy(session *session.Store, validator *validator.Validator, localUserService *services.LocalUserService) *LocalStrategy {
	return &LocalStrategy{
		session,
		validator,
		localUserService,
	}
}
