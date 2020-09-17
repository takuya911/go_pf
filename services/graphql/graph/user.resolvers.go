package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/takuya911/go_pf/services/graphql/graph/generated"
	"github.com/takuya911/go_pf/services/graphql/graph/model"
	pb "github.com/takuya911/go_pf/services/graphql/proto"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	t, err := r.userClient.GetUser(ctx, &pb.GetUserRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	user := &model.User{
		ID:       t.GetId(),
		Name:     t.GetName(),
		Email:    t.GetEmail(),
		Password: t.GetPassword(),
	}
	users := []*model.User{}
	users = append(users, user)

	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *resolver }
type queryResolver struct{ *resolver }
