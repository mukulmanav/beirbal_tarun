package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var DrugLabel = "PHI_DRUG"

type DrugEntry = utils.ProdigyOutput

var sampleTextDrug = []string{
	"Medication %v ",
	"Pharmaceutical %v ",
	"Prescription %v ",
	"Over-the-Counter %v ",
	"Generic Name %v ",
	"Drug Interaction %v ",
	"Drug Class %v ",
	"Drug Class %v ",
	"Addiction %v ",
	"Withdrawal %v ",
	"Narcotic %v ",
	"Depressant %v ",
	"Depressant %v ",
	"Hallucinogen %v ",
	"Controlled Substance %v ",
	"Drug Testing %v ",
	"Drug Abuse %v ",
	"Drug Rehabilitation %v ",
	"Pharmacy %v ",
	"Antibiotic %v ",
	"Psychotropic %v ",
	"Biologic %v ",
	"Clinical Trial %v ",
	"Pharmacokinetics %v ",
	" %v ",
}

var training_statements_drug = utils.GenTrainingStatement(
	sampleTextDrug,
	DrugLabel,
)

func GetTextForFieldDrug(
	drug string,
) DrugEntry {
	return utils.GetRandomEntry(training_statements_drug, drug)
}
