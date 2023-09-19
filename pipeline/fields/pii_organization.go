package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var OrganizationLabel = "PII_ORGANIZATION"

type OrganizationEntry = utils.ProdigyOutput

var sampleTextOrganization = []string{
	"Organization %v ",
	"Company %v ",
	"Corporation %v ",
	"Agency %v ",
	"Association %v ",
	"Enterprise %v ",
	"Institution %v ",
	"Charity %v ",
	"Foundation %v ",
	"Public Sector %v ",
	"Private Sector %v ",
	"NGO %v ",
	"Multinational Corporation %v ",
	" %v ",
}

var training_statements_organization = utils.GenTrainingStatement(
	sampleTextOrganization,
	OrganizationLabel,
)

func GetTextForFieldOrganization(
	organization string,
) OrganizationEntry {
	return utils.GetRandomEntry(training_statements_organization, organization)
}
