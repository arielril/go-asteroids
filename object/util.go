package object

import (
	"strconv"
	"strings"
)

// CreateObjectDataFromFile gets the file values and transform to an object.Data type
func CreateObjectDataFromFile(file string) Data {
	lines := strings.Split(file, "\n")

	var data Data

	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}

		parsedLine := strings.Split(lines[i], " ")

		var valuesList []int

		for _, v := range parsedLine {
			value, _ := strconv.Atoi(v)
			valuesList = append(valuesList, value)
		}

		data = append(data, valuesList)
	}

	return data
}
