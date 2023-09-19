package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var StateLabel = "PII_STATE"

type StateEntry = utils.ProdigyOutput

var sampleTextState = []string{
	" state %v ",
	" state:  %v ",
	" %v ",
}

var training_statements_state = []StateEntry{
	{
		Text: sampleTextState[0],
		Spans: []prose.LabeledEntity{
			{
				Start: 7,
				End:   utils.GetWholeEntity(sampleTextState[0], 7),
				Label: CountryLabel,
			},
		},
	},
	{
		Text: sampleTextState[1],
		Spans: []prose.LabeledEntity{
			{
				Start: 9,
				End:   utils.GetWholeEntity(sampleTextState[1], 9),
				Label: CountryLabel,
			},
		},
	},
	{
		Text: sampleTextState[2],
		Spans: []prose.LabeledEntity{
			{
				Start: 1,
				End:   utils.GetWholeEntity(sampleTextState[2], 1),
				Label: CountryLabel,
			},
		},
	},
}

func GetTextForFieldState(state string) StateEntry {
	return utils.GetRandomEntry(training_statements_state, state)
}
