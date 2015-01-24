package gorgonzola

import (
	"net/http"

	"appengine"
	"appengine/datastore"
)

type Datastore struct{}

func NewDatastore() *Datastore {
	return &Datastore{}
}

func (ds *Datastore) SaveJsonJobs(r *http.Request, jj *JsonJobs) error {
	c := appengine.NewContext(r)
	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "jobs", nil), jj)
	if err != nil {
		return err
	}
	return nil
}

func (ds *Datastore) AddURL(r *http.Request, url string) (string, error) {
	return "", nil
}

func (ds *Datastore) GetJobs(r *http.Request) ([]Job, error) {
	return nil, nil
}

func (ds *Datastore) GetJob(r *http.Request, key string) (*Job, error) {
	return nil, nil
}
