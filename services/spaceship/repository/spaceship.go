package repository

import (
	"context"

	"gorm.io/gorm"

	"imperial-fleet-inventory/services/spaceship/domain/model"
)

type SpaceshipRepo struct {
	db *gorm.DB
}

func NewSpaceshipRepo(db *gorm.DB) *SpaceshipRepo {
	return &SpaceshipRepo{db: db}
}

func (s SpaceshipRepo) CreateSpaceship(ctx context.Context, spaceship model.CreateSpaceshipRequest) (model.Spaceship, error) {
	// TODO implement me
	panic("implement me")
}

func (s SpaceshipRepo) UpdateSpaceship(ctx context.Context, ID int64, spaceship model.UpdateSpaceshipRequest) (model.Spaceship, error) {
	// TODO implement me
	panic("implement me")
}

func (s SpaceshipRepo) FindSpaceship(ctx context.Context, ID int64) (model.Spaceship, error) {
	// TODO implement me
	panic("implement me")
}

func (s SpaceshipRepo) FindAllSpaceships(ctx context.Context, filter model.GetSpaceshipListQuery) ([]model.Spaceship, error) {
	// TODO implement me
	panic("implement me")
}

func (s SpaceshipRepo) DeleteSpaceship(ctx context.Context, ID int64) error {
	// TODO implement me
	panic("implement me")
}
