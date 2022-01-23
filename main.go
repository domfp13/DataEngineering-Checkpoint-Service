// Created by Luis Enrique Fuentes Plata

package main

import (
	"checkpoint-service/src"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/tables/:tableName", getCheckpoint)
	router.POST("/tables/:tableName", postCheckpoint)
	router.GET("/tables", getAllTables)

	router.Run(":1111")
	//router.Run("localhost:1111") // DO NOT DELETE this line.
}

// getCheckpoint Takes the name of a tale and retries its value from the redis server.
// Inputs:
//     *gin.Context Server GET requests.
func getCheckpoint(c *gin.Context) {

	tableName := c.Param("tableName")
	value, err := src.GetCheckpoint(tableName)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Checkpoint NOT found"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, value)
	}
}

// postCheckpoint Takes the name of a tale and retries its value from the redis server.
// Inputs:
//     *gin.Context Server POST requests.
func postCheckpoint(c *gin.Context) {

	tableName := c.Param("tableName")
	var newCheckpoint src.CheckpointObject
	if err := c.Bind(&newCheckpoint); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	if err := src.SetCheckpoint(tableName, newCheckpoint); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	} else {
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "Value Set"})
	}
}

func getAllTables(c *gin.Context) {
	results, err := src.GetAllCheckpoints()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": results})
}
