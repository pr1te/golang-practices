package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/logger"
	"github.com/pr1te/announcify-api/pkg/services"
)

type MeController struct {
	logger         *logger.Logger
	profileService *services.ProfileService
}

func (controller *MeController) GetMyProfile(c *fiber.Ctx) error {
	id := c.Locals("profile-id").(uint)

	profile, err := controller.profileService.GetById(id)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(profile)
}

func NewMe(logger *logger.Logger, profileService *services.ProfileService) *MeController {
	return &MeController{
		logger,
		profileService,
	}
}
