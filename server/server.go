package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := getDbConnection()
	if err != nil {
		log.Fatalf("DB connection error. %v", err)
	}
	router := NewRouter(db)
	middlewares := NewSingleHost(Logger(router), getAllowedHost())
	log.Fatal(http.ListenAndServe(":"+getServerPort(), middlewares))
}

func getServerPort() string {
	var port string = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}

func getAllowedHost() string {
	var allowedHost string = "localhost"
	if h := os.Getenv("allowedHost"); h != "" {
		allowedHost = h
	}
	return allowedHost + ":" + getServerPort()
}
func getDbConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("USER")+":"+os.Getenv("PW")+"@/"+os.Getenv("DB"))
	if err != nil {
		return nil, fmt.Errorf("Failed to get a handle for the database %s. %v\n",
			os.Getenv("DB"), err.Error())
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to connect to DB. Make sure that the required "+
			"environment variables are set. %v\n", err.Error())
	}
	return db, nil
}
