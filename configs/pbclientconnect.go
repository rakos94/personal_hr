package configs

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pbx "personal_hr/pb.client/model"
)

var hostpbclient = "camskoleksi.com:8190"
var CtxClient = context.Background()
var Clients pbx.PersonServiceClient

func ConnPb()  {
	conn, err := grpc.Dial(hostpbclient, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Not connected err =>", err)
	}

	Clients = pbx.NewPersonServiceClient(conn)
}
