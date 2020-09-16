package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"time"

	"github.com/takuya911/go_pf/services/graphql/graph/generated"
	"github.com/takuya911/go_pf/services/graphql/graph/model"
	pb "github.com/takuya911/go_pf/services/graphql/proto"
	"google.golang.org/grpc"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*model.User, error) {
	user := &model.User{
		ID:   input.ID,
		Name: input.Name,
	}
	r.users = append(r.users, user)
	return user, nil
}

func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	// gRPC
	conn, err := grpc.Dial("user:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	userC := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	t, err := userC.GetUser(ctx, &pb.GetUserRequest{Id: "1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Print(t)

	// 本題
	return r.users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
