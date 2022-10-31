package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	spaceshipProto "imperial-fleet-inventory/api/langs/go/spaceship/grpc"
	"imperial-fleet-inventory/services/http/api/converter"
	"imperial-fleet-inventory/services/http/domain/model"
)

type SpaceshipServer struct {
	logger       *zap.Logger
	spaceshipSvc spaceshipProto.SpaceshipServiceClient
}

func NewSpaceshipServer(spaceshipClient spaceshipProto.SpaceshipServiceClient, logger *zap.Logger) *SpaceshipServer {
	return &SpaceshipServer{
		spaceshipSvc: spaceshipClient,
		logger:       logger.With(zap.String("server", "spaceship")),
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
	ctx := c.Request.Context()

	query := model.GetSpaceshipListQuery{}
	if c.ShouldBind(&query) != nil {
		s.errorResponse(c, model.NewErrorResponse(http.StatusBadRequest, errors.New("invalid query")))
		return
	}

	spaceshipsRes, err := s.spaceshipSvc.GetSpaceships(ctx, converter.ToProtoGetSpaceshipsRequest(query))
	if err != nil {
		s.logger.Error("failed to get spaceships", zap.Error(err))
		s.errorResponse(c, converter.UnwrapGRPCError(err))
		return
	}

	c.JSON(http.StatusOK, converter.FromProtoGetSpaceshipsResponse(spaceshipsRes))
}

func (s *SpaceshipServer) GetSpaceship(c *gin.Context) {
	ctx := c.Request.Context()

	spaceShipID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		s.errorResponse(c, model.NewErrorResponse(http.StatusBadRequest, errors.New("invalid spaceship id")))
		return
	}

	spaceshipRes, err := s.spaceshipSvc.GetSpaceship(ctx, &spaceshipProto.GetSpaceshipRequest{Id: spaceShipID})
	if err != nil {
		s.logger.Error("failed to get spaceship", zap.Error(err))
		s.errorResponse(c, converter.UnwrapGRPCError(err))
		return
	}

	// Is a good practice to wrap returned objects in some kind of response wrapper, like GetSpaceshipResponse,
	// where requested entity is not on root level, but in some kind of data field.
	// But in initial task, was requested to return entity on root. So that's this code is commented.
	// res := model.GetSpaceshipResponse{Data: converter.FromProtoSpaceship(spaceshipRes.Spaceship)}

	res := converter.FromProtoSpaceship(spaceshipRes.Spaceship)

	c.JSON(http.StatusOK, res)
}

func (s *SpaceshipServer) CreateSpaceship(c *gin.Context) {
	ctx := c.Request.Context()
	var req model.CreateSpaceshipRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Error("invalid request", zap.Error(err))
		s.errorResponse(c, model.NewErrorResponse(http.StatusBadRequest, errors.New("invalid request")))
		return
	}

	_, err := s.spaceshipSvc.CreateSpaceship(ctx, converter.ToProtoCreateSpaceshipRequest(req))
	if err != nil {
		s.logger.Error("failed to create spaceship", zap.Error(err))
		s.errorResponse(c, converter.UnwrapGRPCError(err))
		return
	}

	c.JSON(http.StatusOK, model.CreateSpaceshipResponse{Success: true})
}

func (s *SpaceshipServer) UpdateSpaceship(c *gin.Context) {
	ctx := c.Request.Context()

	spaceShipID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		s.errorResponse(c, model.NewErrorResponse(http.StatusBadRequest, errors.New("invalid spaceship id")))
		return
	}

	var req model.UpdateSpaceshipRequest

	if err = c.ShouldBindJSON(&req); err != nil {
		s.logger.Error("invalid request", zap.Error(err))
		s.errorResponse(c, model.NewErrorResponse(http.StatusBadRequest, errors.New("invalid request")))
		return
	}

	_, err = s.spaceshipSvc.UpdateSpaceship(ctx, converter.ToProtoUpdateSpaceshipRequest(spaceShipID, req))
	if err != nil {
		s.logger.Error("failed to update spaceship", zap.Error(err))
		s.errorResponse(c, converter.UnwrapGRPCError(err))
		return
	}

	c.JSON(http.StatusOK, model.UpdateSpaceshipResponse{Success: true})
}

func (s *SpaceshipServer) DeleteSpaceship(c *gin.Context) {
	ctx := c.Request.Context()

	spaceShipID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		s.errorResponse(c, model.NewErrorResponse(http.StatusBadRequest, errors.New("invalid spaceship id")))
		return
	}

	_, err = s.spaceshipSvc.DeleteSpaceship(ctx, &spaceshipProto.DeleteSpaceshipRequest{Id: spaceShipID})
	if err != nil {
		s.logger.Error("failed to delete spaceship", zap.Error(err))
		s.errorResponse(c, converter.UnwrapGRPCError(err))
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
