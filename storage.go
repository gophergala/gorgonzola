package gorgonzola

import (
	"net/http"
	"time"
)

type Link struct {
	Url     string
	Added   time.Time
	Fetched time.Time
}

type Job struct {
}

type Storage interface {
	AddURL(r *http.Request, url string) error
	GetJobs(r *http.Request) ([]Job, error)
	GetJob(r *http.Request, key string) (*Job, error)
}
