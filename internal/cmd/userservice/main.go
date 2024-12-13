package main

import (
	"net"
	"os"

	"careem-assignment/internal/database"
	"careem-assignment/internal/logging"
	"careem-assignment/internal/metrics"
	"careem-assignment/internal/repositories"
	"careem-assignment/internal/services"
	"careem-assignment/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	logging.Init()
	db, err := database.DBInit()
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	metrics.StartMetricsServer(":9090")

	userStore := repositories.NewUserStore(db)
	userService := services.NewUserService(userStore, logging.Logger)

	port := os.Getenv("USER_SERVICE_PORT")
	if port == "" {
		port = "50051"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, userService)
	logging.Logger.Infof("UserService listening on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
