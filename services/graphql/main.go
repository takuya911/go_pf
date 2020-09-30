package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/takuya911/go_pf/services/graphql/graph"
	"github.com/takuya911/go_pf/services/graphql/graph/generated"
	pb "github.com/takuya911/go_pf/services/graphql/proto"
	"google.golang.org/grpc"
)

func main() {
	// get env
	graphqlPort := os.Getenv("GRAPHQL_SERVICE_PORT")
	userService := os.Getenv("USER_SERVICE_NAME")
	userServicePort := os.Getenv("USER_SERVICE_PORT")

	ctx := context.Background()

	// User gRPC connect
	conn, err := grpc.DialContext(ctx, userService+":"+userServicePort, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	userC := pb.NewUserServiceClient(conn)

	// Regist handler
	resolver := graph.NewResolver(userC)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", graphqlPort)
	log.Fatal(http.ListenAndServe(":"+graphqlPort, nil))
}
