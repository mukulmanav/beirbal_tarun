package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var BloodTypeLabel = "PHI_BLOOD_TYPE"

type BloodTypeEntry = utils.ProdigyOutput

var sampleTextBloodType = []string{
	"Blood type %v ",
	"Blood Group %v ",
	"Blood Donation %v ",
	"Blood Bank %v ",
	"Blood Type Card %v ",
	"Blood System %v ",
}

var training_statements_blood_type = utils.GenTrainingStatement(
	sampleTextBloodType,
	BloodTypeLabel,
)

func GetTextForFieldBloodType(
	blood_type string,
) BloodTypeEntry {
	return utils.GetRandomEntry(training_statements_blood_type, blood_type)
}
