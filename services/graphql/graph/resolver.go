package graph

import (
	"github.com/takuya911/go_pf/services/graphql/graph/generated"
	pb "github.com/takuya911/go_pf/services/graphql/proto"
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
