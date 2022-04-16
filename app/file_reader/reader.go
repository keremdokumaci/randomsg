package filereader

import (
	"github.com/keremdokumaci/random-message-generator/app/helper"
)

type FileReaderType string

const (
	JSON FileReaderType = ".json"
	YAML                = ".yaml"
	XML                 = ".xml"
)

type IReader interface {
	Read(filePath string) map[string]interface{}
}

func NewFileReader(readerType string) IReader {
	var reader IReader
	switch readerType {
	case string(JSON):
		reader = JsonReader{}
	default:
		helper.ErrorText("Couldn't find the related file reader !")
	}

	return reader
}
