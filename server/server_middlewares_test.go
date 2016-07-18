package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSingleHostMiddlewareFailure(t *testing.T) {
	middlewares := NewSingleHost(
		http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}),
		"example.com:3000")
	ts := httptest.NewServer(middlewares)
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	defer resp.Body.Close()
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 403 {
		t.Error("Expected HTTP response code of 403 but got %d", resp.StatusCode)
	}
}

func TestSingleHostMiddlewareSuccess(t *testing.T) {
	url := "127.0.0.1:4000"
	ln, err := net.Listen("tcp", url)
	if err != nil {
		t.Errorf("Failed to create server. %v", err)
	}
	middlewares := NewSingleHost(
		http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}),
		url)
	s := &httptest.Server{
		Listener: ln,
		Config:   &http.Server{Handler: middlewares},
	}
	s.Start()
	defer s.Close()
	resp, err := http.Get(s.URL)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected HTTP response code of 200 but got %d", resp.StatusCode)
	}
}

func TestLoggerMiddleware(t *testing.T) {
	loggerHandler := Logger(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	ts := httptest.NewServer(loggerHandler)
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	defer resp.Body.Close()
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to read response body. %v", err)
	}
	if string(body) != "" {
		t.Errorf("Expected response body to be empty but got %s", string(body))
	}
}
