package modules

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*CsvReader is a method to read problem questions from csv sources, it takes no parameters*/
func CsvReader() {

	data, err := os.Open("./question.txt")
	check(err)
	r := csv.NewReader(data)
	records, err := r.ReadAll()
	check(err)

	questions := parseLines(records)

	correct := 0
	for index, question := range questions {
		fmt.Printf("Question #%d: %s = \n", index+1, question.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == question.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(questions))

}

func parseLines(lines [][]string) []question {
	ret := make([]question, len(lines))
	for i, line := range lines {
		ret[i] = question{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type question struct {
	q string
	a string
}
