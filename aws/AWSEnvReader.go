package aws

import (
	"fmt"
	"log"
	"os"
)

func ReadAWSEnv() (awsRegion string) {
	accessKey, aFound := os.LookupEnv("AWS_ACCESS_KEY_ID")     // we don't return this because the AWS sdk auto detects this
	secretKey, sFound := os.LookupEnv("AWS_SECRET_ACCESS_KEY") // we don't return this because the AWS sdk auto detects this
	rKey, rFound := os.LookupEnv("AWS_REGION")

	if !aFound {
		log.Fatal(`The "AWS_ACCESS_KEY_ID" env variable was not set.  Please set it.`)
	}
	if !sFound {
		log.Fatal(`The "AWS_SECRET_ACCESS_KEY" env variable was not set.  Please set it.`)
	}
	if !rFound {
		log.Fatal(`The "AWS_REGION" env variable was not set.  Please set it.`)
	}

	fmt.Println("AWS_ACCESS_KEY_ID: " + accessKey)
	fmt.Println("AWS_SECRET_ACCESS_KEY: " + secretKey)
	fmt.Println("AWS_REGION: " + rKey)

	return rKey
}
