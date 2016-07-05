package main

import (
	"flag"
	"github.com/kalharbi/sieveable-profile/dbwriter"
	"io/ioutil"
	"log"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatalln("A JSON file must be given as an argument")
	}
	parsed := parseInFile(args[0])
	dbwriter.Insert(parsed)
}

func parseInFile(file string) dbwriter.Response {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read input file. %v\n", err)
	}
	return dbwriter.Parse(&content)
}
