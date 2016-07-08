package main

import (
	"flag"
	"fmt"
	"github.com/sieveable/sieveable-profile/dbwriter"
	"io/ioutil"
	"log"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatalln("A JSON file must be given as an argument")
	}
	parsed, err := parseInFile(args[0])
	if err != nil {
		log.Fatalf("Parser Error: %v", err)
	}
	if err := dbwriter.Insert(parsed); err != nil {
		log.Fatalf("Failed to insert app features in %s. Reason: %v", args[0], err)
	}
}

func parseInFile(file string) (res dbwriter.Response, err error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return res, fmt.Errorf("Failed to read input file. %v\n", err)
	}
	return dbwriter.Parse(&content)
}
