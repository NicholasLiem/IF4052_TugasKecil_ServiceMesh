package main

import (
	"net/http"
	"time"
)

func CreateHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}