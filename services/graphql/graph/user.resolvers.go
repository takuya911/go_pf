package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/takuya911/go_pf/services/graphql/graph/generated"
	"github.com/takuya911/go_pf/services/graphql/graph/model"
	user "github.com/takuya911/go_pf/services/graphql/proto"
)

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginPayload, error) {
	res, err := r.userClient.Login(ctx, &user.LoginReq{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}

	return &model.LoginPayload{User: res.User, TokenPair: res.TokenPair}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.CreateUserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserByID(ctx context.Context, input model.GetUserForm) (*user.User, error) {
	request := &user.GetUserForm{
		Id: input.ID,
	}
	return r.userClient.GetUserByID(ctx, request)
}

// Mutation returns generated.MutationResolver implementation.
func (r *resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *resolver }
type queryResolver struct{ *resolver }
