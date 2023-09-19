package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var CityLabel = "PII_CITY"

type CityEntry = utils.ProdigyOutput

var sampleTextCity = []string{
	"City %v ",
	"Urban %v ",
	"Metropolitan %v ",
	"Metropolitan Center %v ",
	"Municipality %v ",
	"Capital City %v ",
	"City Center %v ",
	"City Centre %v ",
	"Suburb %v ",
	"Downtown %v ",
	"Street %v ",
	"City Map %v ",
	"City Zoning %v ",
	"City Zone %v ",
	"Zone %v ",
	"Capital %v ",
	"cosmopolis %v ",
	"City-State %v ",
	"Built-Up Area %v ",
	"Metro Area %v ",
	"Urbia %v ",
	"Urban Area %v ",
	"central city %v ",
	" %v ",
}

var training_statements_city = utils.GenTrainingStatement(
	sampleTextCity,
	CityLabel,
)

func GetTextForFieldCity(
	city string,
) CityEntry {
	return utils.GetRandomEntry(training_statements_city, city)
}
