package repositories

import (
	"github.com/pr1te/announcify-api/pkg/database"
	"gorm.io/gorm"
)

type HelperRepository struct {
	db *database.Database
}

type TxFunc func(tx *gorm.DB) error

type GetOptions struct {
	Tx     *gorm.DB
	Fields interface{}
}

type CreateOptions struct {
	Tx *gorm.DB
}

func (repo *HelperRepository) RunInTransaction(fn TxFunc) error {
	tx := repo.db.Client.Begin()

	if err := fn(tx); err != nil {
		tx.Rollback()

		return err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	return tx.Commit().Error
}

func NewHelper(db *database.Database) *HelperRepository {
	return &HelperRepository{db}
}
