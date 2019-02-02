package main

import (
	"log"
	"net"

	"flag"

	wordsearchsystemgrpc "github.com/chrisjpalmer/word_search_system_grpc"
	"google.golang.org/grpc"
)

func main() {
	var (
		err      error
		config   *Config
		listener net.Listener
	)
	//Get config
	configPath := flag.String("config", "config.json", "/path/to/config.json")
	flag.Parse()
	config, err = ParseConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//Announce start
	log.Println("WordSearchSystem has started")

	//Create the word search service
	wordSearchService := NewWordSearchService()

	//Create the listener for the specific address
	listener, err = net.Listen("tcp", config.ListenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//Instantiate new grpc server instance object
	grpcServer := grpc.NewServer()

	//Create the WordSearchSystemServer
	wordSearchSystemServer := NewWordSearchSystemServer(wordSearchService)

	//Connect the Server, with the proto definitions with the instance of the grpcServer
	wordsearchsystemgrpc.RegisterWordSearchSystemServer(grpcServer, wordSearchSystemServer)

	//Connect the grpcServer with the listener so that it can begin accepting requests
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
