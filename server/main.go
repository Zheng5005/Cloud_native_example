package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	calculatorpb "example.com/cloud-native-grpc-calculator/gen/example.com/cloud-native-grpc-calculator/gen/calculator"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	return &calculatorpb.AddResponse{Result: req.A + req.B}, nil
}

func (s *server) Sub(ctx context.Context, req *calculatorpb.SubRequest) (*calculatorpb.SubResponse, error) {
	return &calculatorpb.SubResponse{Result: req.A - req.B}, nil
}

func (s *server) Mul(ctx context.Context, req *calculatorpb.MulRequest) (*calculatorpb.MulResponse, error) {
	return &calculatorpb.MulResponse{Result: req.A * req.B}, nil
}

func (s *server) Div(ctx context.Context, req *calculatorpb.DivRequest) (*calculatorpb.DivResponse, error) {
	return &calculatorpb.DivResponse{Result: req.A / req.B}, nil
}

func main() {
	addr := getEnv("LISTEN_ADDR", ":50051")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", addr, err)
	}

	grpcServer := grpc.NewServer()
	calculatorpb.RegisterCalculatorServer(grpcServer, &server{})

	log.Printf("gRPC server listening on %s", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
