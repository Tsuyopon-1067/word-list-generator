package generator

import (
	"strings"
	"wordlist/parse"
)

func GenerateCsvText(words []parse.Word) string {
	var builder strings.Builder
	for _, line := range words {
		builder.WriteString(line.Position)
		builder.WriteString(",")
		builder.WriteString(line.Name)
		builder.WriteString(",")
		builder.WriteString(line.Meaning)
		builder.WriteString("\n")
	}
	return builder.String()
}
