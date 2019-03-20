// Copyright 2017 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// [START logging_stdlogging]

// Sample stdlogging writes log.Logger logs to the Stackdriver Logging.
package main

import (
	"fmt"
	"log"
	"time"

	// Imports the Stackdriver Logging client package.
	"cloud.google.com/go/logging"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	cmnLabels := make(map[string]string)
	cmnLabels["key1"] = "value1"
	cmnLabels["key1"] = "value1"

	// Sets your Google Cloud Platform project ID.
	projectID := "teamwork-poc-233820"

	// Creates a client.
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
	fmt.Println("Successfully created stackdriver client")
	
	if err := client.Ping(ctx); err != nil {
		fmt.Println(" ping failed..")
	}else{
	    fmt.Println(" ping succeeded.")	
	}
	// Sets the name of the log to write to.
	logName := "NGDS"

	loggerOpt := logging.CommonLabels(cmnLabels)

	logger := client.Logger(logName,loggerOpt)
	//logger := client.Logger(logName).StandardLogger(logging.Info)
	err =logger.LogSync(ctx, logging.Entry{Payload: "ALERT! Something critical happened!"})
	if err != nil {
		log.Fatalf("Failed to log in sync: %v", err)
	}else{
		fmt.Println("log sync  succeeded.")	
	}

	// Logs "hello world", log entry is visible at
	// Stackdriver Logs.
	//logger.Println("NGDS Logs start")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Exiting after logging")
}

// [END logging_stdlogging]
