package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var LanguageLabel = "PII_LANGUAGE"

type LanguageEntry = utils.ProdigyOutput

var sampleTextLanguage = []string{
	"Language %v ",
	"Tongue %v ",
	"Speech %v ",
	"Communication %v ",
	"Mother Tongue %v ",
	" %v ",
}

var training_statements_language = utils.GenTrainingStatement(
	sampleTextLanguage,
	LanguageLabel,
)

func GetTextForFieldLanguage(
	language string,
) UsHealthBenificiaryNumberEntry {
	return utils.GetRandomEntry(training_statements_language, language)
}
