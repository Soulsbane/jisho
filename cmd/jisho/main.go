package main

import (
	"github.com/Soulsbane/jisho/pkg/jisho"
	"github.com/alexflint/go-arg"
)

func main() {
	var args ProgramArgs

	arg.MustParse(&args)
	jisho.LookupWord(args.Word, args.ListAll)
}
