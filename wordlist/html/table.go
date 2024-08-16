package html

import "wordlist/parsecsv"

func GenerateTableHtml(words []parsecsv.Word) string {
	res := generateHead()
	res += "<body>"
	res += generateTableBody(words)
	res += "</body>"
	res += "</html>"
	return res
}

func generateHead() string {
	return `<!DOCTYPE html>
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
</head>`
}

func generateTableBody(words []parsecsv.Word) string {
	res := `<table>
	<tr>
	<th>position</th>
	<th>word</th>
	<th>meaning</th>
	</tr>`
	for _, word := range words {
		res += "<tr>"
		res += "<td>" + word.Position + "</td>"
		res += "<td>" + word.Name + "</td>"
		res += "<td>" + word.Mean + "</td>"
		res += "</tr>"
	}
	res += "</table>"
	return res
}
