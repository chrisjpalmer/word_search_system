package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

//
//HELPER STRUCTURES
//

//wordInfo - represents a word and any of its associated metadata
type wordInfo struct {
	word                  string
	numberOfTimesSearched int64
}

//wordInfoSlice - a slice of wordInfo struct
type wordInfoSlice []wordInfo

//alphabeticalWordInfoSlice - a wordInfoSlice that can be passed to sort.Sort() to sort words that are alphabetically higher to the start of the slice
type alphabeticalWordInfoSlice wordInfoSlice

func (p alphabeticalWordInfoSlice) Len() int { return len(p) }
func (p alphabeticalWordInfoSlice) Less(i, j int) bool {
	return p[i].word < p[j].word
}
func (p alphabeticalWordInfoSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

//searchFrequencyWordInfoSlice - a wordInfoSlice that can be passed to sort.Sort() to sort words that are more frequently searched to the start of the slice
type searchFrequencyWordInfoSlice wordInfoSlice

func (p searchFrequencyWordInfoSlice) Len() int { return len(p) }
func (p searchFrequencyWordInfoSlice) Less(i, j int) bool {
	return p[i].numberOfTimesSearched > p[j].numberOfTimesSearched
}
func (p searchFrequencyWordInfoSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

//wordStat - metadata for a word
type wordStat struct {
	numberOfTimesSearched int64
}

//
//WordSearchService
//

//WordSearchService - a service which allows words to be added and searched, as well as providing statistics on those words
type WordSearchService struct {
	dictionaryWords     map[string]bool
	dictionaryWordStats map[string]*wordStat
}

//NewWordSearchService creates a new instance of WordSearchService
func NewWordSearchService() *WordSearchService {
	newWordSearchService := new(WordSearchService)
	newWordSearchService.dictionaryWords = make(map[string]bool)
	newWordSearchService.dictionaryWordStats = make(map[string]*wordStat)
	newWordSearchService.AddWords([]string{
		"hello",
		"goodbye",
		"list",
		"search",
		"filter",
		"yes",
		"no",
	})
	return newWordSearchService
}

//SearchWord - return whether word exists or not
func (wordSearchService *WordSearchService) SearchWord(word string) (exists bool) {
	lowercaseWord := strings.ToLower(word)

	//Check if the word does not exist
	if wordSearchService.dictionaryWords[lowercaseWord] == false {
		return false
	}

	//The word exists, record that it has been searched
	wordSearchService.incrementNumberOfTimesSearched(lowercaseWord)
	return true
}

//incrementNumberOfTimesSearched - increment the wordStat numberOfTimesSearched property for a given word
func (wordSearchService *WordSearchService) incrementNumberOfTimesSearched(word string) {
	if wordSearchService.dictionaryWordStats[word] == nil {
		wordSearchService.dictionaryWordStats[word] = new(wordStat)
	}
	wordSearchService.dictionaryWordStats[word].numberOfTimesSearched++
}

//AddWords - add words to the list
func (wordSearchService *WordSearchService) AddWords(words []string) (err error) {
	//Convert all words to lowercase before adding them
	lowercaseWords := wordSearchService.wordsToLowercase(words)

	//Validation... do any of the words exist already?
	for i := range lowercaseWords {
		word := lowercaseWords[i]
		if wordSearchService.dictionaryWords[word] == true {
			return errors.New(fmt.Sprintf("%s word already exists", word))
		}
	}

	//Add each of these words to the words
	for i := range lowercaseWords {
		word := lowercaseWords[i]
		wordSearchService.dictionaryWords[word] = true
	}

	return nil
}

//TODO: COMMENT
func (wordSearchService *WordSearchService) wordsToLowercase(words []string) (lowercaseWords []string) {
	_lowercaseWords := make([]string, len(words))
	for i := range words {
		_lowercaseWords[i] = strings.ToLower(words[i])
	}
	return _lowercaseWords
}

//Top5Words - sort the words alphabetically and by order of search frequency, return the top 5
func (wordSearchService *WordSearchService) Top5Words() []string {
	//Extract all the words from the dictionary and put in an array
	_wordInfoSlice := make(wordInfoSlice, len(wordSearchService.dictionaryWords))
	i := 0
	for word := range wordSearchService.dictionaryWords {
		_wordInfoSlice[i].word = word
		if wordSearchService.dictionaryWordStats[word] != nil {
			_wordInfoSlice[i].numberOfTimesSearched = wordSearchService.dictionaryWordStats[word].numberOfTimesSearched
		}
		i++
	}

	//Sort alphabetically
	_alphabeticalWordInfoSlice := alphabeticalWordInfoSlice(_wordInfoSlice)
	sort.Sort(_alphabeticalWordInfoSlice)

	//Sort by search frequency
	_searchFrequencyWordInfoSlice := searchFrequencyWordInfoSlice(_alphabeticalWordInfoSlice)
	sort.Sort(_searchFrequencyWordInfoSlice)

	sliceMax := 5
	if sliceMax > len(_searchFrequencyWordInfoSlice) {
		sliceMax = len(_searchFrequencyWordInfoSlice)
	}

	top5WordInfoSlice := _searchFrequencyWordInfoSlice[:sliceMax]

	top5Words := make([]string, len(top5WordInfoSlice))
	for i := range top5WordInfoSlice {
		top5Words[i] = top5WordInfoSlice[i].word
	}

	return top5Words
}

/**
write unit tests for this struct

*/
