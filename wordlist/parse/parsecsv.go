package parse

import (
	"errors"
	"fmt"
	"strings"
)

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
			Position: strings.TrimSpace(tmp[0]),
			Name:     strings.TrimSpace(tmp[1]),
			Mean:     strings.TrimSpace(tmp[2]),
		}
	}
	return res, nil
}
