package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var CoordinatesLabel = "PII_COORDINATES"

type CoordinatesEntry = utils.ProdigyOutput

var sampleTextCoordinates = []string{
	"coordinates %v ",
	"location %v ",
	"Position %v",
	"Geographic Coordinates %v ",
	"GPS Coordinates %v ",
	"Cartesian Coordinates %v ",
	"Map Coordinates %v ",
	"Grid Coordinates %v ",
	"Spatial Coordinates %v ",
	"Waypoint %v ",
	"Reference Point %v ",
	"Navigation Data %v ",
	"Vector Coordinates %v ",
	"Grid Reference %v ",
	"Decimal Degrees %v ",
	"Longitude and Latitude %v ",
	"Longitude %v ",
	"Latitude %v ",
	"Positional Data %v ",
	"Map Grid %v ",
	"Geospatial Coordinates %v ",
	"Survey Coordinates %v ",
	"Mapping Coordinates %v ",
	" %v ",
}

var training_statements_coordinates = utils.GenTrainingStatement(
	sampleTextCoordinates,
	CoordinatesLabel,
)

func GetTextForFieldCoordinates(
	coordinates string,
) CoordinatesEntry {
	return utils.GetRandomEntry(training_statements_coordinates, coordinates)
}
