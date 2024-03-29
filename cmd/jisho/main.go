package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"unicode"

	"github.com/Soulsbane/jisho/pkg/jisho"
	"github.com/alexflint/go-arg"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jwalton/gchalk"
)

var SlugColor = gchalk.WithBold().Blue
var ReadingColor = gchalk.WithBold().Green
var JishoSearchURL = "https://jisho.org/search/"

func isLatin(word string) bool {
	for _, letter := range word {
		// Kanji and Kana are not supported so bail out early
		if !unicode.In(letter, unicode.Latin) {
			return false
		}
	}

	return true
}

func handleListAll(jishoResult jisho.JishoResult) {
	for _, data := range jishoResult.JishoData {
		outputTable := table.NewWriter()
		slug := data.Slug

		outputTable.SetOutputMirror(os.Stdout)
		outputTable.AppendHeader(table.Row{SlugColor(text.Hyperlink(JishoSearchURL+slug, slug)) + " - " + ReadingColor(data.Japanese[0].Reading)})

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

func handleSingleWord(jishoResult jisho.JishoResult) {
	outputTable := table.NewWriter()
	slug := jishoResult.JishoData[0].Slug
	reading := jishoResult.JishoData[0].Japanese[0].Reading

	outputTable.SetOutputMirror(os.Stdout)
	outputTable.AppendHeader(table.Row{SlugColor(text.Hyperlink(JishoSearchURL+slug, slug)) + " - " + ReadingColor(reading)})

	for _, sense := range jishoResult.JishoData[0].Senses {
		for _, definition := range sense.EnglishDefinitions {
			outputTable.AppendRow(table.Row{definition})
		}
	}

	outputTable.SetStyle(table.StyleRounded)
	outputTable.Render()

	fmt.Println()
}

func main() {
	var args ProgramArgs

	arg.MustParse(&args)

	if isLatin(args.Word) {
		jishoResult, err := jisho.LookupWord(args.Word)

		if err != nil {
			fmt.Println(err)
		} else {
			if args.ListAll {
				handleListAll(jishoResult)
			} else {
				handleSingleWord(jishoResult)
			}
		}

	} else {
		fmt.Println("Kanji and Kana not supported!")
	}
}
