package gorgonzola

import (
	"net/http"
)

type Job struct {
}

type Storage interface {
	SaveJsonJobs(r *http.Request, jj *JsonJobs) error
	AddURL(r *http.Request, url string) (string, error)
	GetJobs(r *http.Request) ([]Job, error)
	GetJob(r *http.Request, key string) (*Job, error)
}
