package main

import (
	"context"
	"flag"
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
	env := os.Getenv("ENV")

	ctx := context.Background()

	// User gRPC connect
	var opts []grpc.DialOption

	if "dev" == env {
		opts = append(opts, grpc.WithInsecure())
	} else {
		serverHost := flag.String("server-host", "", "https://user-repftyfivq-an.a.run.app")
		opts = append(opts, grpc.WithAuthority(*serverHost))
	}
	conn, err := grpc.DialContext(ctx, userService+":"+userServicePort, opts...)
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
