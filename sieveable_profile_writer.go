package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/sieveable/sieveable-profile/dbwriter"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatalln("A JSON file must be given as an argument")
	}
	db := getDbConnection()
	defer db.Close()
	parsed, err := parseInFile(args[0])
	if err != nil {
		log.Fatalf("Parser Error: %v", err)
	}
	if err := dbwriter.Insert(db, parsed); err != nil {
		log.Fatalf("Failed to insert app features in %s. Reason: %v", args[0], err)
	}
}
func getDbConnection() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("USER")+":"+os.Getenv("PW")+"@/"+os.Getenv("DB"))
	if err != nil {
		log.Fatalf("Failed to get a handle for the database %s. %v\n",
			os.Getenv("DB"), err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to connect to DB. Make sure that the required "+
			"environment variables are set. %v\n", err.Error())
	}
	return db
}
func parseInFile(file string) (res dbwriter.Response, err error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return res, fmt.Errorf("Failed to read input file. %v\n", err)
	}
	return dbwriter.Parse(&content)
}
