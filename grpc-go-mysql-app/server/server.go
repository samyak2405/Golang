package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/samyak2405/grpc-go-mysql-app/github.com/samyak2405/grpc-go-mysql-app/proto"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/samyak2405/grpc-go-mysql-app/proto"
)

type Server struct {
	proto.UnimplementedUserServiceServer
	DB *sql.DB
}

func (s *Server) GetUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	var user proto.User
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := s.DB.QueryRow(query, req.Id)
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve user: %v", err)
	}
	return &proto.UserResponse{User: &user}, nil
}
