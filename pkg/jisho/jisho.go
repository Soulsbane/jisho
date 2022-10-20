package jisho

import (
	"fmt"
	"os"

	"github.com/imroc/req/v3"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jwalton/gchalk"
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
			handleListAll(result)
		} else {
			fmt.Println(result.Data[0].Slug)
			fmt.Println(result.Data[0].Japanese[0].Reading)
			fmt.Println(result.Data[0].Senses[0].EnglishDefinitions[0])
		}
	} else {
		fmt.Println("No results found for: " + wordToFind)
	}
}

func handleListAll(result Result) {
	var SlugColor = gchalk.WithBold().Blue
	var ReadingColor = gchalk.WithBold().Green

	for _, data := range result.Data {
		outputTable := table.NewWriter()

		outputTable.SetOutputMirror(os.Stdout)
		outputTable.AppendHeader(table.Row{SlugColor(data.Slug) + " - " + ReadingColor(data.Japanese[0].Reading)})

		for _, sense := range data.Senses {
			for _, definition := range sense.EnglishDefinitions {
				outputTable.AppendRow(table.Row{definition})
			}
		}

		outputTable.SetStyle(table.StyleRounded)
		outputTable.Render()

		fmt.Println()
	}
}
