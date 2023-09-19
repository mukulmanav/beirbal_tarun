package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var FamilyNameLabel = "PII_FAMILY_NAME"

type FamilyNameEntry = utils.ProdigyOutput

var sampleTextFamilyName = []string{
	"Family Name %v ",
	"Surname %v ",
	"Last Name %v ",
	"Patronymic %v ",
	"Matronymic %v ",
	"Middle Name %v ",
	"Given Name %v ",
	"Surname Prefix %v ",
	"Maiden Name %v ",
	"Hyphenated Name %v ",
	"Née %v ",
	"Alias %v ",
	"Genealogy %v ",
	"Family Tree %v ",
	"Ancestry %v ",
	"Pedigree %v ",
	"Clan %v ",
	"Heraldry %v ",
	"Cultural Variations %v ",
	"Surname Meaning %v ",
	"married name %v ",
	"sobriquet %v ",
	"epithet %v ",
	"cognomen %v ",
	" %v ",
}

var training_statements_family_name = utils.GenTrainingStatement(
	sampleTextFamilyName,
	FamilyNameLabel,
)

func GetTextForFieldFamilyName(
	family_name string,
) FamilyNameEntry {
	return utils.GetRandomEntry(training_statements_family_name, family_name)
}
