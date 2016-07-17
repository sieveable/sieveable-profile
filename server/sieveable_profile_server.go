package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

func main() {
	var allowedHost string = "localhost"
	var port string = "3000"
	if h := os.Getenv("allowedHost"); h != "" {
		allowedHost = h
	}
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	db := getDbConnection()
	router := NewRouter(db)
	middlewares := NewSingleHost(Logger(router), allowedHost+":"+port)
	log.Fatal(http.ListenAndServe(":"+port, middlewares))
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
