package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordSearchService(t *testing.T) {
	wordSearchService := NewWordSearchService()
	if wordSearchService == nil {
		t.Error("cannot create TestWordSearchService")
	}
}

func TestWordSearchService_SearchWord(t *testing.T) {
	wordSearchService := NewWordSearchService()
	var (
		results []string
	)

	//it should show that the word "hello" exists in the wordSearchService by default
	results = wordSearchService.SearchWord("hello")
	assert.EqualValues(t, []string{"hello"}, results)

	//it should show that the word "yes" exists in the wordSearchService by default
	results = wordSearchService.SearchWord("yes")
	assert.EqualValues(t, []string{"yes"}, results)
}

func TestWordSearchService_AddWord(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		wordSearchService := NewWordSearchService()
		var (
			results []string
		)

		//it should show that the word "super" DOES NOT exist in the wordSearchService by default
		results = wordSearchService.SearchWord("super")
		assert.EqualValues(t, []string{}, results)

		//it should show that adding a new word "super" can word without an error
		err := wordSearchService.AddWords([]string{"super"})
		if err != nil {
			t.Errorf("AddWords() returned an error %+v", err)
		}

		//it should show that the word "super" now exists in the wordSearchService
		results = wordSearchService.SearchWord("super")
		assert.EqualValues(t, []string{"super"}, results)
	})
	t.Run("case sensitivity test", func(t *testing.T) {
		wordSearchService := NewWordSearchService()

		//it should show that the wordSearchService is not sensitive to case for default words
		assert.EqualValues(t, []string{"no"}, wordSearchService.SearchWord("no"))
		assert.EqualValues(t, []string{"no"}, wordSearchService.SearchWord("NO"))
		assert.EqualValues(t, []string{"no"}, wordSearchService.SearchWord("nO"))

		//it should show that wordSearchService is not sensitive to case for non-default words
		wordSearchService.AddWords([]string{"GO"})
		assert.EqualValues(t, []string{"go", "goodbye"}, wordSearchService.SearchWord("go"))
		assert.EqualValues(t, []string{"go", "goodbye"}, wordSearchService.SearchWord("GO"))
		assert.EqualValues(t, []string{"go", "goodbye"}, wordSearchService.SearchWord("gO"))
	})
	t.Run("add multiple words", func(t *testing.T) {
		wordSearchService := NewWordSearchService()

		//it should allow multiple words to be added at once
		wordSearchService.AddWords([]string{"go", "for", "a", "walk"})
		assert.EqualValues(t, []string{"go", "goodbye"}, wordSearchService.SearchWord("go"))
		assert.EqualValues(t, []string{"for"}, wordSearchService.SearchWord("for"))
		assert.EqualValues(t, []string{"a", "search", "walk"}, wordSearchService.SearchWord("a"))
		assert.EqualValues(t, []string{"walk"}, wordSearchService.SearchWord("walk"))
	})
}

func TestWordSearchService_Top5SearchKeyWords(t *testing.T) {
	wordSearchService := NewWordSearchService()
	var (
		top5SearchKeyWords []string
	)

	//it should return the correct list of words when the wordSearchService has just been initialized
	top5SearchKeyWords = wordSearchService.Top5SearchKeyWords()
	assert.EqualValues(t, []string{}, top5SearchKeyWords)

	//it should push "apple" to into the list when it has been searched once
	wordSearchService.SearchWord("apple")
	top5SearchKeyWords = wordSearchService.Top5SearchKeyWords()
	assert.EqualValues(t, []string{"apple"}, top5SearchKeyWords)

	//it should list "bannana" after "apple" when apple and bannana have both been searched once
	wordSearchService.SearchWord("bannana")
	top5SearchKeyWords = wordSearchService.Top5SearchKeyWords()
	assert.EqualValues(t, []string{"apple", "bannana"}, top5SearchKeyWords)

	//it should list "bannana" before "apple" when bannana has been searched one more time than apple
	wordSearchService.SearchWord("bannana")
	top5SearchKeyWords = wordSearchService.Top5SearchKeyWords()
	assert.EqualValues(t, []string{"bannana", "apple"}, top5SearchKeyWords)

	//it should only list 5 words when many have been searched
	wordSearchService.SearchWord("orange")
	wordSearchService.SearchWord("strawaberry")
	wordSearchService.SearchWord("plum")
	wordSearchService.SearchWord("blueberry")
	top5SearchKeyWords = wordSearchService.Top5SearchKeyWords()
	assert.EqualValues(t, []string{"bannana", "apple", "blueberry", "orange", "plum"}, top5SearchKeyWords)
}
