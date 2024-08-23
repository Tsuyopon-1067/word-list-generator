package parse

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"wordlist/scraping"
)

type goroutineResult struct {
	index int
	word  Word
	err   error
}

func ParseCsv(csv string) ([]Word, error) {
	csv = deleteBlankLIne(csv)
	lines := strings.Split(csv, "\n")

	size := len(lines)
	ch := make(chan goroutineResult, size)
	var wg sync.WaitGroup
	wg.Add(size)

	for i, line := range lines {
		go func(index int, line string, ch chan goroutineResult, wg *sync.WaitGroup) {
			defer wg.Done()
			word, err := createWord(line)
			ch <- goroutineResult{index, word, err}
		}(i, line, ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	res := make([]Word, size)
	for v := range ch {
		if v.err != nil {
			return nil, v.err
		}
		res[v.index] = v.word
	}
	return res, nil
}

func deleteBlankLIne(str string) string {
	split := strings.Split(str, "\n")
	var newSplit []string
	for _, s := range split {
		if s != "" {
			newSplit = append(newSplit, s)
		}
	}
	return strings.Join(newSplit, "\n")
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
		if meaning == "" {
			meaning = "none"
		}
	}

	return Word{
		Position: position,
		Name:     name,
		Meaning:  meaning,
	}, nil
}
