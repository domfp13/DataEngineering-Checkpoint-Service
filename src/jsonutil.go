package src

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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

// WriteFileJsonObject uses marshalling to write a struct into a JSON object
// The file save under tmp/checkpoint/
func WriteFileJsonObject(objectName string, c *gin.Context) bool {

	var newCheckpoint checkpointObject
	var checkpointObjects []checkpointObject

	if err := c.Bind(&newCheckpoint); err != nil {
		log.Println(err)
	}

	if newCheckpoint.Time == "" {
		return false
	}

	checkpointObjects = append(checkpointObjects, newCheckpoint)

	content, err2 := json.Marshal(checkpointObjects)
	if err2 != nil {
		log.Println("Error JSON Marshalling")
		log.Println(err2.Error())
	}

	err3 := ioutil.WriteFile(localDir+prefix+objectName+extension, content, 0644)
	if err3 != nil {
		log.Println(err3)
	}
	return true
}
