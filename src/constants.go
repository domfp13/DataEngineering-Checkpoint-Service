// Created by Luis Enrique Fuentes Plata

package src

import "os"

const (
	prefix    = "checkpoints/"
	localDir  = "tmp/"
	extension = ".json"
	awsRegion = "us-east-1"
)

var awsS3Bucket = os.Getenv("BUCKET_NAME")
