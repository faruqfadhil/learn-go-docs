package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/faruqfadhil/learn-go-docs/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	// doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	doBidirectionalStreaming(c)
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

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	requests := []*calculatorpb.ComputeAverageRequest{
		&calculatorpb.ComputeAverageRequest{
			Number: 1,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 2,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 3,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 4,
		},
	}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("error while calling compute avg %v", err)
	}

	for _, req := range requests {
		fmt.Printf("sending req : %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	msg, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from compute avg %v", err)
	}
	fmt.Printf("response : %v\n", msg)
}

func doBidirectionalStreaming(c calculatorpb.CalculatorServiceClient) {
	reqs := []int32{1, 9, 2, 3, 16, 4, 32}
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("error while creating stream:%v", err)
	}

	waitChan := make(chan struct{})

	go func() {
		for _, req := range reqs {
			fmt.Println("sending req: ", req)
			stream.Send(&calculatorpb.FindMaximumRequest{
				Number: req,
			})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receive data:%v", err)
				break
			}
			fmt.Println("received: ", msg.GetResult())
		}
		close(waitChan)
	}()
	<-waitChan
}
