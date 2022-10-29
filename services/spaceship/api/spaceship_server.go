package api

import (
	"context"

	"github.com/pkg/errors"

	spaceshipProto "imperial-fleet-inventory/api/langs/go/spaceship/grpc"
	"imperial-fleet-inventory/services/spaceship/api/converter"
	"imperial-fleet-inventory/services/spaceship/domain/model"
)

type SpaceshipServer struct {
	spaceshipSrv SpaceshipService
}

type SpaceshipService interface {
	FindAllSpaceships(ctx context.Context, filter model.GetSpaceshipListQuery) ([]model.Spaceship, error)
	CreateSpaceship(ctx context.Context, spaceship model.CreateSpaceshipRequest) (model.Spaceship, error)
	FindSpaceship(ctx context.Context, ID int64) (model.Spaceship, error)
	UpdateSpaceship(ctx context.Context, ID int64, spaceship model.UpdateSpaceshipRequest) (model.Spaceship, error)
	DeleteSpaceship(ctx context.Context, ID int64) error
}

func NewSpaceshipServer(spaceshipSrv SpaceshipService) *SpaceshipServer {
	return &SpaceshipServer{
		spaceshipSrv: spaceshipSrv,
	}
}

func (s *SpaceshipServer) CreateSpaceship(ctx context.Context, request *spaceshipProto.CreateSpaceshipRequest) (*spaceshipProto.CreateSpaceshipResponse, error) {
	req := converter.FromProtoCreateSpaceshipRequest(request)

	spaceship, err := s.spaceshipSrv.CreateSpaceship(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create spaceship")
	}

	return converter.ToProtoCreateSpaceshipResponse(spaceship), nil
}

func (s *SpaceshipServer) GetSpaceships(ctx context.Context, request *spaceshipProto.GetSpaceshipsRequest) (*spaceshipProto.GetSpaceshipsResponse, error) {
	req := converter.FromProtoGetSpaceshipsRequest(request)

	spaceships, err := s.spaceshipSrv.FindAllSpaceships(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get spaceships")
	}

	return converter.ToProtoGetSpaceshipsResponse(spaceships), nil
}

func (s *SpaceshipServer) GetSpaceship(ctx context.Context, request *spaceshipProto.GetSpaceshipRequest) (*spaceshipProto.GetSpaceshipResponse, error) {
	id := request.GetId()

	spaceship, err := s.spaceshipSrv.FindSpaceship(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get spaceship")
	}

	return converter.ToProtoGetSpaceshipResponse(spaceship), nil
}

func (s *SpaceshipServer) UpdateSpaceship(ctx context.Context, request *spaceshipProto.UpdateSpaceshipRequest) (*spaceshipProto.UpdateSpaceshipResponse, error) {
	id := request.GetId()
	req := converter.FromProtoUpdateSpaceshipRequest(request)

	spaceship, err := s.spaceshipSrv.UpdateSpaceship(ctx, id, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update spaceship")
	}

	return converter.ToProtoUpdateSpaceshipResponse(spaceship), nil

}

func (s *SpaceshipServer) DeleteSpaceship(ctx context.Context, request *spaceshipProto.DeleteSpaceshipRequest) (*spaceshipProto.DeleteSpaceshipResponse, error) {
	// TODO implement me
	panic("implement me")
}
