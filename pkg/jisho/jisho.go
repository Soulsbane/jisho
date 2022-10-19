package jisho

import (
	"fmt"

	"github.com/imroc/req/v3"
)

const API_URL = "https://jisho.org/api/v1/search/words"

type Result struct {
	Data []struct {
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
		if listAll {
			for _, data := range result.Data {
				fmt.Println(data.Slug)
				fmt.Println(data.Japanese[0].Reading)
				for _, sense := range data.Senses {
					for _, definition := range sense.EnglishDefinitions {
						fmt.Println(definition)
					}
				}
				fmt.Println()
			}
		} else {
			fmt.Println(result.Data[0].Slug)
			fmt.Println(result.Data[0].Japanese[0].Reading)
			fmt.Println(result.Data[0].Senses[0].EnglishDefinitions[0])
		}
	} else {
		fmt.Println("No results found for: " + wordToFind)
	}
}
