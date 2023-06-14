package main

import (
	"flag"
	"log"

	"github.com/ddomd/sgrep/internal/search"
)

func main() {
	verbose := flag.Bool("v", false, "Display line number and file path")
	regex := flag.Bool("r", false, "Use a regex to form a query")
	workers := flag.Int("w", 10, "Set the number of search workers, default: 10")
	buffer := flag.Int("b", 100, "Set the buffer size, default: 100")
	flag.Parse()

	query, path := parseArgs(flag.Args())

	options := search.NewSearchOptions(*regex, *verbose)
	search := search.NewSearch(*buffer, *workers, options)

	search.Start(path, query)

}

func parseArgs(args []string) (string, string){
	if len(args) < 1 {
		log.Fatal("Not enough arguments: sgrep [-vrwb] <query> [path]")
	}

	if len(args) == 1 {
		return args[0], "."
	}

	return args[0], args[1]
}
