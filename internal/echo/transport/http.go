package transport

import (
	"context"

	"github.com/MishaNiki/prob-gateway/internal/echo/api/pb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type echoServer struct {
}

// NewEchoServer ...
func NewEchoServer() pb.EchoServiceServer {
	return new(echoServer)
}

func (es *echoServer) Echo(ctx context.Context, msg *pb.EchoMessage) (*pb.EchoMessage, error) {
	// Realise busines logic
	return msg, nil
}

// CreateEchoMux ...
func CreateEchoMux(ctx context.Context, endpoint string) (*runtime.ServeMux, error) {
	echoServer := NewEchoServer()
	echoRouter := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterEchoServiceHandlerServer(ctx, echoRouter, echoServer)
	if err != nil {
		return nil, err
	}
	err = pb.RegisterEchoServiceHandlerFromEndpoint(ctx, echoRouter, endpoint, opts)
	if err != nil {
		return nil, err
	}
	return echoRouter, nil
}
