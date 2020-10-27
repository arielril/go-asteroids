package object

import (
	"strconv"
	"strings"

	"github.com/arielril/go-asteroids/point"
	"github.com/go-gl/gl/v2.1/gl"
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

// getObjectGLPosition returns the real position of an instance inside the GL Window
func getObjectGLPosition(p point.Point) point.Point {
	rawPoint := p.Raw()

	var glMatrix [4][4]float32
	newPoint := make([]float32, 4)

	gl.GetFloatv(gl.MODELVIEW_MATRIX, &glMatrix[0][0])

	for i := range newPoint {
		newPoint[i] = glMatrix[0][i]*rawPoint.X +
			glMatrix[1][i]*rawPoint.Y +
			glMatrix[2][i]*rawPoint.Z +
			glMatrix[3][i]
	}

	return point.New(newPoint[0], newPoint[1], newPoint[2])
}
