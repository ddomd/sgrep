package search

import (
	"fmt"
	"log"

	"github.com/ddomd/sgrep/utils"
)



func (s *Search) Start(path, query string) {
	s.discoverDirs(path)
	s.gatherResults(query)
	s.displayResults()
}

func (s *Search) discoverDirs(path string) {
	s.WorkersWG.Add(1)

	go func() {
		defer s.WorkersWG.Done()
		utils.CrawlFiles(&s.Queue, path)
		s.Queue.Finalize(s.Queue.Workers)
	}()
}

func (s *Search) gatherResults(query string) {
	for i := 0; i < s.Queue.Workers; i++ {
		s.WorkersWG.Add(1)
		go func() {
			defer s.WorkersWG.Done()
			for {
				job := s.Queue.GetJob()

				if job.Path == "" {
					return
				}

				jobResults, err := utils.ScanFile(job.Path, query, s.Options.Regex)
				if err != nil {
					log.Fatal("Couldn't read file")
				}

				if jobResults != nil {
					for _, result := range jobResults.LinesFound {
						s.Work <- result
					}
				}
			}
		}()
	}

	go func() {
		s.WorkersWG.Wait()
		// Close channel
		close(s.Done)
	}()
}

func (s *Search) displayResults() {
	s.DisplayWG.Add(1)
	go func() {
		for {
			select {
			case r := <-s.Work:
				if s.Options.Lines {
					fmt.Printf("%s:%d %s\n", r.Path, r.LineHeight, r.Line)
				} else {
					fmt.Println(r.Line)
				}
			case <-s.Done:
				// Make sure channel is empty before aborting display goroutine
				if len(s.Work) == 0 {
					s.DisplayWG.Done()
					return
				}
			}
		}
	}()
	s.DisplayWG.Wait()
}