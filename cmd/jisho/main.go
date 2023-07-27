package main

import (
	"fmt"
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

// NOTE: This was taken from a recent go-pretty commit. Once a new release contains this function then it can be removed
func Hyperlink(url, text string) string {
	if url == "" {
		return text
	}
	if text == "" {
		return url
	}
	// source https://gist.github.com/egmontkob/eb114294efbcd5adb1944c9f3cb5feda
	return fmt.Sprintf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", url, text)
}

func handleListAll(jishoResult jisho.JishoResult) {
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

func handleSingleWord(jishoResult jisho.JishoResult) {
	outputTable := table.NewWriter()
	slug := jishoResult.JishoData[0].Slug
	reading := jishoResult.JishoData[0].Japanese[0].Reading

	outputTable.SetOutputMirror(os.Stdout)
	fmt.Println(Hyperlink(JishoSearchURL+slug, slug))
	outputTable.AppendHeader(table.Row{SlugColor(Hyperlink(JishoSearchURL+slug, slug)) + " - " + ReadingColor(reading)})

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
