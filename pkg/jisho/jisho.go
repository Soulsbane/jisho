package jisho

import (
	"fmt"

	"github.com/imroc/req/v3"
)

const API_URL = "https://jisho.org/api/v1/search/words"

type Result struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		/*Slug     string   `json:"slug"`
		IsCommon bool     `json:"is_common"`
		Tags     []string `json:"tags"`
		Jlpt     []string `json:"jlpt"`
		Japanese []struct {
			Word    string `json:"word"`
			Reading string `json:"reading"`
		} `json:"japanese"`*/
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
		/*Attribution struct {
			Jmdict   bool `json:"jmdict"`
			Jmnedict bool `json:"jmnedict"`
			Dbpedia  bool `json:"dbpedia"`
		} `json:"attribution"`*/
	} `json:"data"`
}

func fetchWord(wordToFind string) (Result, error) {
	client := req.C()
	var result Result

	_, err := client.R().SetQueryParam("keyword", wordToFind).SetResult(&result).Get(API_URL)
	return result, err
}

func LookupWord(wordToFind string, listAll bool) {
	result, err := fetchWord(wordToFind)

	if err != nil {
		panic(err)
	}

	if len(result.Data) > 0 {
		fmt.Println("Found results for: " + wordToFind)
	} else {
		fmt.Println("No results found for: " + wordToFind)
	}
}
