package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var GivenNameLabel = "PII_GIVEN_NAME"

type GivenNameEntry = utils.ProdigyOutput

var sampleTextGivenName = []string{
	"Given Name %v ",
	"First Name %v ",
	"Forename %v ",
	"Personal Name %v ",
	"Middle Name %v ",
	"Full Name %v ",
	"Legal Name %v ",
	"Nickname %v ",
	"Formal Name %v ",
	"Baby Name %v ",
	"Common Name %v ",
	" %v ",
}

var training_statements_Given_name = utils.GenTrainingStatement(
	sampleTextUsHealthBenificiaryNumber,
	UsHealthBenificiaryNumbeLabel,
)

func GetTextForFieldGivenName(
	Given_Name string,
) GivenNameEntry {
	return utils.GetRandomEntry(training_statements_us_health_benificiary, Given_Name)
}
