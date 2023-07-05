package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/exceptions"
	"github.com/pr1te/announcify-api/pkg/libs/validator"
	"github.com/pr1te/announcify-api/pkg/services"
)

type LocalAuthController struct {
	localAuthService *services.LocalAuthService
}

type CreateAccountBodyDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

func (controller *LocalAuthController) CreateAccount(c *fiber.Ctx, validator *validator.Validator) error {
	body := &CreateAccountBodyDTO{}
	c.BodyParser(body)

	if err := validator.ValidateStruct(body); err != nil {
		details := make([]interface{}, len(err))
		for index, value := range err {
			details[index] = value
		}

		return exceptions.NewValidationErrorException("validation error", details)
	}

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
