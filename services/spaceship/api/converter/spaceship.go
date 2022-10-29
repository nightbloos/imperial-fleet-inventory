package converter

import (
	spaceshipProto "imperial-fleet-inventory/api/langs/go/spaceship/grpc"
	"imperial-fleet-inventory/services/spaceship/domain/model"
)

func FromProtoCreateSpaceshipRequest(req *spaceshipProto.CreateSpaceshipRequest) model.CreateSpaceshipRequest {
	return model.CreateSpaceshipRequest{
		Name:     req.GetName(),
		Class:    req.GetClass(),
		Crew:     req.GetCrew(),
		Image:    req.GetImageUrl(),
		Value:    req.GetValue(),
		Status:   req.GetStatus(),
		Armament: FromProtoCreateSpaceshipArmamentRequest(req.GetArmament()),
	}
}

func FromProtoCreateSpaceshipArmamentRequest(req []*spaceshipProto.CreateSpaceshipRequest_Armament) []model.CreateSpaceshipArmamentRequest {
	var armament []model.CreateSpaceshipArmamentRequest
	for _, a := range req {
		armament = append(armament, model.CreateSpaceshipArmamentRequest{
			Title:    a.GetTitle(),
			Quantity: a.GetQuantity(),
		})
	}
	return armament
}

func ToProtoCreateSpaceshipResponse(s model.Spaceship) *spaceshipProto.CreateSpaceshipResponse {
	return &spaceshipProto.CreateSpaceshipResponse{
		Spaceship: ToProtoSpaceship(s),
	}
}

func FromProtoGetSpaceshipsRequest(req *spaceshipProto.GetSpaceshipsRequest) model.GetSpaceshipListQuery {
	return model.GetSpaceshipListQuery{
		Name:   req.GetName(),
		Class:  req.GetClass(),
		Status: req.GetStatus(),
	}
}

func ToProtoGetSpaceshipsResponse(spaceships []model.Spaceship) *spaceshipProto.GetSpaceshipsResponse {
	var res []*spaceshipProto.Spaceship
	for _, s := range spaceships {
		res = append(res, ToProtoSpaceship(s))
	}

	return &spaceshipProto.GetSpaceshipsResponse{
		Spaceships: res,
	}
}

func ToProtoGetSpaceshipResponse(spaceship model.Spaceship) *spaceshipProto.GetSpaceshipResponse {
	return &spaceshipProto.GetSpaceshipResponse{
		Spaceship: ToProtoSpaceship(spaceship),
	}
}

func FromProtoUpdateSpaceshipRequest(req *spaceshipProto.UpdateSpaceshipRequest) model.UpdateSpaceshipRequest {

	return model.UpdateSpaceshipRequest{
		Name:     req.GetName(),
		Class:    req.GetClass(),
		Crew:     req.GetCrew(),
		Image:    req.GetImageUrl(),
		Value:    req.GetValue(),
		Status:   req.GetStatus(),
		Armament: FromProtoUpdateSpaceshipArmamentRequest(req.GetArmament()),
	}
}

func FromProtoUpdateSpaceshipArmamentRequest(req []*spaceshipProto.UpdateSpaceshipRequest_Armament) []model.UpdateSpaceshipArmamentRequest {
	var armament []model.UpdateSpaceshipArmamentRequest
	for _, a := range req {
		armament = append(armament, model.UpdateSpaceshipArmamentRequest{
			Title:    a.GetTitle(),
			Quantity: a.GetQuantity(),
		})
	}
	return armament
}

func ToProtoUpdateSpaceshipResponse(s model.Spaceship) *spaceshipProto.UpdateSpaceshipResponse {
	return &spaceshipProto.UpdateSpaceshipResponse{
		Spaceship: ToProtoSpaceship(s),
	}
}

func ToProtoSpaceship(s model.Spaceship) *spaceshipProto.Spaceship {
	return &spaceshipProto.Spaceship{
		Id:       s.Id,
		Name:     s.Name,
		Class:    s.Class,
		Crew:     s.Crew,
		ImageUrl: s.Image,
		Value:    s.Value,
		Status:   s.Status,
		Armament: ToProtoSpaceshipToProtoSpaceship(s.Armament),
	}
}

func ToProtoSpaceshipToProtoSpaceship(armaments []model.SpaceshipArmament) []*spaceshipProto.Spaceship_Armament {
	var protoArmaments []*spaceshipProto.Spaceship_Armament
	for _, a := range armaments {
		protoArmaments = append(protoArmaments, &spaceshipProto.Spaceship_Armament{
			Title:    a.Title,
			Quantity: a.Quantity,
		})
	}
	return protoArmaments
}
