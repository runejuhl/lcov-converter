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
	formatter := lcov.Formatter{
		Path: "target/coverage/lcov.info",
	}

	paths, err := formatter.Search()
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	log.Printf("paths: %+v", paths)
	report, err := formatter.Format()
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	log.Printf("reports: %+v\n", report)

	// e, err := env.New()
	// if err != nil {
	//	log.Fatalf("err: %s", err)
	// }

	// var gitHead, _ = env.GetHead()
	// log.Printf("git: %+v\n", gitHead)

	// m := map[string]interface{}{}

	// g := structs.Map(e.Git)
	// for k, v := range g {
	//	m[k] = v
	// }

	// g = structs.Map(e.CI)
	// for k, v := range g {
	//	m[k] = v
	// }

	// j, err := json.Marshal(m)
	// if err != nil {
	//	log.Fatalf("err: %s", err)
	// }

	// j, err := e.MarshalJSON()
	// if err != nil {
	//	log.Fatalf("err: %s", err)
	// }
	// log.Printf("json: %s", j)
	// ioutil.WriteFile("climate.json", j, 0644)

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
	err = cf.Save()
	if err != nil {
		log.Fatalf("runFormatter failed: %s", err)
	}

	ioutil.WriteFile("climate.json", bb.Bytes(), 0644)
}
