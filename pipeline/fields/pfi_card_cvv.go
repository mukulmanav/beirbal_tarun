package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var CardCvvLabel = "PFI_CARD_CVV"

type CardcvvEntry = utils.ProdigyOutput

var sampleTextCardCvv = []string{
	"CVV %v ",
	"CVV1 %v ",
	"CVV2 %v ",
	"CVC2 %v ",
	"CSC %v ",
	"Card Security Value %v ",
	"Card Verification Code %v ",
	"Three-Digit Code %v ",
	"Security Code %v ",
	"Authentication Code %v ",
	"Card Verification %v ",
	"Secure Code %v ",
	"Anti-Fraud Code %v ",
	"Payment Card Security %v ",
	"CIN %v ",
	" %v ",
}

var training_statements_card_cvv = utils.GenTrainingStatement(
	sampleTextCardCvv,
	CardCvvLabel,
)

func GetTextForFieldCardCvv(
	card_cvv string,
) CardcvvEntry {
	return utils.GetRandomEntry(training_statements_card_cvv, card_cvv)
}
