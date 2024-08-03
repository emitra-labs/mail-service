package main

import (
	"context"
	"fmt"
	"net"

	"github.com/appleboy/graceful"
	"github.com/caitlinelfring/go-env-default"
	"github.com/emitra-labs/common/log"
	"github.com/emitra-labs/mail-service/rpc"
	pb "github.com/emitra-labs/pb/mail"
	"google.golang.org/grpc"
)

var grpcPort = env.GetIntDefault("GRPC_PORT", 4000)

func main() {
	// Setup gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMailServer(grpcServer, &rpc.Server{})

	m := graceful.NewManager()

	m.AddRunningJob(func(ctx context.Context) error {
		log.Infof("gRPC server is listening at %s", lis.Addr())
		return grpcServer.Serve(lis)
	})

	m.AddShutdownJob(func() error {
		grpcServer.GracefulStop()
		return nil
	})

	<-m.Done()
}
