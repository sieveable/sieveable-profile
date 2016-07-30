package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRoutes(t *testing.T) {
	db, err := getDbConnection()
	if err != nil {
		t.Fatalf("Failed to get a handle for the database %s. %v\n",
			os.Getenv("DB"), err.Error())
	}
	router := NewRouter(db)
	ts := httptest.NewServer(router)
	defer ts.Close()
	var routes []string = []string{"/apps/features/:feature_name",
		"/features/apps/:packageName", "/features/categories/:cat_name"}
	for _, route := range routes {
		res, err := http.Get(ts.URL + route)
		if err != nil {
			t.Errorf("Route %s failed", route)
		}
		if res.StatusCode != 200 {
			t.Errorf("Route %s returned %d HTTP status code", route, res.StatusCode)
		}
	}
}
