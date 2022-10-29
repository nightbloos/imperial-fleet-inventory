package grpc

import (
	"context"

	"google.golang.org/grpc"
	grpcMetadata "google.golang.org/grpc/metadata"

	"imperial-fleet-inventory/common/request_metadata"
)

const (
	metadataRequestID string = "app-request-id"
)

func NewClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		md := requestMetadata.GetFromContext(ctx)
		ctx = grpcMetadata.AppendToOutgoingContext(ctx,
			metadataRequestID, md.RequestID,
		)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func NewGRPCServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		grpcMD, ok := grpcMetadata.FromIncomingContext(ctx)
		if ok {
			ctx = requestMetadata.AddToContext(
				ctx,
				requestMetadata.Metadata{
					RequestID: firstStringOrEmpty(grpcMD.Get(metadataRequestID)),
				})
		}

		return handler(ctx, req)
	}
}

func firstStringOrEmpty(strSlice []string) string {
	if len(strSlice) == 0 {
		return ""
	}

	return strSlice[0]
}
