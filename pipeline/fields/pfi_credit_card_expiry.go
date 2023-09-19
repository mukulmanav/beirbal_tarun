package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var CreditCardexpiryLabel = "PFI_CREDIT_CARD_EXPIRY"

type CreditCardExpiryEntry = utils.ProdigyOutput

var sampleTextCreditCardExpiry = []string{
	"Expiration Date %v ",
	"Card Expiry %v ",
	"Card Expiration Date %v ",
	"Valid Through %v ",
	"Expires On %v ",
	"Expiration Month %v ",
	"Expiration Year %v ",
	"Card Validity %v ",
	"Card Renewal %v ",
	"Renewal Date %v ",
	"CVV Expiry %v ",
	"Card Replacement %v ",
	"Cardholder Agreement %v ",
	"Valid Until %v ",
	"Month/Year Expiry %v ",
	"Expiration Notice %v ",
	"Card Reactivation %v ",
	"Card Inactivity %v ",
	"Card Status %v ",
	"Card Reissue %v ",
	" %v ",
}

var training_statements_credit_card_expiry = utils.GenTrainingStatement(
	sampleTextCreditCardExpiry,
	CreditCardexpiryLabel,
)

func GetTextForFieldCreditCardExpiry(
	credit_card_expiry string,
) CreditCardExpiryEntry {
	return utils.GetRandomEntry(training_statements_credit_card_expiry, credit_card_expiry)
}
