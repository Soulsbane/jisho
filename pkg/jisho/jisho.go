package jisho

import (
	"fmt"
	"os"

	"github.com/imroc/req/v3"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jwalton/gchalk"
)

const API_URL = "https://jisho.org/api/v1/search/words"

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

	_, err := client.R().SetQueryParam("keyword", wordToFind).SetResult(&jishoResult).Get(API_URL)
	return jishoResult, err
}

func LookupWord(wordToFind string, listAll bool) {
	jishoResult, err := fetchWord(wordToFind)

	if err != nil {
		panic(err)
	}

	if len(jishoResult.JishoData) > 0 {
		if listAll {
			handleListAll(jishoResult)
		} else {
			fmt.Println(jishoResult.JishoData[0].Slug)
			fmt.Println(jishoResult.JishoData[0].Japanese[0].Reading)
			fmt.Println(jishoResult.JishoData[0].Senses[0].EnglishDefinitions[0])
		}
	} else {
		fmt.Println("No results found for: " + wordToFind)
	}
}

func handleListAll(jishoResult JishoResult) {
	var SlugColor = gchalk.WithBold().Blue
	var ReadingColor = gchalk.WithBold().Green

	for _, data := range jishoResult.JishoData {
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
