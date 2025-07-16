package grpc

import (
	userpb "github.com/Brrocat/project-protos/proto/user"
	"github.com/Brrocat/users-service/internal/user"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Handler struct {
	svc *user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.User) (*userpb.User, error) {
	log.Printf("Getting user with ID: %d", req.GetId())

	u, err := h.svc.GetUserByID(uint(req.GetId()))
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &userpb.User{
		Id:    uint32(u.ID),
		Email: u.Email,
	}, nil
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	log.Printf("Creating user with email: %s", req.GetEmail())

	u := &user.User{
		Email: req.GetEmail(),
	}

	if err := h.svc.CreateUser(u); err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:    uint32(u.ID),
			Email: u.Email,
		},
	}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	log.Printf("Updating user with ID: %d", req.GetId())

	u := &user.User{
		Email: req.GetEmail(),
	}

	updatedUser, err := h.svc.UpdateUserByID(uint(req.GetId()), u)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return nil, status.Error(codes.Internal, "failed to update user")
	}

	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:    uint32(updatedUser.ID),
			Email: updatedUser.Email,
		},
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	log.Printf("Deleting user with ID: %d", req.GetId())

	if err := h.svc.DeleteUserByID(uint(req.GetId())); err != nil {
		log.Printf("Error deleting user: %v", err)
		return nil, status.Error(codes.Internal, "failed to delete user")
	}

	return &userpb.DeleteUserResponse{
		Success: true,
	}, nil

}

func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	log.Printf("Listing users with limit: %d, offset: %d", req.GetLimit(), req.GetOffset())

	users, err := h.svc.GetAllUsers()
	if err != nil {
		log.Printf("Error listing users: %v", err)
		return nil, status.Error(codes.Internal, "failed to list users")
	}

	var pbUsers []*userpb.User
	for _, u := range users {
		pbUsers = append(pbUsers, &userpb.User{
			Id:    uint32(u.ID),
			Email: u.Email,
		})
	}

	return &userpb.ListUsersResponse{
		Users: pbUsers,
	}, nil
}
