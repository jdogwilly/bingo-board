package main

import (
	"bytes"
	"text/template"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// HTMLTemplate is the test bingoboard template
var HTMLTemplate = `
<head>
	<title>{{ .BoardTitle }}</title>
	<style>
		table, td, th {
			border: 1px solid black;
			text-align: center;
		}
		td {
			padding: 1rem
		}
		#container {
			text-align: center;
			margin: auto;
			max-width: 75%;
		}
	</style>
</head>
<body>
	<div id="container">
		<h1>{{ .BoardTitle }}</h1>
		<table id="card">
			<tr>
				<td>{{ index .Board 0 0 }}</td>
				<td>{{ index .Board 0 1 }}</td>
				<td>{{ index .Board 0 2 }}</td>
				<td>{{ index .Board 0 3 }}</td>
				<td>{{ index .Board 0 4 }}</td>
			</tr>
			<tr>
				<td>{{ index .Board 1 0 }}</td>
				<td>{{ index .Board 1 1 }}</td>
				<td>{{ index .Board 1 2 }}</td>
				<td>{{ index .Board 1 3 }}</td>
				<td>{{ index .Board 1 4 }}</td>
			</tr>
			<tr>
				<td>{{ index .Board 2 0 }}</td>
				<td>{{ index .Board 2 1 }}</td>
				<td>{{ index .Board 2 2 }}</td>
				<td>{{ index .Board 2 3 }}</td>
				<td>{{ index .Board 2 4 }}</td>
			</tr>
			<tr>
				<td>{{ index .Board 3 0 }}</td>
				<td>{{ index .Board 3 1 }}</td>
				<td>{{ index .Board 3 2 }}</td>
				<td>{{ index .Board 3 3 }}</td>
				<td>{{ index .Board 3 4 }}</td>
			</tr>
			<tr>
				<td>{{ index .Board 4 0 }}</td>
				<td>{{ index .Board 4 1 }}</td>
				<td>{{ index .Board 4 2 }}</td>
				<td>{{ index .Board 4 3 }}</td>
				<td>{{ index .Board 4 4 }}</td>
			</tr>
		</table>
	</div>
</body>
`

// BingoBoard type to represent a bingo board
type BingoBoard struct {
	BoardTitle string
	Board      [5][5]string
}

// BuildPage builds the bingo board from a given template board
// and given bingo layout
func BuildPage(htmlTemplate string, bingoBoard BingoBoard) *bytes.Buffer {
	var bodyBuffer bytes.Buffer
	t := template.New("template")
	var templates = template.Must(t.Parse(htmlTemplate))
	templates.Execute(&bodyBuffer, bingoBoard)
	return &bodyBuffer
}

// Handler is the main function being run by Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var testBoard = BingoBoard{
		BoardTitle: "Test Board",
		Board: [5][5]string{
			{
				"Does all my resume have to be CS stuff?",
				"\"But I'm a sophomore.\"",
				"Shameless plug for club",
				"I only know C",
				"Data Structures & Algorithms",
			},
			{
				"Team Experiences",
				"What programming languages/technologies?",
				"Career Fairs",
				"Co-ops",
				"Some idiot refuses to use Slack",
			},
			{
				"\"Sophomore by credit hours\"",
				"Panelist says you don't need social skills",
				"Free Space",
				"Machine Learning",
				"\"But I'm a freshman.\"",
			},
			{
				"Panelist talks way too long",
				"Technical Interviews",
				"Research",
				"What if I don't know how to code?",
				"Personal Projects",
			},
			{
				"Shameless plug for past employer",
				"\"But I'm a junior.\"",
				"Panelist complains CS majors have no social skills",
				"Google",
				"Git",
			},
		},
	}
	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"content-type": "text/html"},
		Body:       BuildPage(HTMLTemplate, testBoard).String(),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
