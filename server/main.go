package main

import (
	"os"
	"sync"

	"github.com/BradHacker/Br4vo6ix/ent"
	"github.com/sirupsen/logrus"
)

func main() {
	// Print the pretty ASCII art
	PrintCoverArt()

	// Check that all necessary env vars exist
	if err := CheckEnvVars(); err != nil {
		logrus.Errorf("missing environment variable(s): %v", err)
		os.Exit(1)
	}

	var client ent.Client

	var wg sync.WaitGroup
	wg.Add(2)

	go RunAPI(&client, &wg)
	go RunC2(&client, &wg)
	wg.Wait()
}
