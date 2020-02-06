package main

import "fmt"

import "google.golang.org/grpc"

import "log"

import "github.com/faruqfadhil/learn-go-docs/grpc-go-course/greet/greetpb"

import "context"

import "io"

import "time"

func main() {
	fmt.Println("hello Iam client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	// doUnary(c)
	// doServerStreaming(c)
	doClientStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "faruq",
			LastName:  "fadhil",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC: %v", err)
	}
	log.Printf("response from greet : %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do server streaming RPC...")

	res, err := c.GreetManyTimes(context.Background(), &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "faruq",
			LastName:  "fadhil",
		},
	})
	if err != nil {
		log.Fatalf("error while calling greetManyTimes RPC : %v", err)
	}

	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream %v", err)
		}
		log.Printf("response : %v", msg.GetResult())
	}

}

func doClientStreaming(c greetpb.GreetServiceClient) {
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "faruq",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "fadhil",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "marfuah",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling logGreat %v", err)
	}

	for _, req := range requests {
		fmt.Printf("sending req : %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)

	}

	msg, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from long greet %v", err)
	}
	fmt.Printf("response : %v\n", msg)

}
