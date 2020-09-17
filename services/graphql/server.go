package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/takuya911/go_pf/services/graphql/graph"
	"github.com/takuya911/go_pf/services/graphql/graph/generated"
	pb "github.com/takuya911/go_pf/services/graphql/proto"
	"google.golang.org/grpc"
)

func main() {
	port := "80"
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "user:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	userC := pb.NewUserServiceClient(conn)
	resolver := graph.NewResolver(userC)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
