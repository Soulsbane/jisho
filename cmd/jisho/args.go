package main

type ProgramArgs struct {
	Word    string `arg:"positional, required"`
	ListAll bool   `arg:"-a,--list-all" default:"false"`
}

func (args ProgramArgs) Description() string {
	return "Gets the definition, reading and kanji for the given romaji input"
}
