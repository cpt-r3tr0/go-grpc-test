package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/cpt-r3tr0/go-grpc-test/proto"
)

const (
	port = ":50051"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorClient(conn)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter expression (q to quit): ")

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		line = strings.TrimSpace(line)
		if line == "q" || line == "quit" {
			break
		}

		req := &pb.CalculationRequest{Expression: line}
		res, err := c.Calculate(context.Background(), req)

		if err != nil {
			log.Fatalf("Error calculating: %v", err)
		}

		log.Printf("Result: %d", res.Result)
	}
}
