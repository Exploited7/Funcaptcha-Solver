package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func solveCaptcha(publickey, website, host, blob string) {
	data := map[string]string{
		"host":      host,
		"publickey": publickey,
		"website":   website,
		"blob":      blob,
		"solvekey":  "TM-upBsYRUYH4SfbpQ13j0AuBpf7giICLVBb5yDDPZYwDSjs",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	resp, err := http.Post("http://23.137.104.216:5000/api/funcaptcha", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error making POST request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	fmt.Println(result["token"])
}

func main() {
	solveCaptcha("73BEC076-3E53-30F5-B1EB-84F494D43DBA", "https://signin.ea.com", "ea-api.arkoselabs.com", "undefined")
}
