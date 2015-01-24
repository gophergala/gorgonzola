package gorgonzola

import (
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
)

type Datastore struct{}

func NewDatastore() *Datastore {
	return &Datastore{}
}

func (ds *Datastore) saveJsonJobs(c appengine.Context, jj *JsonJobs) error {
	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Jobs", nil), jj)
	if err != nil {
		return err
	}
	return nil
}

func (ds *Datastore) AddURL(r *http.Request, url string) error {
	c := appengine.NewContext(r)
	link := &Link{
		Url:   url,
		Added: time.Now(),
	}
	key := datastore.NewKey(c, "Link", url, 0, nil)
	if _, err := datastore.Put(c, key, link); err != nil {
		return err
	}
	var jj JsonJobs
	if err := getJSONJobs(c, link.Url, &jj); err != nil {
		return err
	}
	if err := ds.saveJsonJobs(c, &jj); err != nil {
		return err
	}
	link.Fetched = time.Now()
	if _, err := datastore.Put(c, key, link); err != nil {
		return err
	}
	return nil
}

func (ds *Datastore) GetJobs(r *http.Request) ([]Job, error) {
	return nil, nil
}

func (ds *Datastore) GetJob(r *http.Request, key string) (*Job, error) {
	return nil, nil
}
