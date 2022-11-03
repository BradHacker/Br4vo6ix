package main

import (
	"fmt"
	"os"
)

// All required env variables
var REQUIRED_ENV_VARS = []string{"BR4VO_KEY", "PORT"}

func CheckEnvVars() error {
	for _, envVar := range REQUIRED_ENV_VARS {
		if _, exists := os.LookupEnv(envVar); !exists {
			return fmt.Errorf("environment variable \"%s\" does not exist. please define it", envVar)
		}
	}
	return nil
}
