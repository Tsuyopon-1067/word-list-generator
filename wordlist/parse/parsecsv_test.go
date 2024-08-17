package parse

import (
	"fmt"
	"testing"
)

func TestParseCsvOk(t *testing.T) {
	data := `p10 10%, available, 使える
	p20 30%, currently, 現在は
	p30 50%, annual, 毎年恒例の`

	actual, err := ParseCsv(data)
	if err != nil {
		t.Errorf("error occured: %s", err)
		return
	}
	expected := []Word{
		{Position: "p10 10%", Name: "available", Meaning: "使える"},
		{Position: "p20 30%", Name: "currently", Meaning: "現在は"},
		{Position: "p30 50%", Name: "annual", Meaning: "毎年恒例の"},
	}

	isOk := true
	for i := range actual {
		if actual[i] != expected[i] {
			isOk = false
			fmt.Printf("i=%d, actual: %v\nexpected: %v\n", i, actual[i], expected[i])
			break
		}
	}

	if !isOk {
		t.Errorf("actual: %v\nexpected: %v", actual, expected)
	}
}

func TestParseCsvFailed(t *testing.T) {
	data := `p10 10%, available, 使える
	p20 30%, 
	p30 50%, annual, 毎年恒例の`

	actual, err := ParseCsv(data)
	if err == nil {
		t.Errorf("actual: %v", actual)
	}
}
