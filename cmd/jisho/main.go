package main

import (
	"errors"
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

func getOutputTable(headerText string, maxWidth int) table.Writer {
	outputTable := table.NewWriter()

	outputTable.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignLeft, WidthMax: maxWidth},
	})

	outputTable.SetOutputMirror(os.Stdout)
	outputTable.AppendHeader(table.Row{headerText})
	outputTable.SetStyle(table.StyleRounded)
	outputTable.Style().Options.SeparateRows = false

	return outputTable
}

func handleListAll(jishoResult jisho.JishoResult, maxWidth int) {
	for _, data := range jishoResult.JishoData {
		slug := data.Slug
		word := data.Japanese[0].Word
		reading := data.Japanese[0].Reading
		headerText := text.Hyperlink(JishoSearchURL+slug, WordColor(word)) + " - " + ReadingColor(reading)
		outputTable := getOutputTable(headerText, maxWidth)

		for _, sense := range data.Senses {
			for _, definition := range sense.EnglishDefinitions {
				outputTable.AppendRow(table.Row{definition})
			}
		}

		outputTable.Render()
		fmt.Println()
	}
}

func handleSingleWord(jishoResult jisho.JishoResult, maxWidth int) {
	slug := jishoResult.JishoData[0].Slug
	word := jishoResult.JishoData[0].Japanese[0].Word
	reading := jishoResult.JishoData[0].Japanese[0].Reading
	headerText := text.Hyperlink(JishoSearchURL+slug, WordColor(word)) + " - " + ReadingColor(reading)
	outputTable := getOutputTable(headerText, maxWidth)

	for _, sense := range jishoResult.JishoData[0].Senses {
		for _, definition := range sense.EnglishDefinitions {
			outputTable.AppendRow(table.Row{definition})
		}
	}

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
			if errors.Is(err, jisho.NoResultsErr) {
				fmt.Println("No results found for: " + args.Word)
			} else if errors.Is(err, jisho.FailedToLookupWordErr) {
				fmt.Println(err)
			} else {
				fmt.Println(err)
			}
		} else {
			if args.ListAll {
				handleListAll(jishoResult, args.MaxWidth)
			} else if args.Copy {
				handleCopyToClipboard(jishoResult)
			} else {
				handleSingleWord(jishoResult, args.MaxWidth)
			}
		}

	} else {
		fmt.Println("Kanji and Kana not supported!")
	}
}
