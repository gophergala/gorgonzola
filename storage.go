package gorgonzola

import (
	"net/http"
)

type Storage interface {
	SaveJsonJobs(r *http.Request, jj *JsonJobs) error
	AddURL(url string) (string, error)
	GetJobs() ([]Job, error)
}
