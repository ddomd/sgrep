package workqueue

type Job struct {
	Path string
}

type WorkQueue struct {
	Jobs      chan Job
	Workers   int
}
