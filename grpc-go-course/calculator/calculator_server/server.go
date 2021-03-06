package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"

	"github.com/faruqfadhil/learn-go-docs/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) Calculate(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	res := &calculatorpb.CalculatorResponse{
		Result: req.GetCalculating().GetNumbOne() + req.GetCalculating().GetNumbTwo(),
	}
	return res, nil
}

func (*server) PrimeNumberDecomp(req *calculatorpb.PrimeNumberDecompRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompServer) error {
	k := int32(2)
	N := req.GetPrimeNumber().GetPrime()
	for N > 1 {
		if N%k == 0 {
			res := &calculatorpb.PrimeNumberDecompResponse{
				Result: k,
			}
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
			N = N / k
		} else {
			k = k + 1
		}
	}
	return nil
}

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	var result float32
	divider := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Result: result / float32(divider),
			})
		}
		if err != nil {
			log.Fatalf("error while reading client stream : %v", err)
		}
		number := req.GetNumber()
		divider++
		result += float32(number)
	}
}

func (*server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	maximum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v", err)
			return err
		}
		numb := req.GetNumber()
		if numb > maximum {
			maximum = numb
			err = stream.Send(&calculatorpb.FindMaximumResponse{
				Result: numb,
			})
			if err != nil {
				log.Fatalf("error while sending to client:%v", err)
				return err
			}
		}
	}
}

func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintln("negative number : ", number),
		)
	}
	res := math.Sqrt(float64(number))
	return &calculatorpb.SquareRootResponse{
		Result: res,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

}
