package main

import "fmt"

import "google.golang.org/grpc"

import "log"

import "github.com/faruqfadhil/learn-go-docs/grpc-go-course/greet/greetpb"

import "context"

func main() {
	fmt.Println("hello Iam client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
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
