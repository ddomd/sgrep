package search

import (
	"sync"

	"github.com/ddomd/sgrep/internal/workqueue"
	"github.com/ddomd/sgrep/utils"
)

type SearchOptions struct {
	Regex bool
	Lines bool
}

type Search struct {
	WorkersWG sync.WaitGroup
	DisplayWG sync.WaitGroup
	Work			chan utils.Result
	Done			chan struct{}
	Queue     workqueue.WorkQueue
	Options   SearchOptions
}

func NewSearch(bufsize, workers int, options SearchOptions) Search{
	return Search{
		Work: make(chan utils.Result, bufsize),
		Done: make(chan struct{}),
		Queue: workqueue.NewWorkQueue(bufsize, workers),
		Options: options,
	}
}

func NewSearchOptions(regex, lines bool) SearchOptions{
	return SearchOptions{regex, lines}
}
