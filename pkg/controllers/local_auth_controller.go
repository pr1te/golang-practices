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

func (controller *LocalAuthController) CreateAccount(c *fiber.Ctx) {
	body := new(CreateAccountBodyDTO)
	c.BodyParser(&body)

	result, err := controller.localAuthService.CreateAccount(body.Email)

	if err != nil {
		c.JSON(err)
	} else {
		c.JSON(result)
	}
}

func NewLocalAuth(localAuthService *services.LocalAuthService) *LocalAuthController {
	return &LocalAuthController{localAuthService}
}
