package utils

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/ddomd/sgrep/internal/workqueue"
)

func CrawlFiles(queue *workqueue.WorkQueue, path string) {
	files, err := crawlDirs(path)
	if err != nil {
		log.Printf("Couldn't read directory: %s", err.Error())
	}

	for _, file := range files {
		next := filepath.Join(path, file.Name())

		if file.IsDir() {
			CrawlFiles(queue, next)
		} else {
			queue.AddJob(workqueue.NewJob(next))
		}
	}
}


//crawlDirs is a modified version of os.ReadDir that returns a list of
//all files in a directory without sorting them alphabetically
func crawlDirs(name string) ([]fs.DirEntry, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dirs, err := f.ReadDir(-1)
	return dirs, err
}