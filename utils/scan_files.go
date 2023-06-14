package utils

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

func NewResult(path, line string, lineHeight int) Result {
	return Result{path, line, lineHeight}
}

//Opens a new scanner on current text file and looks for the query string
func ScanFile(path , tofind string, isRegex bool) (*Results, error){
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	
	fileResults := Results{make([]Result, 0)}
	lineHeight := 1
	

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//Check if the file is a valid text file
		if !utf8.ValidString(scanner.Text()) {
			return nil, nil
		}

		if isRegex {
			ok, err := regexp.MatchString(tofind, scanner.Text())
			if err != nil {
				return nil, errors.New("Malformed regex")
			}
			if ok {
				found := NewResult(path, scanner.Text(), lineHeight)
				fileResults.LinesFound = append(fileResults.LinesFound, found)
			}
		} else {
			if strings.Contains(scanner.Text(), tofind) {
				found := NewResult(path, scanner.Text(), lineHeight)
				fileResults.LinesFound = append(fileResults.LinesFound, found)
			}
		}

		lineHeight++
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	if len(fileResults.LinesFound) <= 0 {
		return nil, nil
	}

	return &fileResults, nil
}


