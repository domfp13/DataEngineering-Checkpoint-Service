package src

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// checkpointObject represents data about a checkpoint JSON object.
type checkpointObject struct {
	Time string `json:"Time"`
}

// ReadFileJsonObject uses unmarshalling to open an existent JSON object
// and returns its content.
func ReadFileJsonObject(objectName string) []checkpointObject {
	content, err := ioutil.ReadFile(localDir + prefix + objectName + extension)
	if err != nil {
		log.Println(err.Error())
	}

	var checkpointObjects []checkpointObject
	err2 := json.Unmarshal(content, &checkpointObjects)
	if err2 != nil {
		log.Println("Error JSON Unmarshalling")
		log.Println(err2.Error())
	}

	return checkpointObjects
}
