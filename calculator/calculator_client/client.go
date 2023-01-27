package main

import (
	"fmt"
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Calculator Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := pb.NewCalculatorServiceClient(cc)
	//doSum(c)
	//doPrimes(c)
	//doAvg(c)
	//doMax(c)
	doSqrt(c, 10)
	doSqrt(c, -10)
}
