package main

import (
	"github.com/Brrocat/users-service/internal/database"
	transportgrpc "github.com/Brrocat/users-service/internal/transport/grpc"
	"github.com/Brrocat/users-service/internal/user"
	"log"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	repo := user.NewRepository(db)
	svc := user.NewService(repo)

	if err := transportgrpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}
