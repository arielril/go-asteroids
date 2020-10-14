package util

import (
	"fmt"
	"io/ioutil"

	"github.com/arielril/go-asteroids/object"
)

// GetObjectDataFromFile executes the parsing of the template file to an understandable type
func GetObjectDataFromFile(filename string) object.Data {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(fmt.Errorf("failed to parse the file %v. %v", filename, err))
	}

	return object.CreateObjectDataFromFile(string(data))
}
