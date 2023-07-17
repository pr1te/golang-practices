package repositories

import (
	"errors"

	"github.com/pr1te/left-it-api/pkg/database"
	"github.com/pr1te/left-it-api/pkg/models"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *database.Database
}

func (repo *ProfileRepository) Create(profile *models.Profile, options ...CreateOptions) (*models.Profile, error) {
	opts := CreateOptions{}

	if len(options) > 0 {
		if options[0].Tx != nil {
			opts.Tx = options[0].Tx
		}
	}

	var result *gorm.DB

	if opts.Tx != nil {
		result = opts.Tx.Create(&profile)
	} else {
		result = repo.db.Client.Create(&profile)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return profile, nil
}

func (repo *ProfileRepository) GetById(id uint, options ...GetOptions) (*models.Profile, error) {
	profile := &models.Profile{}

	opts := GetOptions{
		Fields: []string{"*"},
	}

	if len(options) > 0 {
		if options[0].Tx != nil {
			opts.Tx = options[0].Tx
		}

		if options[0].Fields != nil {
			opts.Fields = options[0].Fields
		}
	}

	var result *gorm.DB

	if opts.Tx != nil {
		result = opts.Tx.First(profile, "id = ? AND deleted = ?", id, false).Select(opts.Fields)
	} else {
		result = repo.db.Client.First(profile, "id = ? AND deleted = ?", id, false).Select(opts.Fields)
	}

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return profile, nil
}

func NewProfile(db *database.Database) *ProfileRepository {
	return &ProfileRepository{db}
}
