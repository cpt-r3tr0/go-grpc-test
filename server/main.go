package main

import (
	"context"
	"log"
	"net"

	pb "github.com/cpt-r3tr0/go-grpc-test/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

// Calculate implements proto.CalculatorServer.
func (s *server) Calculate(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	log.Printf("Received: %v", req.GetExpression())

	// Evaluate expression
	result := EvaluateExp(req.GetExpression())

	res := &pb.CalculationResponse{
		Result: int32(result),
	}
	return res, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterCalculatorServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
