package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"time"

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

	// 取得
	t, err := userC.GetUser(ctx, &pb.GetUserRequest{Id: "1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	user := &model.User{
		ID:       t.GetId(),
		Name:     t.GetName(),
		Email:    t.GetEmail(),
		Password: t.GetPassword(),
	}
	r.users = append(r.users, user)

	return r.users, nil
}
