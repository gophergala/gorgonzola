package gorgonzola

import (
	"encoding/json"
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
	for _, job := range jj.getJobs() {
		job.Hash = job.getHash()
		key := datastore.NewKey(c, "Job", job.Hash, 0, nil)
		_, err := datastore.Put(c, key, job)
		if err != nil {
			return err
		}
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
	var jjraw []byte
	var err error
	if jjraw, err = getJSONJobsDoc(r, link.Url); err != nil {
		return err
	}
	if err := validateDoc(string(jjraw)); err != nil {
		return err
	}
	var jj JsonJobs
	json.Unmarshal(jjraw, &jj)
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
