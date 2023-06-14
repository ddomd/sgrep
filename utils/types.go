package utils

type Result struct {
	Path       string
	Line       string
	LineHeight int
}

type Results struct {
	LinesFound []Result
}