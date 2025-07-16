package grpc

import (
	userpb "github.com/Brrocat/project-protos/proto/user"
	"log"
	"net"

	"github.com/Brrocat/users-service/internal/user"
	"google.golang.org/grpc"
)

func RunGRPC(svc *user.Service) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	userpb.RegisterUserServiceServer(srv, NewHandler(svc))

	log.Println("Starting gRPC server on :50051")
	return srv.Serve(lis)
}
