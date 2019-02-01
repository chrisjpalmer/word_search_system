package main

import (
	"context"
	"log"
	"net"

	wordsearchsystemgrpc "github.com/chrisjpalmer/word_search_system_grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SearchWord(ctx context.Context, in *wordsearchsystemgrpc.SearchWordRequest) (*wordsearchsystemgrpc.SearchWordReply, error) {
	return nil, nil
}

func (s *server) AddWords(ctx context.Context, in *wordsearchsystemgrpc.AddWordsRequest) (*wordsearchsystemgrpc.AddWordsReply, error) {
	return nil, nil
}

func (s *server) Top5SearchKeyWords(ctx context.Context, in *wordsearchsystemgrpc.Top5SearchKeyWordsRequest) (*wordsearchsystemgrpc.Top5SearchKeyWordsReply, error) {
	return nil, nil
}

func main() {
	//Create new socket connection to listen on the port
	socketConnection, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//Instantiate new grpc server instance class
	grpcServer := grpc.NewServer()

	//Register the WordSearchSystem proto with the grpc server instance
	wordsearchsystemgrpc.RegisterWordSearchSystemServer(grpcServer, &server{})

	//Call server on the grpc server passing the socket connection
	if err := grpcServer.Serve(socketConnection); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
