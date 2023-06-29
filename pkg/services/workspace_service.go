package services

import (
	"github.com/pr1te/announcify-api/pkg/models"
	"github.com/pr1te/announcify-api/pkg/repositories"
)

type WorkspaceService struct {
	workspaceRepository *repositories.WorkspaceRepository
}

func (service *WorkspaceService) Create(workspace *models.Workspace) *models.Workspace {
	return service.workspaceRepository.Create(workspace)
}

func (service *WorkspaceService) Find() *[]models.Workspace {
	return service.workspaceRepository.Find()
}

func NewWorkspace(workspaceRepository *repositories.WorkspaceRepository) *WorkspaceService {
	return &WorkspaceService{
		workspaceRepository,
	}
}
