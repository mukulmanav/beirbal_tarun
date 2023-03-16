package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
	"github.com/jdkato/prose/v2"
)

var EuDriverLicenceLabel = "pii-eu-driver-licence"

type EuDriverLicenceEntry = utils.ProdigyOutput

var sampleTextEuDriverLicence = []string{
	"driverlic %v",
	"driverlics %v",
	"driverlicense %v",
	"driverlicenses %v",
	"driverlicence %v",
	"driverlicences %v",
	"driver lic %v",
	"driver lics %v",
	"driver license %v",
	"driver licenses %v",
	"driver licence %v",
	"driver licences %v",
	"driverslic %v",
	"driverslics %v",
	"driverslicence %v",
	"driverslicences %v",
	"driverslicense %v",
	"driverslicenses %v",
	"drivers lic %v",
	"drivers lics %v",
	"drivers license %v",
	"drivers licenses %v",
	"drivers licence %v",
	"drivers licence %v",
	"driver'lic %v",
	"driver'lics %v",
	"driver'license %v",
	"driver'licenses %v",
	"driver'licence %v",
	"driver'licences %v",
	"driver' lic %v",
	"driver' lics %v",
	"driver' license %v",
	"driver' licenses %v",
	"driver' licence %v",
	"driver' licences %v",
	"driver'slic %v",
	"driver'slics %v",
	"driver'slicense %v",
	"driver'slicenses %v",
	"driver'slicence %v",
	"driver'slicences %v",
	"driver's lic %v",
	"driver's lics %v",
	"driver's license %v",
	"driver's licenses %v",
	"driver's licence %v",
	"driver's licences %v",
	"dl# %v",
	"dls# %v",
	"driverlic# %v",
	"driverlics# %v",
	"driverlicense# %v",
	"driverlicenses# %v",
	"driverlicence# %v",
	"driverlicences# %v",
	"driver lic# %v",
	"driver lics# %v",
	"driver license# %v",
	"driver licenses# %v",
	"driver licences# %v",
	"driverslic# %v",
	"driverslics# %v",
	"driverslicense# %v",
	"driverslicenses# %v",
	"driverslicence# %v",
	"driverslicences# %v",
	"drivers lic# %v",
	"drivers lics# %v",
	"drivers license# %v",
	"drivers licenses# %v",
	"drivers licence# %v",
	"drivers licences# %v",
	"driver'lic# %v",
	"driver'lics# %v",
	"driver'license# %v",
	"driver'licenses# %v",
	"driver'licence# %v",
	"driver'licences# %v",
	"driver' lic# %v",
	"driver' lics# %v",
	"driver' license# %v",
	"driver' licenses# %v",
	"driver' licence# %v",
	"driver' licences# %v",
	"driver'slic# %v",
	"driver'slics# %v",
	"driver'slicense# %v",
	"driver'slicenses# %v",
	"driver'slicence# %v",
	"driver'slicences# %v",
	"driver's lic# %v",
	"driver's lics# %v",
	"driver's license# %v",
	"driver's licenses# %v",
	"driver's licence# %v",
	"driver's licences# %v",
	"driving licence %v",
	"driving license %v",
	"dlno# %v",
	"driv lic %v",
	"driv licen %v",
	"driv license %v",
	"driv licenses %v",
	"driv licence %v",
	"driv licences %v",
	"driver licen %v",
	"drivers licen %v",
	"driver's licen %v",
	"driving lic %v",
	"driving licen %v",
	"driving licenses %v",
	"driving licence %v",
	"driving licences %v",
	"driving permit %v",
	"dl no %v",
	"dlno %v",
	"dl number %v",
}

var training_statements_eu_driver_licence = []utils.ProdigyOutput{
	VoterIdEntry{
		Text: sampleTextEuDriverLicence[0],
		Spans: []prose.LabeledEntity{
			{
				Start: 5,
				End:   utils.GetWholeEntity(sampleTextEuDriverLicence[0], 5),
				Label: CSINLabel,
			},
		},
	},
}
