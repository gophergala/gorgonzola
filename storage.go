package gorgonzola

type Storage interface {
	SaveJsonJobs(jj JsonJobs) error
}
