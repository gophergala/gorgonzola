package gorgonzola

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Link struct {
	Url     string
	Added   time.Time
	Fetched time.Time
}

type Job struct {
	Hash                  string
	CompanyName           string
	CompanyURL            string
	CompanyRemoteFriendly bool
	CompanyMarket         string
	CompanySize           string
	Position              string
	Title                 string
	Description           string
	Url                   string
	Type                  string
	Posted                string
	Location              string
	Skills                []string
	SalaryRangeFrom       int
	SalaryRangeTos        int
	SalaryRangeCurrency   string
	EquityFrom            float32
	EquityTo              float32
	Perks                 []string
	Apply                 string
}

type Storage interface {
	AddURL(r *http.Request, url string) error
	GetJobs(r *http.Request) ([]Job, error)
	GetJob(r *http.Request, key string) (*Job, error)
}

func (j *Job) getHash() string {
	h := md5.New()
	io.WriteString(h, j.CompanyName)
	io.WriteString(h, j.Position)
	io.WriteString(h, j.Title)
	io.WriteString(h, j.Url)
	io.WriteString(h, j.Posted)
	return fmt.Sprintf("%x", h.Sum(nil))
}
