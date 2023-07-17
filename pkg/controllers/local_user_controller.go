package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/errors"
	"github.com/pr1te/announcify-api/pkg/libs/validator"
	"github.com/pr1te/announcify-api/pkg/services"
)

type LocalUserController struct {
	localUserService *services.LocalUserService
}

type createAccountBodyDTO struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,password"`
	FirstName string `json:"firstName" validate:"required,alpha,max=100"`
	LastName  string `json:"lastName" validate:"required,alpha,max=100"`
}

func (controller *LocalUserController) CreateAccount(c *fiber.Ctx, validator *validator.Validator) error {
	body := &createAccountBodyDTO{}
	c.BodyParser(body)

	if err := validator.ValidateStruct(body); err != nil {
		details := make([]interface{}, len(err))
		for index, value := range err {
			details[index] = value
		}

		return errors.NewValidationFailed("validation error", details)
	}

	result, err := controller.localUserService.CreateAccount(services.LocalUserServiceCreateAccountInfo{
		Email:     body.Email,
		Password:  body.Password,
		FirstName: body.FirstName,
		LastName:  body.LastName,
	})

	if err != nil {
		return err
	}

	c.Status(200).JSON(result)

	return nil
}

func NewLocalUser(localUserService *services.LocalUserService) *LocalUserController {
	return &LocalUserController{localUserService}
}
