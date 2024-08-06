package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/samyak2405/grpc-go-mysql-app/github.com/samyak2405/grpc-go-mysql-app/proto"
	"github.com/samyak2405/grpc-go-mysql-app/server"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"log"
	"net"
	"os"
)

// Config struct to match the structure of application.yml
type Config struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	GRPC struct {
		Port int `yaml:"port"`
	} `yaml:"grpc"`
}

func main() {
	// Load configuration
	config, err := loadConfig("application.yml")
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Build the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	// Connect to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Start the gRPC server
	srv := grpc.NewServer()
	proto.RegisterUserServiceServer(srv, &server.Server{DB: db})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GRPC.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Starting gRPC server on port %d...", config.GRPC.Port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Function to load configuration from a YAML file
func loadConfig(file string) (*Config, error) {
	config := &Config{}
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	return config, err
}
