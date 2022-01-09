package main

import (
	"checkpoint-service/src"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	prefix    = "checkpoints/"
	extension = ".json"
)

func main() {
	router := gin.Default()

	router.GET("/tables/:tableName", getTableInfoByName)
	router.POST("/tables/:tableName", postTableInfoByTime)

	// use this for local development instead
	router.Run("localhost:8080")

	// Use this for container run
	//router.Run(":1111")
}

// getTableInfoByName locates the S3 object in the checkpoint prefix whose name matches
// the parameter sent by the client, then returns the checkpointObject in a JSON format.
func getTableInfoByName(c *gin.Context) {

	objectName := c.Param("tableName")
	var awsS3Key string = prefix + objectName + extension
	if src.GetS3BucketFile(awsS3Key) {
		var checkpointObjects = src.ReadFileJsonObject(objectName)
		c.IndentedJSON(http.StatusOK, checkpointObjects)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Checkpoint not found"})

}

// postTableInfoByTime adds/modifies S3 object from JSON received in the request body.
func postTableInfoByTime(c *gin.Context) {

	objectName := c.Param("tableName")
	if src.WriteFileJsonObject(objectName, c) {
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "File Uploaded"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Params not pass or Time empty"})
	}
}
