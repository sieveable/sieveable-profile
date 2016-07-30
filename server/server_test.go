package main

import (
	"os"
	"testing"
)

func setValidEnv() {
	os.Setenv("USER", "travis")
	os.Setenv("DB", "test_apps")
	os.Unsetenv("PW")
}

func TestSuccessfulConnection(t *testing.T) {
	setValidEnv()
	db, err := getDbConnection()
	if err != nil || db == nil {
		t.Fail()
	}
}

func TestFailedConnection(t *testing.T) {
	os.Unsetenv("USER")
	os.Unsetenv("DB")
	db, err := getDbConnection()
	if err == nil || db != nil {
		t.Fail()
	}
	setValidEnv()
}

func TestGetAllowedHost(t *testing.T) {
	// testing default allowed host
	if host := getAllowedHost(); host != "localhost:3000" {
		t.Errorf("Unexpected default allowed host. %s", host)
	}
	// testing custom allowed host
	os.Setenv("allowedHost", "example.com")
	os.Setenv("PORT", "4000")
	if host := getAllowedHost(); host != "example.com:4000" {
		t.Errorf("Unexpected custom allowed host value. %s", host)
	}
	os.Unsetenv("allowedHost")
	os.Unsetenv("PORT")
}
