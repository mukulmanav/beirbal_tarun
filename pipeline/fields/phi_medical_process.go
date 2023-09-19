package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var MedicalProcessLabel = "PHI_MEDICAL_PROCESS"

type MedicalProcessEntry = utils.ProdigyOutput

var sampleTextMedicalProcess = []string{
	"Medical process %v ",
	"Diagnosis %v ",
	"Treatment %v ",
	"Medical Procedure %v ",
	"Medical Examination %v ",
	"Medication %v ",
	"Medical Records %v ",
	"Medical test %v ",
	" %v ",
}

var training_statements_medical_process = utils.GenTrainingStatement(
	sampleTextMedicalProcess,
	MedicalProcessLabel,
)

func GetTextForFieldMedicalProcess(
	medical_process string,
) UsHealthBenificiaryNumberEntry {
	return utils.GetRandomEntry(training_statements_medical_process, medical_process)
}
