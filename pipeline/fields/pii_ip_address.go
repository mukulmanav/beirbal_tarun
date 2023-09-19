package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var IPaddressLabel = "PII_IP_ADDRESS"

type IPaddressEntry = utils.ProdigyOutput

var sampleTextIPaddress = []string{
	"IP Address %v ",
	"IPv4 %v ",
	"IPv6 %v ",
	"Subnet %v ",
	"Gateway %v ",
	"DNS %v ",
	"Router %v ",
	"LAN %v ",
	"WAN %v ",
	"Firewall %v ",
	" %v ",
}

var training_statements_ip_address = utils.GenTrainingStatement(
	sampleTextIPaddress,
	IPaddressLabel,
)

func GetTextForFieldIPaddress(
	ip_address string,
) IPaddressEntry {
	return utils.GetRandomEntry(training_statements_ip_address, ip_address)
}
