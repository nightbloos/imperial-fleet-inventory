package model

type GetSpaceshipListQuery struct {
	Name   string `form:"name"`
	Class  string `form:"class"`
	Status string `form:"status"`
}

type GetSpaceshipsListResponse struct {
	Data []SpaceshipListShort `json:"data"`
}

type SpaceshipListShort struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type GetSpaceshipResponse struct {
	Data Spaceship `json:"data"`
}

type CreateSpaceshipRequest struct {
	Name     string              `json:"name"`
	Class    string              `json:"class"`
	Crew     int64               `json:"crew"`
	Image    string              `json:"image"`
	Value    float64             `json:"value"`
	Status   string              `json:"status"`
	Armament []SpaceshipArmament `json:"armament"`
}

type CreateSpaceshipResponse struct {
	Success bool `json:"success"`
}

type UpdateSpaceshipRequest struct {
	Name     string              `json:"name"`
	Class    string              `json:"class"`
	Crew     int64               `json:"crew"`
	Image    string              `json:"image"`
	Value    float64             `json:"value"`
	Status   string              `json:"status"`
	Armament []SpaceshipArmament `json:"armament"`
}

type UpdateSpaceshipResponse struct {
	Success bool `json:"success"`
}

type DeleteSpaceshipResponse struct {
	Success bool `json:"success"`
}

type Spaceship struct {
	ID       int64               `json:"id"`
	Name     string              `json:"name"`
	Class    string              `json:"class"`
	Crew     int64               `json:"crew"`
	Image    string              `json:"image"`
	Value    float64             `json:"value"`
	Status   string              `json:"status"`
	Armament []SpaceshipArmament `json:"armament"`
}

type SpaceshipArmament struct {
	Title string `json:"title"`
	Qty   int64  `json:"message_count,string"`
}
