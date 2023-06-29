package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/logger"
	"github.com/pr1te/announcify-api/pkg/models"
	"github.com/pr1te/announcify-api/pkg/services"
)

type WorkspaceController struct {
	logger           *logger.Logger
	workspaceService *services.WorkspaceService
}

func (controller *WorkspaceController) CreateWorkspace(c *fiber.Ctx) error {
	workspace := &models.Workspace{}

	c.BodyParser(workspace)
	result := controller.workspaceService.Create(workspace)

	error := c.JSON(result)

	return error
}

func (controller *WorkspaceController) GetWorkspace(c *fiber.Ctx) error {
	result := controller.workspaceService.Find()
	error := c.JSON(result)

	return error
}

func NewWorkspace(logger *logger.Logger, workspaceService *services.WorkspaceService) *WorkspaceController {
	return &WorkspaceController{
		logger,
		workspaceService,
	}
}
