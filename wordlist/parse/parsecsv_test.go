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

func TestParseCsvComplementOk(t *testing.T) {
	data := `p10 10%, available
	p20 30%, currently
	p30 50%, annual`

	actual, err := ParseCsv(data)
	if err != nil {
		t.Errorf("error occured: %s", err)
		return
	}
	expected := []Word{
		{Position: "p10 10%", Name: "available", Meaning: "(すぐに)利用できる、入手できる、得られる、(…に)(すぐ)利用できて、入手できて、手に入る、(…に)手に入って、手があいていて、(…に)(ひまで)会ってもらえて"},
		{Position: "p20 30%", Name: "currently", Meaning: "現在は、今のところ、目下、一般に、広く"},
		{Position: "p30 50%", Name: "annual", Meaning: "一年の、年々の、例年の、年一回の、年刊の、一年生の"},
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
