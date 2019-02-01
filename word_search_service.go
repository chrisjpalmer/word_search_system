package main

import (
	"fmt"
	"sort"

	"github.com/pkg/errors"
)

type WordStat struct {
	numberOfTimesSearched int64
}

type WordSearchService struct {
	dictionaryWords     map[string]bool
	dictionaryWordStats map[string]*WordStat
}

type WordInfo struct {
	word                  string
	numberOfTimesSearched int64
}

type WordInfoSlice []WordInfo

type AlphabeticalWordInfoSlice WordInfoSlice

func (p AlphabeticalWordInfoSlice) Len() int { return len(p) }
func (p AlphabeticalWordInfoSlice) Less(i, j int) bool {
	return p[i].word < p[j].word
}
func (p AlphabeticalWordInfoSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type SearchFrequencyWordInfoSlice WordInfoSlice

func (p SearchFrequencyWordInfoSlice) Len() int { return len(p) }
func (p SearchFrequencyWordInfoSlice) Less(i, j int) bool {
	return p[i].numberOfTimesSearched < p[j].numberOfTimesSearched
}
func (p SearchFrequencyWordInfoSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func NewWordSearchService() *WordSearchService {
	newWordSearchService := new(WordSearchService)
	newWordSearchService.dictionaryWords = make(map[string]bool)
	newWordSearchService.dictionaryWordStats = make(map[string]*WordStat)
	newWordSearchService.AddWords([]string{
		"hello",
		"goodbye",
		"list",
		"search",
		"filter",
		"yes",
		"no ",
	})
	return newWordSearchService
}

//SearchWord - return whether word exists or not
func (wordSearchService *WordSearchService) SearchWord(word string) (exists bool) {
	//Check if the word does not exist
	if wordSearchService.dictionaryWords[word] == false {
		return false
	}

	//The word exists, record that it has been searched
	wordSearchService.incrementNumberOfTimesSearched(word)
	return true
}

//incrementNumberOfTimesSearched - increment the wordStat numberOfTimesSearched property for a given word
func (wordSearchService *WordSearchService) incrementNumberOfTimesSearched(word string) {
	if wordSearchService.dictionaryWordStats[word] == nil {
		wordSearchService.dictionaryWordStats[word] = new(WordStat)
	}
	wordSearchService.dictionaryWordStats[word].numberOfTimesSearched++
}

//AddWords - add words to the list
func (wordSearchService *WordSearchService) AddWords(words []string) (err error) {
	//Validation... do any of the words exist already?
	for i := range words {
		word := words[i]
		if wordSearchService.dictionaryWords[word] == true {
			return errors.New(fmt.Sprintf("%s word already exists", word))
		}
	}

	//Add each of these words to the words
	for i := range words {
		word := words[i]
		wordSearchService.dictionaryWords[word] = true
	}

	return nil
}

//Top5Words - sort the words alphabetically and by order of search frequency, return the top 5
func (wordSearchService *WordSearchService) Top5Words() []string {
	//Extract all the words from the dictionary and put in an array
	wordInfoSlice := make(WordInfoSlice, len(wordSearchService.dictionaryWords))
	i := 0
	for word := range wordSearchService.dictionaryWords {
		wordInfoSlice[i].word = word
		if wordSearchService.dictionaryWordStats[word] != nil {
			wordInfoSlice[i].numberOfTimesSearched = wordSearchService.dictionaryWordStats[word].numberOfTimesSearched
		}
		i++
	}

	//Sort alphabetically
	alphabeticalWordInfoSlice := AlphabeticalWordInfoSlice(wordInfoSlice)
	sort.Sort(alphabeticalWordInfoSlice)

	//Sort by search frequency
	searchFrequencyWordInfoSlice := SearchFrequencyWordInfoSlice(alphabeticalWordInfoSlice)
	sort.Sort(searchFrequencyWordInfoSlice)

	sliceMax := 5
	if sliceMax > len(searchFrequencyWordInfoSlice) {
		sliceMax = len(searchFrequencyWordInfoSlice)
	}

	top5WordInfoSlice := searchFrequencyWordInfoSlice[:sliceMax]

	top5Words := make([]string, len(top5WordInfoSlice))
	for i := range top5WordInfoSlice {
		top5Words[i] = top5WordInfoSlice[i].word
	}

	return top5Words
}

/**
TODO: need to comment the above code nicely
write unit tests for this struct

*/
