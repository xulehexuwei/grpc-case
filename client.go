package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc-case/sorter"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewSorterClient(conn)

	var raw = []int64{1, 8, 3, 5, 6}
	index := int64(2)
	response, err := c.QuickSort(context.Background(), &pb.Numbers{Numbers: raw, Index: &index})
	if err != nil {
		log.Fatalf("Error when calling QuickSort: %s", err)
	}
	fmt.Printf("Response from server: %v；index: %v", response.Numbers, *response.Index)
}
