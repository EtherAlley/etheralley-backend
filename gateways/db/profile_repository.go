package db

import (
	"github.com/eflem00/go-example-app/common"
	"github.com/eflem00/go-example-app/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// TODO: fill in the persistant storage piece

type ProfileRepository struct {
	db     *gorm.DB
	logger *common.Logger
}

func NewProfileRepository(db *gorm.DB, logger *common.Logger) *ProfileRepository {
	return &ProfileRepository{
		db,
		logger,
	}
}

func (repo *ProfileRepository) GetProfileByAddress(address string) (entities.Profile, error) {
	var profile entities.Profile

	if err := repo.db.Preload("Elements").First(&profile, "address = ?", address).Error; err != nil {
		repo.logger.Err(err, "error getting profile by address")
		return entities.Profile{}, err
	}

	return profile, nil
}

func (repo *ProfileRepository) SaveProfile(address string, profile entities.Profile) error {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("profile_address = ?", profile.Address).Delete(&entities.Element{}).Error; err != nil {
			return err
		}

		if err := tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(profile).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
