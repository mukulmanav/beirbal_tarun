package fields

import (
	"github.com/bihari123/beirbal/pipeline/utils"
)

var filenameLabel = "PII_FILENAME"

type FilenameEntry = utils.ProdigyOutput

var sampleTextfilename = []string{
	"Filename %v ",
	"File Path %v ",
	"Directory %v ",
	"File Type %v ",
	"File Format %v ",
	"File Attribute %v ",
	"File Backup %v ",
	"File Compression %v ",
	"File Encryption %v ",
	"File Metadata %v ",
	"File Explorer %v ",
	" %v ",
}

var training_statements_filename = utils.GenTrainingStatement(
	sampleTextfilename,
	filenameLabel,
)

func GetTextForFieldFilename(
	filename string,
) FilenameEntry {
	return utils.GetRandomEntry(training_statements_filename, filename)
}
