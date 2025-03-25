/*
 * Copyright (c) 2025 Michael Plunkett (https://github.com/michplunkett)
 * All rights reserved.
 * Used to delete Reddit history of a given user.
 */

package main

import (
	"fmt"
	"log"
	"os"
)

type envVars struct {
	appID        string
	appSecret    string
	userName     string
	userPassword string
}

func arrayHasNoEmptyStrings(envVars []string) bool {
	for _, value := range envVars {
		if value == "" {
			return false
		}
	}

	return true
}

func main() {
	e := envVars{
		os.Getenv("REDDIT_APP_ID"),
		os.Getenv("REDDIT_SECRET"),
		os.Getenv("REDDIT_USER_ID"),
		os.Getenv("REDDIT_USER_PASSWORD"),
	}

	if !arrayHasNoEmptyStrings([]string{e.appID, e.appSecret, e.userName, e.userPassword}) {
		log.Fatal(fmt.Errorf("one of the last.fm environment variables is not present in your system"))
	}
}
