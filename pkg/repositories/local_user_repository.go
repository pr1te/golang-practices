package repositories

import (
	"github.com/pr1te/announcify-api/pkg/database"
	"github.com/pr1te/announcify-api/pkg/models"
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
		return nil, result.Error
	}

	return user, nil
}

func (repo *LocalUserRepository) Create() string {
	return "created"
}

func NewLocalUser(db *database.Database) *LocalUserRepository {
	db.Client.AutoMigrate(&models.LocalUser{})

	return &LocalUserRepository{db}
}
