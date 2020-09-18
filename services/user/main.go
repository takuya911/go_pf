package main

import (
	"log"
	"net"
	"os"

	"github.com/takuya911/go_pf/services/user/infrastructure"
	"github.com/takuya911/go_pf/services/user/interface/controllers"
	"github.com/takuya911/go_pf/services/user/interface/repository"
	pb "github.com/takuya911/go_pf/services/user/proto"
	"github.com/takuya911/go_pf/services/user/usecase/interactor"
)

func main() {
	// db connect
	dbConn, err := infrastructure.NewGormConnect()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	lis, err := net.Listen("tcp", ":"+os.Getenv("USER_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// userservice
	userRepository := repository.NewUserRepository(dbConn)
	userInteractor := interactor.NewUserInteractor(userRepository)
	userController := controllers.NewUserController(userInteractor)

	// services regist
	server := infrastructure.NewGrpcServer(userController)
	pb.RegisterUserServiceServer(server, userController)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
