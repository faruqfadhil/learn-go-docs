package main

import "google.golang.org/grpc"

import "log"

import "github.com/faruqfadhil/learn-go-docs/grpc-go-course/calculator/calculatorpb"

import "context"

import "io"

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	// doUnary(c)
	doServerStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.CalculatorRequest{
		Calculating: &calculatorpb.Calculating{
			NumbOne: 3,
			NumbTwo: 10,
		},
	}

	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling calculator RPC: %v", err)
	}
	log.Printf("response = %v", res.Result)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	res, err := c.PrimeNumberDecomp(context.Background(), &calculatorpb.PrimeNumberDecompRequest{
		PrimeNumber: &calculatorpb.PrimeNumber{
			Prime: 120,
		},
	})
	if err != nil {
		log.Fatalf("error while calling primeDecomp %v", err)
	}

	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading streaming %v", err)
		}
		log.Printf("response : %v", msg.GetResult())
	}
}
