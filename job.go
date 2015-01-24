package main

import (
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

const (
	schemaURL = "https://raw.githubusercontent.com/lukasz-madon/json-job/master/schema.json"
)

type SalaryRange struct {
	From     int    `json:"from"`
	To       int    `json:"to"`
	Currency string `json:"currency"`
}

type Equity struct {
	From float32 `json:"from"`
	To   float32 `json:"to"`
}

type Job struct {
	Position    string      `json:"position"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Url         string      `json:"url"`
	Type        string      `json:"type"`
	Posted      string      `json:"posted"`
	Location    string      `json:"location`
	Skills      []string    `json:"skills"`
	SalaryRange SalaryRange `json:"salaryRange"`
	Equity      Equity      `json:"equity"`
	Perks       []string    `json:"perks"`
	Apply       string      `json:"apply"`
}

type JsonJobs struct {
	Company        string `json:"company"`
	Url            string `json:"url"`
	RemoteFriendly bool   `json:"remoteFriendly"`
	Market         string `json:"market"`
	Size           string `json:"size"`
	Jobs           []Job  `json:"jobs"`
}

func validateDoc(document string) error {
	schemaLoader := gojsonschema.NewReferenceLoader(schemaURL)
	documentLoader := gojsonschema.NewStringLoader(document)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err
	}
	if result.Valid() {
		return nil
	}
	errList := ""
	for _, desc := range result.Errors() {
		errList += desc.String() + "; "
	}
	return errors.New(errList)
}
