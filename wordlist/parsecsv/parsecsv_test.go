package parsecsv

import "testing"

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
		{Position: "p10 10%", Name: "available", Mean: "使える"},
		{Position: "p20 30%", Name: "currently", Mean: "現在は"},
		{Position: "p30 50%", Name: "annual", Mean: "毎年恒例の"},
	}

	isOk := true
	for i := range actual {
		if actual[i] != expected[i] {
			isOk = false
			break
		}
	}

	if isOk {
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
