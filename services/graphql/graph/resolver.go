package graph

import (
	"github.com/takuya911/gqlgen-grpc/services/graphql/graph/generated"
	pb "github.com/takuya911/gqlgen-grpc/services/graphql/proto"
)

// Resolver struct
type resolver struct {
	userClient pb.UserServiceClient
}

// NewResolver function
func NewResolver(u pb.UserServiceClient) generated.ResolverRoot {
	return &resolver{
		userClient: u,
	}
}
