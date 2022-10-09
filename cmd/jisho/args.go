package main

type ProgramArgs struct {
	Word    string `arg:"positional, required"`
	ListAll bool   `arg:"-a,--list-all" default:"false"`
}
