package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	commonDomain "imperial-fleet-inventory/common/sevice/domain"
	"imperial-fleet-inventory/services/spaceship/domain/model"
)

type SpaceshipRepo struct {
	db *gorm.DB
}

func NewSpaceshipRepo(db *gorm.DB) *SpaceshipRepo {
	return &SpaceshipRepo{db: db}
}

func (r *SpaceshipRepo) CreateSpaceship(_ context.Context, spaceship model.Spaceship) (model.Spaceship, error) {
	err := r.db.Create(&spaceship).Error
	if err != nil {
		return model.Spaceship{}, commonDomain.WrapWithInternalError(err, "create spaceship failed")
	}

	return spaceship, nil
}

func (r *SpaceshipRepo) UpdateSpaceship(_ context.Context, spaceship model.Spaceship) (model.Spaceship, error) {
	assocErr := r.db.Where("spaceship_id = ?", spaceship.ID).Delete(&model.SpaceshipArmament{}).Error
	if assocErr != nil {
		return model.Spaceship{}, commonDomain.WrapWithInternalError(assocErr, "update spaceship failed")
	}

	err := r.db.Save(&spaceship).Error
	if err != nil {
		return model.Spaceship{}, commonDomain.WrapWithInternalError(err, "update spaceship failed")
	}
	return spaceship, nil
}

func (r *SpaceshipRepo) FindSpaceship(_ context.Context, ID int64) (model.Spaceship, error) {
	spaceship := model.Spaceship{}
	err := r.db.Where("id = ?", ID).Preload("Armament").First(&spaceship).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Spaceship{}, commonDomain.WrapWithNotFoundError(err, "spaceship not found")
		}
		return model.Spaceship{}, commonDomain.WrapWithInternalError(err, "find spaceship by id failed")
	}
	return spaceship, nil
}

func (r *SpaceshipRepo) FindAllSpaceships(_ context.Context, filter model.GetSpaceshipListQuery) ([]model.Spaceship, error) {
	f := model.Spaceship{}
	if filter.Name != "" {
		f.Name = filter.Name
	}
	if filter.Status != "" {
		f.Status = filter.Status
	}
	if filter.Class != "" {
		f.Class = filter.Class
	}

	var spaceships []model.Spaceship
	err := r.db.Where(&f).Find(&spaceships).Error
	if err != nil {
		return nil, commonDomain.WrapWithInternalError(err, "find all spaceships failed")
	}

	return spaceships, nil
}

func (r *SpaceshipRepo) DeleteSpaceship(_ context.Context, ID int64) error {
	err := r.db.Select("Armament").Delete(&model.Spaceship{ID: ID}).Error
	if err != nil {
		return commonDomain.WrapWithInternalError(err, "remove spaceship failed")
	}
	return nil
}
