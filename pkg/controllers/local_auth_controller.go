package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/services"
)

type LocalAuthController struct {
	localAuthService *services.LocalAuthService
}

type CreateAccountBodyDTO struct {
	Email string `json:"email"`
}

func (controller *LocalAuthController) CreateAccount(c *fiber.Ctx) error {
	body := new(CreateAccountBodyDTO)
	c.BodyParser(&body)

	result, err := controller.localAuthService.CreateAccount(body.Email)

	if err != nil {
		return err
	}

	c.JSON(result)

	return nil
}

func NewLocalAuth(localAuthService *services.LocalAuthService) *LocalAuthController {
	return &LocalAuthController{localAuthService}
}
