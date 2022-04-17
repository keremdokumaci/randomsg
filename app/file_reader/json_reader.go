package filereader

import (
	"encoding/json"
	"os"

	"github.com/keremdokumaci/randomsg/app/helper"
)

type JsonReader struct {
	content map[string]interface{}
}

func (r JsonReader) Read(filePath string) map[string]interface{} {
	r.content = make(map[string]interface{})

	file, err := os.Open(filePath)
	if err != nil {
		helper.ErrorText("An error occured while opening the file.\n" + err.Error())
	}

	fileContent := make(map[string]interface{})

	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&fileContent); err != nil {
		helper.ErrorText("An error occured while parsing the file\n" + err.Error())
	}
	r.content = fileContent

	return fileContent
}
