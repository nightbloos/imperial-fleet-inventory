package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"imperial-fleet-inventory/services/http/domain/model"
)

type SpaceshipServer struct {
	logger *zap.Logger
}

func NewSpaceshipServer(logger *zap.Logger) *SpaceshipServer {
	return &SpaceshipServer{
		logger: logger.With(zap.String("server", "spaceship")),
	}
}

func (s *SpaceshipServer) Register(router *gin.Engine) {
	transactionsRoutes := router.Group("/spaceships")
	{
		transactionsRoutes.GET("/", s.GetSpaceships)
		transactionsRoutes.POST("/", s.CreateSpaceship)
		transactionsRoutes.GET("/:id", s.GetSpaceship)
		transactionsRoutes.PUT("/:id", s.UpdateSpaceship)
		transactionsRoutes.DELETE("/:id", s.DeleteSpaceship)
	}
}

func (s *SpaceshipServer) GetSpaceships(c *gin.Context) {
	query := model.GetSpaceshipListQuery{}
	if c.ShouldBind(&query) != nil {
		s.errorResponse(c, model.NewErrorResponse(http.StatusBadRequest, errors.New("invalid query")))
		return
	}

	res := model.GetSpaceshipsListResponse{
		Data: []model.SpaceshipListShort{
			{
				ID:     1,
				Name:   "Devastator",
				Status: "Operational",
			},
			{
				ID:     2,
				Name:   "Red Five",
				Status: "Damaged",
			},
		},
	}

	c.JSON(http.StatusOK, res)
}

func (s *SpaceshipServer) GetSpaceship(c *gin.Context) {
	spaceshipID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		s.errorResponse(c, model.NewErrorResponse(http.StatusBadRequest, errors.New("invalid spaceship id")))
		return
	}

	res := model.Spaceship{
		Id:     spaceshipID,
		Name:   "Devastator",
		Class:  "Star Destroyer",
		Crew:   35000,
		Image:  `https:\\url.to.image`,
		Value:  1999.99,
		Status: "operational",
		Armament: []model.SpaceshipArmament{
			{
				Title: "Turbo Laser",
				Qty:   60,
			},
		},
	}

	c.JSON(http.StatusOK, res)
}

func (s *SpaceshipServer) CreateSpaceship(c *gin.Context) {

	res := model.CreateSpaceshipResponse{
		Success: true,
	}
	c.JSON(http.StatusOK, res)
}

func (s *SpaceshipServer) UpdateSpaceship(c *gin.Context) {
	_, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		s.errorResponse(c, model.NewErrorResponse(http.StatusBadRequest, errors.New("invalid spaceship id")))
		return
	}

	res := model.UpdateSpaceshipResponse{
		Success: true,
	}
	c.JSON(http.StatusOK, res)
}

func (s *SpaceshipServer) DeleteSpaceship(c *gin.Context) {
	_, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		s.errorResponse(c, model.NewErrorResponse(http.StatusBadRequest, errors.New("invalid spaceship id")))
		return
	}

	res := model.DeleteSpaceshipResponse{
		Success: true,
	}
	c.JSON(http.StatusOK, res)
}

func (s *SpaceshipServer) errorResponse(c *gin.Context, errResp model.ErrorResponse) {
	c.AbortWithStatusJSON(errResp.Code, errResp)
}
