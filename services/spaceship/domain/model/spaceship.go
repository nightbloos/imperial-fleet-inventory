package model

type Spaceship struct {
	ID       int64               `gorm:"column:id;primaryKey"`
	Name     string              `gorm:"column:name"`
	Class    string              `gorm:"column:class"`
	Crew     int64               `gorm:"column:crew"`
	Image    string              `gorm:"column:image"`
	Value    float64             `gorm:"column:value"`
	Status   string              `gorm:"column:status"`
	Armament []SpaceshipArmament `gorm:"foreignKey:spaceship_id"`
}

func (Spaceship) TableName() string {
	return "spaceships"
}

type SpaceshipArmament struct {
	ID          int64  `gorm:"column:id;primaryKey"`
	SpaceshipId int64  `gorm:"column:spaceship_id"`
	Title       string `gorm:"column:title"`
	Quantity    int64  `gorm:"column:quantity"`
}

func (SpaceshipArmament) TableName() string {
	return "spaceship_armament"
}

type GetSpaceshipListQuery struct {
	Name   string
	Class  string
	Status string
}

type CreateSpaceshipRequest struct {
	Name     string
	Class    string
	Crew     int64
	Image    string
	Value    float64
	Status   string
	Armament []CreateSpaceshipArmamentRequest
}

type CreateSpaceshipArmamentRequest struct {
	Title    string
	Quantity int64
}

type UpdateSpaceshipRequest struct {
	Name     string
	Class    string
	Crew     int64
	Image    string
	Value    float64
	Status   string
	Armament []UpdateSpaceshipArmamentRequest
}

type UpdateSpaceshipArmamentRequest struct {
	Title    string
	Quantity int64
}
