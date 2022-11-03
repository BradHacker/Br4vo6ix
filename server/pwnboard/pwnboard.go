package pwnboard

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type Data struct {
	IPs  string `json:"ip"`   // Target IP address as a string
	Type string `json:"type"` // Describes what implant pwnboard is being updated from
}

// Sends a post request with information about a target to pwnboard.
func SendUpdate(ip string, info string) {

	// Grab the the URL/IP of the pwnboard instance that SendUpdate
	// 	is sending the data to.
	PWN_URL, exists := os.LookupEnv("PWN_URL")
	if !exists {
		logrus.Warn("PWN_URL is not set (export PWN_URL=<PWNBOARD URL>)")
		return
	}

	//use the Data struct to organize the data that will be sent to pwnboard
	data := Data{
		IPs:  ip,
		Type: info,
	}

	// Turn data struct into json
	mData, err := json.Marshal(data)
	if err != nil {
		logrus.Warnf("failed to marshal pwnboard data: %v", err)
		return
	}

	// Send json data to pwnboard
	req, err := http.Post(PWN_URL, "application/json", bytes.NewBuffer(mData))
	if err != nil {
		logrus.Warnf("failed to send a post request to pwnboard: %v", err)
		return
	}

	// If anything is returned from pwnboard (usually nothing), print it to the terminal.
	var decoded map[string]interface{}
	json.NewDecoder(req.Body).Decode(&decoded)
	// fmt.Println(decoded)
}
