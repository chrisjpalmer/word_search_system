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

//keyWordStat - represents a word and any of its associated metadata
type keyWordStat struct {
	word                  string
	numberOfTimesSearched int64
}

//keyWordStatSlice - a slice of keyWordStat struct
type keyWordStatSlice []*keyWordStat

//alphabeticalKeyWordStatSlice - a keyWordStatSlice that can be passed to sort.Sort() to sort words that are alphabetically higher to the start of the slice
type alphabeticalKeyWordStatSlice keyWordStatSlice

func (p alphabeticalKeyWordStatSlice) Len() int { return len(p) }
func (p alphabeticalKeyWordStatSlice) Less(i, j int) bool {
	return p[i].word < p[j].word
}
func (p alphabeticalKeyWordStatSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

//searchFrequencyKeyWordStatSlice - a keyWordStatSlice that can be passed to sort.Sort() to sort words that are more frequently searched to the start of the slice
type searchFrequencyKeyWordStatSlice keyWordStatSlice

func (p searchFrequencyKeyWordStatSlice) Len() int { return len(p) }
func (p searchFrequencyKeyWordStatSlice) Less(i, j int) bool {
	return p[i].numberOfTimesSearched > p[j].numberOfTimesSearched
}
func (p searchFrequencyKeyWordStatSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

//
//WordSearchService
//

//WordSearchService - a service which allows words to be added and searched, as well as providing statistics on those words
type WordSearchService struct {
	dictionaryWords map[string]bool
	keyWordStatsMap map[string]*keyWordStat
	keyWordStats    []*keyWordStat
}

//NewWordSearchService creates a new instance of WordSearchService
func NewWordSearchService() *WordSearchService {
	newWordSearchService := new(WordSearchService)
	newWordSearchService.dictionaryWords = make(map[string]bool)
	newWordSearchService.keyWordStatsMap = make(map[string]*keyWordStat)
	newWordSearchService.keyWordStats = make([]*keyWordStat, 0, 0)
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

//SearchWord - returns possible matches for the keyword provided
func (wordSearchService *WordSearchService) SearchWord(keyWord string) (matches []string) {
	//convert the keyword to lowercase
	lowercaseKeyWord := strings.ToLower(keyWord)

	//record the the key word as being searches
	wordSearchService.recordKeyWord(lowercaseKeyWord)

	//Check if the word does not exist
	possibleMatches := make([]string, 0, len(wordSearchService.dictionaryWords))
	for dictionaryWord := range wordSearchService.dictionaryWords {
		if strings.Contains(dictionaryWord, lowercaseKeyWord) {
			possibleMatches = append(possibleMatches, dictionaryWord)
		}
	}

	//Order the matches alphabetically
	sort.Strings(possibleMatches)

	return possibleMatches
}

//recordKeyWord - increment the wordStat numberOfTimesSearched property for a given search keyword
func (wordSearchService *WordSearchService) recordKeyWord(lowercaseKeyWord string) {
	if wordSearchService.keyWordStatsMap[lowercaseKeyWord] == nil {
		keyWordStat := new(keyWordStat)
		keyWordStat.word = lowercaseKeyWord
		wordSearchService.keyWordStatsMap[lowercaseKeyWord] = keyWordStat
		wordSearchService.keyWordStats = append(wordSearchService.keyWordStats, keyWordStat)
	}
	wordSearchService.keyWordStatsMap[lowercaseKeyWord].numberOfTimesSearched++
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

//wordsToLowercase - converts the []string to a lowercase []string
func (wordSearchService *WordSearchService) wordsToLowercase(words []string) (lowercaseWords []string) {
	_lowercaseWords := make([]string, len(words))
	for i := range words {
		_lowercaseWords[i] = strings.ToLower(words[i])
	}
	return _lowercaseWords
}

//Top5SearchKeyWords - returns the top 5 most searched keywords
func (wordSearchService *WordSearchService) Top5SearchKeyWords() []string {
	//Clone the source slice
	keyWordStats := append(wordSearchService.keyWordStats[:0:0], wordSearchService.keyWordStats...)

	//Sort alphabetically
	_alphabeticalWordInfoSlice := alphabeticalKeyWordStatSlice(keyWordStats)
	sort.Sort(_alphabeticalWordInfoSlice)

	//Sort by search frequency
	_searchFrequencyWordInfoSlice := searchFrequencyKeyWordStatSlice(_alphabeticalWordInfoSlice)
	sort.Sort(_searchFrequencyWordInfoSlice)

	//Get first 5 words or less of the sorted words
	sliceMax := 5
	if sliceMax > len(_searchFrequencyWordInfoSlice) {
		sliceMax = len(_searchFrequencyWordInfoSlice)
	}
	top5WordInfoSlice := _searchFrequencyWordInfoSlice[:sliceMax]
	top5Words := make([]string, len(top5WordInfoSlice))
	for i := range top5WordInfoSlice {
		top5Words[i] = top5WordInfoSlice[i].word
	}

	//return the result
	return top5Words
}
