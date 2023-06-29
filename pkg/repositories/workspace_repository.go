package repositories

import (
	"github.com/pr1te/announcify-api/pkg/database"
	"github.com/pr1te/announcify-api/pkg/models"
)

type WorkspaceRepository struct {
	db *database.Database
}

func (repo *WorkspaceRepository) Create(workspace *models.Workspace) *models.Workspace {
	repo.db.Client.Create(&workspace)

	return workspace
}

func (repo *WorkspaceRepository) Find() *[]models.Workspace {
	var workspaces *[]models.Workspace

	repo.db.Client.Find(&workspaces)

	return workspaces
}

func NewWorkspace(db *database.Database) *WorkspaceRepository {
	db.Client.AutoMigrate(&models.Workspace{})

	return &WorkspaceRepository{db}
}
