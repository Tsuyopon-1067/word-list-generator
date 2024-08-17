package parse

import (
	"errors"
	"fmt"
	"strings"
	"wordlist/scraping"
)

func ParseCsv(csv string) ([]Word, error) {
	lines := strings.Split(csv, "\n")
	res := make([]Word, len(lines))
	for i, line := range lines {
		word, err := createWord(line)
		if err != nil {
			return nil, err
		}
		res[i] = word
	}
	return res, nil
}

func createWord(line string) (Word, error) {
	splited := strings.Split(line, ",")
	if len(splited) < 2 {
		msg := fmt.Sprintf("too few columns: %s", splited)
		err := errors.New(msg)
		return Word{}, err
	}

	position := strings.TrimSpace(splited[0])
	name := strings.TrimSpace(splited[1])
	meaning := ""
	if len(splited) >= 3 {
		meaning = strings.TrimSpace(splited[2])
	}

	if strings.TrimSpace(name) == "" {
		msg := fmt.Sprintf("empty word %s", splited)
		err := errors.New(msg)
		return Word{}, err
	}

	if len(splited) == 2 || strings.TrimSpace(meaning) == "" {
		meaning = scraping.FetchWordMeaning(name)
		meaning = strings.TrimSpace(meaning)
	}

	return Word{
		Position: position,
		Name:     name,
		Meaning:  meaning,
	}, nil
}
