package main

import (
	"job-application/cmd"
	handler_grpc "job-application/handlers/grpc"
	pb "job-application/proto"
	"job-application/usecase"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		return
	}

	opt := []grpc.ServerOption{}

	server := grpc.NewServer(opt...)

	repository, err := cmd.InitRepository()
	if err != nil {
		return
	}
	usecase := usecase.NewUsecase(*repository)

	userHandler := handler_grpc.NewUserHandler(usecase.UserUsecase)
	jobHandler := handler_grpc.NewJobHandler(usecase.JobUsecase)
	applicationHandler := handler_grpc.NewApplicationHandler(usecase.ApplicationUsecase)

	pb.RegisterUserServiceServer(server, userHandler)
	pb.RegisterJobServiceServer(server, jobHandler)
	pb.RegisterApplicationServiceServer(server, applicationHandler)

	if err = server.Serve(listener); err != nil {
		return
	}

}
