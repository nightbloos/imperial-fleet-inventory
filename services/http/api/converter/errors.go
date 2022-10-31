package converter

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	apiModel "imperial-fleet-inventory/services/http/domain/model"
)

// UnwrapGRPCError convert known GRPC errors to HTTP error responses
// If error is unknown then StatusInternalServerError used
func UnwrapGRPCError(err error) apiModel.ErrorResponse {
	st, ok := status.FromError(err)
	if !ok {
		return apiModel.NewErrorResponse(http.StatusInternalServerError, err)
	}

	switch st.Code() {
	case codes.NotFound:
		return apiModel.NewErrorResponseWithMessage(http.StatusNotFound, st.Message())
	default:
		return apiModel.NewErrorResponseWithMessage(http.StatusInternalServerError, st.Message())
	}
}
