package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestWordSearchService(t *testing.T) {
	wordSearchService := NewWordSearchService()
	if wordSearchService == nil {
		t.Error("cannot create TestWordSearchService")
	}
}

func TestWordSearchService_SearchWord(t *testing.T) {
	wordSearchService := NewWordSearchService()

	//it should show that the word "hello" exists in the wordSearchService by default
	helloExists := wordSearchService.SearchWord("hello")
	if !helloExists {
		t.Error("'hello' could not be found in the wordSearchService")
	}

	//it should show that the word "yes" exists in the wordSearchService by default
	yesExists := wordSearchService.SearchWord("yes")
	if !yesExists {
		t.Error("'yes' could not be found in the wordSearchService")
	}
}

func TestWordSearchService_AddWord(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		wordSearchService := NewWordSearchService()

		//it should show that the word "super" DOES NOT exist in the wordSearchService by default
		superExistsBefore := wordSearchService.SearchWord("super")
		if superExistsBefore {
			t.Error("'super' already in the wordSearchService")
		}

		//it should show that adding a new word "super" can word without an error
		err := wordSearchService.AddWords([]string{"super"})
		if err != nil {
			t.Errorf("AddWords() returned an error %+v", err)
		}

		//it should show that the word "super" now exists in the wordSearchService
		superExists := wordSearchService.SearchWord("super")
		if !superExists {
			t.Error("'super' could not be found in the wordSearchService")
		}
	})
	t.Run("case sensitivity test", func(t *testing.T) {
		wordSearchService := NewWordSearchService()

		//it should show that the wordSearchService is not sensitive to case for default words
		if !wordSearchService.SearchWord("no") {
			t.Error("searching for 'no' failed")
		}
		if !wordSearchService.SearchWord("NO") {
			t.Error("searching for 'NO' failed")
		}
		if !wordSearchService.SearchWord("nO") {
			t.Error("searching for 'nO' failed")
		}

		//it should show that wordSearchService is not sensitive to case for non-default words
		wordSearchService.AddWords([]string{"GO"})
		if !wordSearchService.SearchWord("go") {
			t.Error("searching for 'go' failed")
		}
		if !wordSearchService.SearchWord("GO") {
			t.Error("searching for 'GO' failed")
		}
		if !wordSearchService.SearchWord("gO") {
			t.Error("searching for 'gO' failed")
		}
	})
	t.Run("add multiple words", func(t *testing.T) {
		wordSearchService := NewWordSearchService()

		//it should allow multiple words to be added at once
		wordSearchService.AddWords([]string{"go", "for", "a", "walk"})
		if !wordSearchService.SearchWord("go") {
			t.Error("searching for 'go' failed")
		}
		if !wordSearchService.SearchWord("for") {
			t.Error("searching for 'for' failed")
		}
		if !wordSearchService.SearchWord("a") {
			t.Error("searching for 'a' failed")
		}
		if !wordSearchService.SearchWord("walk") {
			t.Error("searching for 'walk' failed")
		}
	})
}

func TestWordSearchService_Top5Words(t *testing.T) {
	wordSearchService := NewWordSearchService()
	var (
		top5Words []string
	)

	//it should return the correct list of words when the wordSearchService has just been initialized
	top5Words = wordSearchService.Top5Words()
	assert.EqualValues(t, []string{"filter", "goodbye", "hello", "list", "no"}, top5Words)

	//it should push goodbye to the top of the list when it has been searched more than the other words
	wordSearchService.SearchWord("goodbye")
	top5Words = wordSearchService.Top5Words()
	assert.EqualValues(t, []string{"goodbye", "filter", "hello", "list", "no"}, top5Words)

	//it should push a new word to the top of the list if it has been searched more than the other words
	wordSearchService.AddWords([]string{"zebra"})
	wordSearchService.SearchWord("zebra")
	wordSearchService.SearchWord("zebra")
	top5Words = wordSearchService.Top5Words()
	assert.EqualValues(t, []string{"zebra", "goodbye", "filter", "hello", "list"}, top5Words)
}
