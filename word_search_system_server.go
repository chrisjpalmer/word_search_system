package main

import (
	"context"

	wordsearchsystemgrpc "github.com/chrisjpalmer/word_search_system_grpc"
)

//WordSearchSystemServer - an struct which implements the wordsearchsystemgrpc.WordSearchSystemServer interface to handle gRPC requests.
// It handles the requests and executes logic on the wordSearchService object which is the brain of the application
type WordSearchSystemServer struct {
	wordSearchService *WordSearchService
}

//NewWordSearchSystemServer - initializes a new wordSearchSystemServer
func NewWordSearchSystemServer(wordSearchService *WordSearchService) *WordSearchSystemServer {
	_newWordSearchSystemServer := new(WordSearchSystemServer)
	_newWordSearchSystemServer.wordSearchService = wordSearchService
	return _newWordSearchSystemServer
}

//SearchWord - handles SearchWord request to search for words in the words list
func (wordSearchSystemServer *WordSearchSystemServer) SearchWord(ctx context.Context, in *wordsearchsystemgrpc.SearchWordRequest) (*wordsearchsystemgrpc.SearchWordReply, error) {
	matches := wordSearchSystemServer.wordSearchService.SearchWord(in.KeyWord)
	return &wordsearchsystemgrpc.SearchWordReply{Matches: matches}, nil
}

//AddWords - handles the AddWords request to add words to the words list
func (wordSearchSystemServer *WordSearchSystemServer) AddWords(ctx context.Context, in *wordsearchsystemgrpc.AddWordsRequest) (*wordsearchsystemgrpc.AddWordsReply, error) {
	err := wordSearchSystemServer.wordSearchService.AddWords(in.Words)
	if err != nil {
		return nil, err
	}
	return &wordsearchsystemgrpc.AddWordsReply{}, nil
}

//Top5SearchKeyWords - handles the Top5SearchKeyWords to get the top 5 keywords that were searched
func (wordSearchSystemServer *WordSearchSystemServer) Top5SearchKeyWords(ctx context.Context, in *wordsearchsystemgrpc.Top5SearchKeyWordsRequest) (*wordsearchsystemgrpc.Top5SearchKeyWordsReply, error) {
	keyWords := wordSearchSystemServer.wordSearchService.Top5SearchKeyWords()
	return &wordsearchsystemgrpc.Top5SearchKeyWordsReply{Keywords: keyWords}, nil
}
