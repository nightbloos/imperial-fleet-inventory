package converter

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"imperial-fleet-inventory/common/sevice/domain"
)

// getGRPCCode tries to map known error types into GRPC error codes
func getGRPCCode(e error) codes.Code {
	var err domain.Error
	if errors.As(e, &err) {
		switch err.GetErrorType() {
		case domain.ErrNotFound:
			return codes.NotFound
		case domain.ErrInternal:
			return codes.Internal
		default:
			return codes.Internal
		}
	}

	return codes.Unknown
}

// CreateGRPCErrorResponse converts error into GRPC error
func CreateGRPCErrorResponse(err error) error {
	if err == nil {
		return status.Error(codes.Unknown, "")
	}
	if _, ok := status.FromError(err); ok {
		return err
	}
	return status.Error(getGRPCCode(err), err.Error())
}

// FromGRPCErrorResponse converts gRPC error into domain error.
// If error is nil or not a domain error, it returns "as is".
// Any gRPC error without existing domain analog returns "as is".
func FromGRPCErrorResponse(err error) error {
	if err == nil {
		return nil
	}
	st, ok := status.FromError(err)
	if !ok {
		return err
	}

	switch st.Code() {
	case codes.NotFound:
		return domain.NewNotFoundError(st.Message())
	case codes.Internal:
		return domain.NewInternalError(st.Message())
	}

	return err
}
