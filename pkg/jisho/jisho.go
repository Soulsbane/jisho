package jisho

import (
	"errors"
	"fmt"

	"github.com/imroc/req/v3"
)

const ApiUrl = "https://jisho.org/api/v1/search/words"

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

func fetchWord(wordToFind string) (JishoResult, error) {
	client := req.C()
	var jishoResult JishoResult

	_, err := client.R().SetQueryParam("keyword", wordToFind).SetSuccessResult(&jishoResult).Get(ApiUrl)

	if err != nil {
		return jishoResult, fmt.Errorf("failed to fetch word: %w", err)
	}

	return jishoResult, nil
}

func LookupWord(wordToFind string) (JishoResult, error) {
	jishoResult, err := fetchWord(wordToFind)

	if err != nil {
		return jishoResult, ErrFailedToLookupWord
	}

	if len(jishoResult.JishoData) > 0 {
		return jishoResult, nil
	} else {
		return jishoResult, ErrNoResults
	}
}
