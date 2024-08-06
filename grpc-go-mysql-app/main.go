package main

import (
	"github.com/samyak2405/grpc-go-mysql-app/github.com/samyak2405/grpc-go-mysql-app/proto"
	"log"
	"net"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/samyak2405/grpc-go-mysql-app/server"
	"google.golang.org/grpc"
)

func main() {
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/grpc_db")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	srv := grpc.NewServer()
	proto.RegisterUserServiceServer(srv, &server.Server{DB: db})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Starting gRPC server on port 50051...")
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
