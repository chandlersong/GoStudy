package cases

import "testing"

/**
how to use an interface as property
*/
func TestCreateObject(t *testing.T) {
	var job Job
	job = &SimpleJob{} //job is pointer, because it's a interface variable, which is a pointer

	service := Service{job: job}

	service.job.execute(t)
}

type Service struct {
	job Job
}

type Job interface {
	execute(t *testing.T)
}

type SimpleJob struct{}

func (sj SimpleJob) execute(t *testing.T) {
	t.Log("simple job was execute")
}
