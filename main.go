package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/andreasyunanto/socmed-sgrpc/configs"
	"github.com/andreasyunanto/socmed-sgrpc/database"
	"github.com/andreasyunanto/socmed-sgrpc/pb"
	"github.com/andreasyunanto/socmed-sgrpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// RUN MIGRATOR
	database.Connect()
	database.Migrate(database.DB)

	// START THE SERVER
	lis, err := net.Listen("tcp", ":"+configs.Config("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// START GRPC SERVER
	s := grpc.NewServer()

	// REGISTER SERVICE SERVER
	pb.RegisterPostServiceServer(s, &server.PostServer{})
	pb.RegisterCommentServiceServer(s, &server.CommentServer{})

	// SET REFLECTION FOR DEBUG
	reflection.Register(s)

	go func() {
		fmt.Println("Starting gRPC server socmed...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Stopping gRPC server socmed..")
	s.Stop()
	fmt.Println("Stopping listener...")
	lis.Close()
	fmt.Println("End of Program")

}
