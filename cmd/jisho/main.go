package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/tiagomelo/go-clipboard/clipboard"
	"os"
	"unicode"

	"github.com/Soulsbane/jisho/pkg/jisho"
	"github.com/alexflint/go-arg"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jwalton/gchalk"
)

var WordColor = gchalk.WithBold().Blue
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
		word := data.Japanese[0].Word
		reading := data.Japanese[0].Reading

		outputTable.SetOutputMirror(os.Stdout)
		outputTable.AppendHeader(table.Row{text.Hyperlink(JishoSearchURL+slug, WordColor(word)) + " - " + ReadingColor(reading)})

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
	word := jishoResult.JishoData[0].Japanese[0].Word
	reading := jishoResult.JishoData[0].Japanese[0].Reading

	outputTable.SetOutputMirror(os.Stdout)
	outputTable.AppendHeader(table.Row{text.Hyperlink(JishoSearchURL+slug, WordColor(word)) + " - " + ReadingColor(reading)})

	for _, sense := range jishoResult.JishoData[0].Senses {
		for _, definition := range sense.EnglishDefinitions {
			outputTable.AppendRow(table.Row{definition})
		}
	}

	outputTable.SetStyle(table.StyleRounded)
	outputTable.Render()

	fmt.Println()
}

func handleCopyToClipboard(jishoResult jisho.JishoResult) {
	c := clipboard.New()
	word := jishoResult.JishoData[0].Japanese[0].Word

	if err := c.CopyText(word); err != nil {
		fmt.Println(err)
	}
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
			} else if args.Copy {
				handleCopyToClipboard(jishoResult)
			} else {
				handleSingleWord(jishoResult)
			}
		}

	} else {
		fmt.Println("Kanji and Kana not supported!")
	}
}
