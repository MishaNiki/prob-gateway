package transport

import (
	"context"

	"github.com/MishaNiki/prob-gateway/internal/helloworld/api/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type helloServer struct {
}

// NewHelloServer ...
func NewHelloServer() pb.HelloWorldServiceServer {
	return new(helloServer)
}

func (hs *helloServer) HelloWorld(ctx context.Context, em *empty.Empty) (*pb.HelloWorldResponse, error) {
	// Realise busines logic
	return &pb.HelloWorldResponse{Message: "Hello world"}, nil
}

// CreateHelloMux ...
func CreateHelloMux(ctx context.Context, endpoint string) (*runtime.ServeMux, error) {
	echoServer := NewHelloServer()
	echoRouter := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterHelloWorldServiceHandlerServer(ctx, echoRouter, echoServer)
	if err != nil {
		return nil, err
	}
	err = pb.RegisterHelloWorldServiceHandlerFromEndpoint(ctx, echoRouter, endpoint, opts)
	if err != nil {
		return nil, err
	}
	return echoRouter, nil
}
