package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var MaritalStatusLabel = "PII_MARITAL_STATUS"

type MaritalStatusEntry = utils.ProdigyOutput

var sampleTextMaritalStatus = []string{
	"Marital Status %v ",
	"Matrimony %v ",
	"Matrimonial Status %v ",
	" %v ",
}

var training_statements_marital_status = utils.GenTrainingStatement(
	sampleTextMaritalStatus,
	MaritalStatusLabel,
)

func GetTextForFieldMaritalstatus(
	marital_status string,
) MaritalStatusEntry {
	return utils.GetRandomEntry(training_statements_marital_status, marital_status)
}
