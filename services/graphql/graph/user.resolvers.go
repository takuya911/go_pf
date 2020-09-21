package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/takuya911/go_pf/services/graphql/graph/generated"
	"github.com/takuya911/go_pf/services/graphql/graph/model"
	user "github.com/takuya911/go_pf/services/graphql/proto"
)

func (r *queryResolver) GetUserByID(ctx context.Context, input model.GetUserForm) (*user.User, error) {
	request := &user.GetUserForm{
		Id: input.ID,
	}
	return r.userClient.GetUserByID(ctx, request)

}

// Query returns generated.QueryResolver implementation.
func (r *resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *resolver }
