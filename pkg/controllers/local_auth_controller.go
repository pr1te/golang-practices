package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/errors"
	"github.com/pr1te/announcify-api/pkg/libs/validator"
	"github.com/pr1te/announcify-api/pkg/services"
)

type LocalAuthController struct {
	localAuthService *services.LocalAuthService
}

type createAccountBodyDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type loginBodyDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (controller *LocalAuthController) CreateAccount(c *fiber.Ctx, validator *validator.Validator) error {
	body := &createAccountBodyDTO{}
	c.BodyParser(body)

	if err := validator.ValidateStruct(body); err != nil {
		details := make([]interface{}, len(err))
		for index, value := range err {
			details[index] = value
		}

		return errors.NewValidationFailed("validation error", details)
	}

	result, err := controller.localAuthService.CreateAccount(body.Email, body.Password)

	if err != nil {
		return err
	}

	c.Status(200).JSON(result)

	return nil
}

func (controller *LocalAuthController) Login(c *fiber.Ctx, validator *validator.Validator) error {
	body := &loginBodyDTO{}
	c.BodyParser(body)

	if err := validator.ValidateStruct(body); err != nil {
		details := make([]interface{}, len(err))
		for index, value := range err {
			details[index] = value
		}

		return errors.NewValidationFailed("validation error", details)
	}

	result, err := controller.localAuthService.Login(body.Email, body.Password)

	if err != nil {
		return err
	}

	if result != nil {
		c.Status(200).JSON(result)

		return nil
	}

	return errors.NewUnauthorized("unauthorized")
}

func NewLocalAuth(localAuthService *services.LocalAuthService) *LocalAuthController {
	return &LocalAuthController{localAuthService}
}
