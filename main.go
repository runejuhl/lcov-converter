package main

import (
	"log"

	"github.com/codeclimate/test-reporter/env"
	"github.com/codeclimate/test-reporter/formatters/lcov"
)

// import "bytes"
// import "encoding/json"

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

	var gitHead, _ = env.GetHead()

	log.Printf("git: %+v\n", gitHead)

}
