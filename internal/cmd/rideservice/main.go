package main

import (
	"log"
	"net"
	"os"

	"careem-assignment/internal/database"
	"careem-assignment/internal/logging"
	"careem-assignment/internal/metrics"
	"careem-assignment/internal/repositories"
	"careem-assignment/internal/services"
	"careem-assignment/proto"
	"google.golang.org/grpc"
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

	metrics.StartMetricsServer(":9092")

	ridesRepo := repositories.NewRidesRepository(db)
	ridesService := services.NewRidesService(ridesRepo, logging.Logger)

	port := os.Getenv("RIDE_SERVICE_PORT")
	if port == "" {
		port = "50053"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterRidesServiceServer(grpcServer, ridesService)
	logging.Logger.Infof("RidesService listening on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
