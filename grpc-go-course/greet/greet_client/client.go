package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/faruqfadhil/learn-go-docs/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
	// doClientStreaming(c)
	// doBidirectionalStreaming(c)
	doUnaryWithDeadline(c, 5*time.Second) // should be complete
	doUnaryWithDeadline(c, 1*time.Second) // should timeout

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

func doBidirectionalStreaming(c greetpb.GreetServiceClient) {
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error while creatiing stream: %v", err)
		return
	}
	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "faruq",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "fadhil",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "marfuah",
			},
		},
	}

	waitChan := make(chan struct{})
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending req:%v\n", req)
			stream.Send(req)
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
				log.Printf("Error while receive data: %v", err)
				break
			}
			fmt.Printf("received : %v\n", msg.GetResult())

		}
		close(waitChan)

	}()

	<-waitChan

}

func doUnaryWithDeadline(c greetpb.GreetServiceClient, timeOut time.Duration) {
	fmt.Println("starting to do unary with deadline RPC...")
	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "faruq",
			LastName:  "fadhil",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout")
			} else {
				fmt.Println("Unexpected err: ", statusErr)
			}
		} else {
			log.Fatalf("error while calling greetWithDeadline RPC: %v", err)
		}

		return
	}
	log.Printf("response from greetWithDeadline : %v", res.Result)
}
