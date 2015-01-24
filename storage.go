package main

type JobStorage interface {
	SaveJsonJobs(jj JsonJobs) error
}
