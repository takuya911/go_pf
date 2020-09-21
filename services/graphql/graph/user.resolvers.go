package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/takuya911/go_pf/services/graphql/graph/generated"
	user "github.com/takuya911/go_pf/services/graphql/proto"
)

func (r *queryResolver) GetUser(ctx context.Context) (*user.User, error) {
	return r.userClient.GetUserByID(ctx, &user.GetUserReq{Id: 1})
}

// Query returns generated.QueryResolver implementation.
func (r *resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *resolver }
