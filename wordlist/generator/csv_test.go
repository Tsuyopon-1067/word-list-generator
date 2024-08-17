package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"wordlist/parse"
)

func TestParseHtmlOk(t *testing.T) {
	data := []parse.Word{
		{Position: "p10 10%", Name: "available", Meaning: "使える"},
		{Position: "p20 30%", Name: "currently", Meaning: "現在は"},
		{Position: "p30 50%", Name: "annual", Meaning: "毎年恒例の"},
	}

	expected := `p10 10%, available, 使える
	p20 30%, currently, 現在は
	p30 50%, annual, 毎年恒例の`

	actual := GenerateCsvText(data)

	expected = deleteBlank(expected)
	actual = deleteBlank(actual)
	expected = tagReturn(expected)
	actual = tagReturn(actual)

	assert.Equal(t, expected, actual)
}
