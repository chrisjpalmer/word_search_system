package main

import (
	"context"
	"log"
	"net"

	"flag"

	wordsearchsystemgrpc "github.com/chrisjpalmer/word_search_system_grpc"
	"google.golang.org/grpc"
)

type wordSearchSystemServer struct {
	wordSearchService *WordSearchService
}

func newWordSearchSystemServer(wordSearchService *WordSearchService) *wordSearchSystemServer {
	_newWordSearchSystemServer := new(wordSearchSystemServer)
	_newWordSearchSystemServer.wordSearchService = wordSearchService
	return _newWordSearchSystemServer
}

func (_wordSearchSystemServer *wordSearchSystemServer) SearchWord(ctx context.Context, in *wordsearchsystemgrpc.SearchWordRequest) (*wordsearchsystemgrpc.SearchWordReply, error) {
	matches := _wordSearchSystemServer.wordSearchService.SearchWord(in.KeyWord)
	return &wordsearchsystemgrpc.SearchWordReply{Matches: matches}, nil
}

func (_wordSearchSystemServer *wordSearchSystemServer) AddWords(ctx context.Context, in *wordsearchsystemgrpc.AddWordsRequest) (*wordsearchsystemgrpc.AddWordsReply, error) {
	err := _wordSearchSystemServer.wordSearchService.AddWords(in.Words)
	if err != nil {
		return nil, err
	}
	return &wordsearchsystemgrpc.AddWordsReply{}, nil
}

func (_wordSearchSystemServer *wordSearchSystemServer) Top5SearchKeyWords(ctx context.Context, in *wordsearchsystemgrpc.Top5SearchKeyWordsRequest) (*wordsearchsystemgrpc.Top5SearchKeyWordsReply, error) {
	keyWords := _wordSearchSystemServer.wordSearchService.Top5SearchKeyWords()
	return &wordsearchsystemgrpc.Top5SearchKeyWordsReply{Keywords: keyWords}, nil
}

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
	wordSearchSystemServer := newWordSearchSystemServer(wordSearchService)

	//Connect the Server, with the proto definitions with the instance of the grpcServer
	wordsearchsystemgrpc.RegisterWordSearchSystemServer(grpcServer, wordSearchSystemServer)

	//Connect the grpcServer with the listener so that it can begin accepting requests
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
