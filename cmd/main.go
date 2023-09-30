package main

import (
	"Authorization_Service/configs"
	"Authorization_Service/internal/pb"
	"Authorization_Service/internal/postgres"
	"Authorization_Service/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg := configs.InitConfig()
	h := postgres.InitDB(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode)

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatal("failed to listen: ", err.Error())
	}

	grpcServer := grpc.NewServer()

	s := service.Server{H: h}
	pb.RegisterAuthServiceServer(grpcServer, &s)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("failed to serve: ", err.Error())
	}
}
