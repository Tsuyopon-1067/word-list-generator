package html

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"wordlist/parsecsv"
)

func TestParseCsvOk(t *testing.T) {
	data := []parsecsv.Word{
		{Position: "p10 10%", Name: "available", Mean: "使える"},
		{Position: "p20 30%", Name: "currently", Mean: "現在は"},
		{Position: "p30 50%", Name: "annual", Mean: "毎年恒例の"},
	}

	expected := `<table>
        <tr>
            <th>position</th>
            <th>word</th>
            <th>meaning</th>
        </tr>
        <tr>
            <td>p10 10%</td>
            <td>available</td>
            <td>使える</td>
        </tr>
        <tr>
            <td>p20 30%</td>
            <td>currently</td>
            <td>現在は</td>
        </tr>
        <tr>
            <td>p30 50%</td>
            <td>annual</td>
            <td>毎年恒例の</td>
        </tr>
    </table>`

	actual := generateTableBody(data)

	expected = deleteBlank(expected)
	actual = deleteBlank(actual)
	expected = tagReturn(expected)
	actual = tagReturn(actual)

	assert.Equal(t, expected, actual)
}

func deleteBlank(text string) string {
	text = strings.Replace(text, "\t", "", -1)
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, " ", "", -1)
	return text
}

func tagReturn(text string) string {
	text = strings.Replace(text, ">", ">\n", -1)
	return text
}
