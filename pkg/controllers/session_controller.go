package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type SessionController struct {
	storage *session.Store
}

func (controller *SessionController) Destroy(c *fiber.Ctx) error {
	session, err := controller.storage.Get(c)

	if err != nil {
		return err
	}

	if err := session.Destroy(); err != nil {
		return err
	}

	return nil
}

func NewSession(storage *session.Store) *SessionController {
	return &SessionController{storage}
}
