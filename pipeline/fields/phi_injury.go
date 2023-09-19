package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var InjuryLabel = "PHI_INJURY"

type InjuryEntry = utils.ProdigyOutput

var sampleTextInjury = []string{
	"Injury %v ",
	"Trauma %v ",
	"Wound %v ",
	"Laceration %v ",
	"Abrasion %v ",
	"Contusion %v ",
	"Fracture %v ",
	"Sprain %v ",
	"Strain %v ",
	"Dislocation %v ",
	"Concussion %v ",
	"Burn %v ",
	"Frostbite %v ",
	"Heatstroke %v ",
	"Avulsion %v ",
	"Incision %v ",
	"Puncture %v ",
	"Infection %v ",
	"Treatment %v ",
	"Recovery %v ",
	"First Aid %v ",
	"Emergency Response %v ",
	" %v ",
}

var training_statements_injury = utils.GenTrainingStatement(
	sampleTextInjury,
	InjuryLabel,
)

func GetTextForFieldInjury(
	injury string,
) InjuryEntry {
	return utils.GetRandomEntry(training_statements_us_health_benificiary, injury)
}
