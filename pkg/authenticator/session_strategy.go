package authenticator

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/pr1te/announcify-api/pkg/errors"
)

type SessionStrategy struct {
	storage *session.Store
}

func (strategy *SessionStrategy) Authenticate(c *fiber.Ctx) (*AuthUser, error) {
	session, err := strategy.storage.Get(c)

	if err != nil {
		return nil, err
	}

	userID := session.Get("user-id")
	profileID := session.Get("profile-id")

	if userID == nil || profileID == nil {
		return nil, errors.NewUnauthorized("unauthorized")
	}

	authUser := AuthUser{UserID: userID.(uint), ProfileID: profileID.(uint)}
	strategy.Append(c, authUser)

	return &authUser, nil
}

func (strategy *SessionStrategy) Append(c *fiber.Ctx, authUser AuthUser) error {
	c.Locals("user-id", authUser.UserID)
	c.Locals("profile-id", authUser.ProfileID)

	return nil
}

func NewSessionStrategy(storage *session.Store) *SessionStrategy {
	return &SessionStrategy{storage}
}
