package repositories

import (
	"errors"

	"github.com/pr1te/left-it-api/pkg/database"
	"github.com/pr1te/left-it-api/pkg/models"
	"gorm.io/gorm"
)

type UserProfileLinkRepository struct {
	db *database.Database
}

func (repo *UserProfileLinkRepository) Create(account *models.UserProfileLink, options ...CreateOptions) (*models.UserProfileLink, error) {
	opts := CreateOptions{}

	if len(options) > 0 {
		if options[0].Tx != nil {
			opts.Tx = options[0].Tx
		}
	}

	var result *gorm.DB

	if opts.Tx != nil {
		result = opts.Tx.Create(&account)
	} else {
		result = repo.db.Client.Create(&account)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return account, nil
}

func (repo *UserProfileLinkRepository) GetByUserIDAndType(UserID uint, accountType models.UserType, options ...GetOptions) (*models.UserProfileLink, error) {
	linked := &models.UserProfileLink{}

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
		result = opts.Tx.First(linked, "user_id = ? AND type = ?", UserID, accountType).Select(opts.Fields)
	} else {
		result = repo.db.Client.First(linked, "user_id = ? AND type = ?", UserID, accountType).Select(opts.Fields)
	}

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return linked, nil
}

func NewUserProfileLink(db *database.Database) *UserProfileLinkRepository {
	return &UserProfileLinkRepository{db}
}
