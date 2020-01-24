package main

import "google.golang.org/grpc"

import "log"

import "github.com/faruqfadhil/learn-go-docs/grpc-go-course/calculator/calculatorpb"

import "context"

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
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
