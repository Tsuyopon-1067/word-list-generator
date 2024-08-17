package generator

import (
	"strings"
	"wordlist/parse"
)

func GenerateCsvText(words []parse.Word) string {
	var builder strings.Builder
	for i, line := range words {
		builder.WriteString(line.Position)
		builder.WriteString(",")
		builder.WriteString(line.Name)
		builder.WriteString(",")
		builder.WriteString(line.Meaning)
		if i < len(words)-1 {
			builder.WriteString("\n")
		}
	}
	return builder.String()
}
