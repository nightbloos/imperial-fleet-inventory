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
	CreateSpaceship(ctx context.Context, spaceship model.Spaceship) (model.Spaceship, error)
	FindSpaceship(ctx context.Context, ID int64) (model.Spaceship, error)
	UpdateSpaceship(ctx context.Context, spaceship model.Spaceship) (model.Spaceship, error)
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

func (s *Service) CreateSpaceship(ctx context.Context, createReq model.CreateSpaceshipRequest) (model.Spaceship, error) {
	armaments := make([]model.SpaceshipArmament, 0, len(createReq.Armament))
	for _, armament := range createReq.Armament {
		armaments = append(armaments, model.SpaceshipArmament{
			Title:    armament.Title,
			Quantity: armament.Quantity,
		})
	}
	spaceship := model.Spaceship{
		Name:     createReq.Name,
		Class:    createReq.Class,
		Crew:     createReq.Crew,
		Image:    createReq.Image,
		Value:    createReq.Value,
		Status:   createReq.Status,
		Armament: armaments,
	}

	return s.repo.CreateSpaceship(ctx, spaceship)
}

func (s *Service) FindSpaceship(ctx context.Context, ID int64) (model.Spaceship, error) {
	return s.repo.FindSpaceship(ctx, ID)
}

func (s *Service) UpdateSpaceship(ctx context.Context, ID int64, updateReq model.UpdateSpaceshipRequest) (model.Spaceship, error) {
	armaments := make([]model.SpaceshipArmament, 0, len(updateReq.Armament))
	for _, armament := range updateReq.Armament {
		armaments = append(armaments, model.SpaceshipArmament{
			Title:    armament.Title,
			Quantity: armament.Quantity,
		})
	}
	spaceship := model.Spaceship{
		ID:       ID,
		Name:     updateReq.Name,
		Class:    updateReq.Class,
		Crew:     updateReq.Crew,
		Image:    updateReq.Image,
		Value:    updateReq.Value,
		Status:   updateReq.Status,
		Armament: armaments,
	}

	return s.repo.UpdateSpaceship(ctx, spaceship)
}

func (s *Service) DeleteSpaceship(ctx context.Context, ID int64) error {
	return s.repo.DeleteSpaceship(ctx, ID)
}
