package repositories

import (
	"errors"

	"github.com/pr1te/left-it-api/pkg/database"
	"github.com/pr1te/left-it-api/pkg/models"
	"gorm.io/gorm"
)

type LocalUserRepository struct {
	db *database.Database
}

func (repo *LocalUserRepository) GetByEmail(email string, options ...GetOptions) (*models.LocalUser, error) {
	user := &models.LocalUser{}

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
		result = opts.Tx.First(user, "email = ? AND deleted = ?", email, false).Select(opts.Fields)
	} else {
		result = repo.db.Client.First(user, "email = ? AND deleted = ?", email, false).Select(opts.Fields)
	}

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return user, nil
}

func (repo *LocalUserRepository) Create(user *models.LocalUser, options ...CreateOptions) (*models.LocalUser, error) {
	opts := CreateOptions{}

	if len(options) > 0 {
		if options[0].Tx != nil {
			opts.Tx = options[0].Tx
		}
	}

	var result *gorm.DB

	if opts.Tx != nil {
		result = opts.Tx.Create(&user)
	} else {
		result = repo.db.Client.Create(&user)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func NewLocalUser(db *database.Database) *LocalUserRepository {
	return &LocalUserRepository{db}
}
