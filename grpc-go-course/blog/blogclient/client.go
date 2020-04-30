package main

import (
	"context"
	"fmt"
	"log"

	"github.com/faruqfadhil/learn-go-docs/grpc-go-course/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello Iam client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	// doBidirectionalStreaming(c)
	// doUnaryWithDeadline(c, 5*time.Second) // should be complete
	// doUnaryWithDeadline(c, 1*time.Second) // should timeout

}

func doUnary(c blogpb.BlogServiceClient) {
	fmt.Println("starting to do unary RPC...")
	req := &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			Id:       "sas",
			AuthorId: "fadhil",
			Content:  "ehehe",
			Title:    "s.Tr.Kom",
		},
	}
	res, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC: %v", err)
	}
	log.Printf("response from greet : %v", res.GetBlog())
}
