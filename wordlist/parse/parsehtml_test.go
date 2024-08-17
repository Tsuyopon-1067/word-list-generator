package parse

import "testing"

func TestParseHtmlOk(t *testing.T) {
	data := `<!DOCTYPE html>
<html lang="ja">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Word List</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
        }

        table {
            width: 100%;
            border-collapse: separate;
            border-spacing: 0;
            border: 1px solid #ddd;
            border-radius: 8px;
            overflow: hidden;
        }

        th,
        td {
            padding: 12px 15px;
            text-align: left;
        }

        th {
            background-color: #f8f9fa;
            font-weight: 600;
            text-transform: uppercase;
            font-size: 14px;
            letter-spacing: 0.03em;
            border-bottom: 2px solid #ddd;
        }

        tr:nth-child(even) {
            background-color: #f8f9fa;
        }

        tr:hover {
            background-color: #e9ecef;
        }

        th:nth-child(1),
        td:nth-child(1) {
            width: 14.28%;
            /* 1/7 of the total width */
        }

        th:nth-child(2),
        td:nth-child(2),
        th:nth-child(3),
        td:nth-child(3) {
            width: 42.86%;
        }
    </style>
</head>

<body>
    <table>
        <tr>
            <th>position</th>
            <th>word</th>
            <th>meaning</th>
        </tr>
        <tr>
            <td>p10 10%</td>
            <td> available</td>
            <td> 使える</td>
        </tr>
        <tr>
            <td> p20 30%</td>
            <td> currently</td>
            <td> 現在は</td>
        </tr>
        <tr>
            <td> p30 50%</td>
            <td> annual</td>
            <td> 毎年恒例の</td>
        </tr>
    </table>
</body>

</html>
`

	actual, err := ParseHtml(data)
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
			break
		}
	}

	if !isOk {
		t.Errorf("actual: %v\nexpected: %v", actual, expected)
	}
}
