package converter

import (
	spaceshipProto "imperial-fleet-inventory/api/langs/go/spaceship/grpc"
	"imperial-fleet-inventory/services/http/domain/model"
)

func ToProtoCreateSpaceshipRequest(req model.CreateSpaceshipRequest) *spaceshipProto.CreateSpaceshipRequest {
	armament := make([]*spaceshipProto.CreateSpaceshipRequest_Armament, 0, len(req.Armament))
	for _, a := range req.Armament {
		armament = append(armament, &spaceshipProto.CreateSpaceshipRequest_Armament{
			Title:    a.Title,
			Quantity: a.Qty,
		})
	}
	return &spaceshipProto.CreateSpaceshipRequest{
		Name:     req.Name,
		Class:    req.Class,
		Crew:     req.Crew,
		ImageUrl: req.Image,
		Value:    req.Value,
		Status:   req.Status,
		Armament: armament,
	}
}

func ToProtoUpdateSpaceshipRequest(spaceShipID int64, req model.UpdateSpaceshipRequest) *spaceshipProto.UpdateSpaceshipRequest {
	armament := make([]*spaceshipProto.UpdateSpaceshipRequest_Armament, 0, len(req.Armament))
	for _, a := range req.Armament {
		armament = append(armament, &spaceshipProto.UpdateSpaceshipRequest_Armament{
			Title:    a.Title,
			Quantity: a.Qty,
		})
	}
	return &spaceshipProto.UpdateSpaceshipRequest{
		Id:       spaceShipID,
		Name:     req.Name,
		Class:    req.Class,
		Crew:     req.Crew,
		ImageUrl: req.Image,
		Value:    req.Value,
		Status:   req.Status,
		Armament: armament,
	}
}

func ToProtoGetSpaceshipsRequest(q model.GetSpaceshipListQuery) *spaceshipProto.GetSpaceshipsRequest {
	return &spaceshipProto.GetSpaceshipsRequest{
		Name:   q.Name,
		Class:  q.Class,
		Status: q.Status,
	}
}

func FromProtoGetSpaceshipsResponse(res *spaceshipProto.GetSpaceshipsResponse) model.GetSpaceshipsListResponse {
	spaceships := make([]model.SpaceshipListShort, 0, len(res.Spaceships))
	for _, s := range res.Spaceships {
		spaceships = append(spaceships,
			model.SpaceshipListShort{
				ID:     s.Id,
				Name:   s.Name,
				Status: s.Status,
			},
		)
	}
	return model.GetSpaceshipsListResponse{
		Data: spaceships,
	}
}

func FromProtoSpaceship(spaceship *spaceshipProto.Spaceship) model.Spaceship {
	armament := make([]model.SpaceshipArmament, 0, len(spaceship.Armament))
	for _, a := range spaceship.Armament {
		armament = append(armament, model.SpaceshipArmament{
			Title: a.Title,
			Qty:   a.Quantity,
		})
	}
	return model.Spaceship{
		ID:       spaceship.Id,
		Name:     spaceship.Name,
		Class:    spaceship.Class,
		Crew:     spaceship.Crew,
		Image:    spaceship.ImageUrl,
		Value:    spaceship.Value,
		Status:   spaceship.Status,
		Armament: armament,
	}
}
