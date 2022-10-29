package requestMetadata

import "context"

type metadataKey struct{}

type Metadata struct {
	RequestID string
}

func AddToContext(ctx context.Context, md Metadata) context.Context {
	return context.WithValue(ctx, metadataKey{}, md)
}

func GetFromContext(ctx context.Context) Metadata {
	val := ctx.Value(metadataKey{})
	md, ok := val.(Metadata)
	if !ok {
		return Metadata{}
	}
	return md
}

func GetRequestID(ctx context.Context) string {
	return GetFromContext(ctx).RequestID
}
