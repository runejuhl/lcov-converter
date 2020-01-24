package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"

	"github.com/codeclimate/test-reporter/formatters"
	"github.com/codeclimate/test-reporter/formatters/lcov"
)

// CoverageFormatter is stolen
type CoverageFormatter struct {
	CoveragePath string
	In           formatters.Formatter
	InputType    string
	Output       string
	Prefix       string
	AddPrefix    string
	writer       io.Writer
}

// Save is stolen
func (f CoverageFormatter) Save() error {
	rep, err := f.In.Format()
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	if len(rep.SourceFiles) == 0 {
		log.Fatalf("could not find coverage info for source files: %s", err)
	}

	if f.writer == nil {
		log.Fatalf("missing writer")
	}

	err = rep.Save(f.writer)
	if err != nil {
		log.Fatalf("could not save file: %s", err)
	}
	return nil
}

func main() {
	bb := &bytes.Buffer{}

	cf := CoverageFormatter{
		Prefix:    ".",
		InputType: "lcov",
		writer:    bb,
		In: &lcov.Formatter{
			Path: "target/coverage/lcov.info",
		},
		Output: "-",
	}

	log.Printf("would write now: %+v", cf)
	err := cf.Save()
	if err != nil {
		log.Fatalf("runFormatter failed: %s", err)
	}

	ioutil.WriteFile("climate.json", bb.Bytes(), 0644)
}
