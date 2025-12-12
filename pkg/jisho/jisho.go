package jisho

import (
	"encoding/json"
	"errors"
	"net/http"
)

const ApiUrl = "https://jisho.org/api/v1/search/words?keyword="

var ErrNoResults = errors.New("no results found")
var ErrFailedToLookupWord = errors.New("failed to lookup word. Try again later")

type JishoResult struct {
	JishoData []struct {
		Slug     string `json:"slug"`
		Japanese []struct {
			Word    string `json:"word"`
			Reading string `json:"reading"`
		} `json:"japanese"`
		Senses []struct {
			EnglishDefinitions []string `json:"english_definitions"`
			PartsOfSpeech      []string `json:"parts_of_speech"`
			Links              []any    `json:"links"`
			Tags               []any    `json:"tags"`
			Restrictions       []any    `json:"restrictions"`
			SeeAlso            []any    `json:"see_also"`
			Antonyms           []any    `json:"antonyms"`
			Source             []any    `json:"source"`
			Info               []any    `json:"info"`
		} `json:"senses"`
	} `json:"data"`
}

func LookupWord(wordToFind string) (JishoResult, error) {
	var jishoResult JishoResult

	resp, err := http.Get(ApiUrl + wordToFind)

	if err != nil {
		return jishoResult, ErrFailedToLookupWord
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&jishoResult)

	if err != nil {
		return jishoResult, ErrNoResults
	}

	if len(jishoResult.JishoData) > 0 {
		return jishoResult, nil
	} else {
		return jishoResult, ErrNoResults
	}
}
