package grpc

import (
	"context"
	"log"

	pb "personal_hr/models"

	"google.golang.org/grpc"
)

var host = "localhost:1111"

// Ctx ...
var Ctx = context.Background()

// Client ...
var Client pb.UserServiceClient

// Conn ...
var Conn grpc.ClientConn

// ClientConnect ...
func ClientConnect() {
	Conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Not connected err =>", err)
	}
	Client = pb.NewUserServiceClient(Conn)
}
