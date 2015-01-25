package gorgonzola

import (
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

const (
	schemaURL = "file://./schema.json"
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

type JsonJob struct {
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
	Company        string    `json:"company"`
	Url            string    `json:"url"`
	RemoteFriendly bool      `json:"remoteFriendly"`
	Market         string    `json:"market"`
	Size           string    `json:"size"`
	Jobs           []JsonJob `json:"jobs"`
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

func (jj *JsonJobs) getJobs() []*Job {
	var result []*Job
	for _, job := range jj.Jobs {
		result = append(result, &Job{
			CompanyName:           jj.Company,
			CompanyURL:            jj.Url,
			CompanyRemoteFriendly: jj.RemoteFriendly,
			CompanyMarket:         jj.Market,
			CompanySize:           jj.Size,
			Position:              job.Position,
			Title:                 job.Title,
			Description:           job.Description,
			Url:                   job.Url,
			Type:                  job.Type,
			Posted:                job.Posted,
			Location:              job.Location,
			Skills:                job.Skills,
			SalaryRangeFrom:       job.SalaryRange.From,
			SalaryRangeTos:        job.SalaryRange.To,
			SalaryRangeCurrency:   job.SalaryRange.Currency,
			EquityFrom:            job.Equity.From,
			EquityTo:              job.Equity.To,
			Perks:                 job.Perks,
			Apply:                 job.Apply,
		})
	}
	return result
}
