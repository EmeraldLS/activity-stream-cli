package main

import (
	"log"
	"net"

	pb "github.com/EmeraldLS/daily_activity/proto"
	"github.com/EmeraldLS/daily_activity/server/activity"
	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("An error occured when starting the listerner :: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterActivityServiceServer(grpcServer, &activity.ActivityServer{})
	log.Printf("Acrivity Service started @ :: %v", list.Addr())
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("An error occured while serving the gRPC service %v", err)
	}
}
