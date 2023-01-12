package main

import (
	"context"
	"fmt"
	qsort "grpc-case/qsort"
	pb "grpc-case/sorter"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedSorterServer
}

func (s *server) QuickSort(ctx context.Context, in *pb.Numbers) (*pb.Numbers, error) {
	qsort.QuickSort(in.Numbers)
	r := *(in.Index) + 1
	fmt.Println(r)
	in.Index = &r
	return in, nil
}

func main() {
	lis, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterSorterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
