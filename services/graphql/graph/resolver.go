package graph

import (
	"github.com/takuya911/go_pf/services/graphql/graph/model"
	pb "github.com/takuya911/go_pf/services/graphql/proto"
)

// Resolver struct
type Resolver struct {
	users      []*model.User
	UserClient pb.UserServiceClient
}
