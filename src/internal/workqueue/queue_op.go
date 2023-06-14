package workqueue

func NewWorkQueue(bufsize int, workers int) WorkQueue {
	return WorkQueue{make(chan Job, bufsize), workers}
}

func NewJob(path string) Job {
	return Job{path}
}

func (queue *WorkQueue) GetJob() Job{
	job := <- queue.Jobs
	return job
}

func (queue *WorkQueue) AddJob(job Job) {
	queue.Jobs <- job
}

func (queue *WorkQueue) Finalize(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		queue.AddJob(Job{""})
	}
}