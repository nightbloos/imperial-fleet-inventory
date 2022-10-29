package spaceship

import (
	"context"

	"go.uber.org/zap"

	"imperial-fleet-inventory/services/spaceship/domain/model"
)

type Service struct {
	repo   Repository
	logger *zap.Logger
}

type Repository interface {
	FindAllSpaceships(ctx context.Context, filter model.GetSpaceshipListQuery) ([]model.Spaceship, error)
	CreateSpaceship(ctx context.Context, spaceship model.CreateSpaceshipRequest) (model.Spaceship, error)
	FindSpaceship(ctx context.Context, ID int64) (model.Spaceship, error)
	UpdateSpaceship(ctx context.Context, ID int64, spaceship model.UpdateSpaceshipRequest) (model.Spaceship, error)
	DeleteSpaceship(ctx context.Context, ID int64) error
}

func NewService(repo Repository, logger *zap.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger.With(zap.String("internal-service", "spaceship")),
	}
}

func (s *Service) FindAllSpaceships(ctx context.Context, filter model.GetSpaceshipListQuery) ([]model.Spaceship, error) {
	return s.repo.FindAllSpaceships(ctx, filter)
}

func (s *Service) CreateSpaceship(ctx context.Context, spaceship model.CreateSpaceshipRequest) (model.Spaceship, error) {
	return s.repo.CreateSpaceship(ctx, spaceship)
}

func (s *Service) FindSpaceship(ctx context.Context, ID int64) (model.Spaceship, error) {
	return s.repo.FindSpaceship(ctx, ID)
}

func (s *Service) UpdateSpaceship(ctx context.Context, ID int64, spaceship model.UpdateSpaceshipRequest) (model.Spaceship, error) {
	return s.repo.UpdateSpaceship(ctx, ID, spaceship)
}

func (s *Service) DeleteSpaceship(ctx context.Context, ID int64) error {
	return s.repo.DeleteSpaceship(ctx, ID)
}
