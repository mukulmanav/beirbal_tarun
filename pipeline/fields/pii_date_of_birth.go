package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var DateOfBirthLabel = "PII_DATE_OF_BIRTH"

type DateOfBirthEntry = utils.ProdigyOutput

var sampleTextDateOfBirth = []string{
	"Date of Birth %v ",
	"Birthdate %v ",
	"Birth Year %v ",
	"Birth Month %v ",
	"Birth Day %v ",
	"Birth Day %v ",
	"Age %v ",
	"DOB Format %v ",
	"Date of Birth Field %v ",
	"DOB Verification %v ",
	"Age Calculation %v ",
	"Legal Age %v ",
	"Age Group %v ",
	"DOB Privacy %v ",
	"DOB Record %v ",
	"DOB Certificate %v ",
	"DOB Retrieval %v ",
	"DOB Format Conversion %v ",
	"DOB Format Conversion %v ",
	"DOB Accuracy %v ",
	"DOB Authentication %v ",
	"DOB %v ",
	"birthday %v ",
	" %v ",
}

var training_statements_date_of_birth = utils.GenTrainingStatement(
	sampleTextDateOfBirth,
	DateOfBirthLabel,
)

func GetTextForFieldDateOfBirth(
	date_of_birth string,
) DateOfBirthEntry {
	return utils.GetRandomEntry(training_statements_date_of_birth, date_of_birth)
}
