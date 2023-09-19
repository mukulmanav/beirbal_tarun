package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var DateIntervalLabel = "PII_DATE_INTERVAL"

type DateIntervalEntry = utils.ProdigyOutput

var sampleTextDateInterval = []string{
	"date interval %v ",
	"Date Range %v ",
	"Time Interval %v ",
	"Time Period %v ",
	"Duration %v ",
	"Time Span %v ",
	"Date Window %v ",
	"Date Boundary %v ",
	"Date Bracket %v ",
	"Date Range Selection %v ",
	"Date Interval Calculation %v ",
	"Date Gap %v ",
	"Date Overlap %v ",
	"Fiscal Year %v ",
	"Quarter %v ",
	"Month %v ",
	"Week %v ",
	"Day %v ",
	"Business Days %v ",
	"Calendar Year %v ",
	"Financial Year %v ",
	"Date Formatting %v ",
	"Date Precision %v ",
	"following dates %v ",
	"minimum period %v ",
	"range of date %v ",
	"said time interval %v ",
	"time range %v ",
	"span of time %v ",
	"following date %v ",
	" %v ",
}

var training_statements_date_interval = utils.GenTrainingStatement(
	sampleTextDateInterval,
	DateIntervalLabel,
)

func GetTextForFieldDateInterval(
	date_interval string,
) DateIntervalEntry {
	return utils.GetRandomEntry(training_statements_date_interval, date_interval)
}
