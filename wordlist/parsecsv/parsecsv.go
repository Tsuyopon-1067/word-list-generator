package parsecsv

import (
	"errors"
	"fmt"
	"strings"
)

type Word struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Mean     string `json:"mean"`
}

func ParseCsv(csv string) ([]Word, error) {
	fmt.Printf("csv: %v\n", csv)
	lines := strings.Split(csv, "\n")
	res := make([]Word, len(lines))
	for i, line := range lines {
		tmp := strings.Split(line, ",")
		if len(tmp) < 3 {
			msg := fmt.Sprintf("too few columns: %s", tmp)
			err := errors.New(msg)
			return nil, err
		}
		res[i] = Word{
			Position: tmp[0],
			Name:     tmp[1],
			Mean:     tmp[2],
		}
	}
	return res, nil
}
