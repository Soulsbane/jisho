package jisho

import (
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
)

const ApiUrl = "https://jisho.org/api/v1/search/words"

var NoResultsErr = errors.New("no results found")
var FailedToLookupWordErr = errors.New("failed to lookup word. Try again later")

type JishoResult struct {
	JishoData []struct {
		Slug     string `json:"slug"`
		Japanese []struct {
			Word    string `json:"word"`
			Reading string `json:"reading"`
		} `json:"japanese"`
		Senses []struct {
			EnglishDefinitions []string      `json:"english_definitions"`
			PartsOfSpeech      []string      `json:"parts_of_speech"`
			Links              []interface{} `json:"links"`
			Tags               []interface{} `json:"tags"`
			Restrictions       []interface{} `json:"restrictions"`
			SeeAlso            []interface{} `json:"see_also"`
			Antonyms           []interface{} `json:"antonyms"`
			Source             []interface{} `json:"source"`
			Info               []interface{} `json:"info"`
		} `json:"senses"`
	} `json:"data"`
}

func fetchWord(wordToFind string) (JishoResult, error) {
	client := req.C()
	var jishoResult JishoResult

	_, err := client.R().SetQueryParam("keyword", wordToFind).SetSuccessResult(&jishoResult).Get(ApiUrl)
	return jishoResult, fmt.Errorf("failed to fetch word: %w", err)
}

func LookupWord(wordToFind string) (JishoResult, error) {
	jishoResult, err := fetchWord(wordToFind)

	if err != nil {
		return jishoResult, FailedToLookupWordErr
	}

	if len(jishoResult.JishoData) > 0 {
		return jishoResult, nil
	} else {
		return jishoResult, NoResultsErr
	}
}
