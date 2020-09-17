package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/takuya911/go_pf/services/graphql/graph/generated"
	user "github.com/takuya911/go_pf/services/graphql/proto"
)

func (r *queryResolver) GetUser(ctx context.Context) (*user.User, error) {
	return r.userClient.GetUser(ctx, &user.GetUserReq{Id: 1})
}

func (r *userResolver) CreatedAt(ctx context.Context, obj *user.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) UpdatedAt(ctx context.Context, obj *user.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) DeletedAt(ctx context.Context, obj *user.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *resolver) User() generated.UserResolver { return &userResolver{r} }

type queryResolver struct{ *resolver }
type userResolver struct{ *resolver }
