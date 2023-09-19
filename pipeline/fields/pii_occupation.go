package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var OccupationLabel = "PII_OCCUPATION"

type OccupationEntry = utils.ProdigyOutput

var sampleTextOccupation = []string{
	"Occupation %v ",
	"Career %v ",
	"Job %v ",
	"Profession %v ",
	"Employment %v ",
	"Work %v ",
	"Job Title %v ",
	"Trade %v ",
	" %v ",
}

var training_statements_occupation = utils.GenTrainingStatement(
	sampleTextOccupation,
	OccupationLabel,
)

func GetTextForFieldOccupation(
	occupation string,
) OccupationEntry {
	return utils.GetRandomEntry(training_statements_occupation, occupation)
}
