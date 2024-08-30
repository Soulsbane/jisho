package main

import (
	"fmt"

	"github.com/carlmjohnson/versioninfo"
)

type ProgramArgs struct {
	Word    string `arg:"positional, required"`
	ListAll bool   `arg:"-a,--list-all" default:"false"`
	Copy    bool   `arg:"-c,--copy" default:"false" help:"Copy kanji output to the clipboard"`
	MaxWidth int    `arg:"-w,--width" default:"80" help:"The max width of the output table"`
}

func (args ProgramArgs) Description() string {
	return "Gets the definition, reading and kanji for the given romaji input"
}

func (ProgramArgs) Version() string {
	return fmt.Sprintln("Version: ", versioninfo.Short())
}
