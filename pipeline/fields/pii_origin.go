package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var OriginLabel = "PII_ORIGIN"

type OriginEntry = utils.ProdigyOutput

var sampleTextorigin = []string{
	"Origin %v ",
	"Source %v ",
	"Birthplace %v ",
	"Ancestry %v ",
	"Roots %v ",
	"Native %v ",
	" %v ",
}

var training_statements_origin = utils.GenTrainingStatement(
	sampleTextorigin,
	OriginLabel,
)

func GetTextForFieldOrigin(
	origin string,
) UsHealthBenificiaryNumberEntry {
	return utils.GetRandomEntry(training_statements_origin, origin)
}
