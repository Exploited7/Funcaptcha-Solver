package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SolveCaptcha(publickey, website string) {
	data := map[string]string{
		"host":      website,
		"publickey": publickey,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	resp, err := http.Post("http://23.137.104.216:5000/api/funcaptcha", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	fmt.Println(result["token"])
}

func main() {
	SolveCaptcha("73BEC076-3E53-30F5-B1EB-84F494D43DBA", "ea-api.arkoselabs.com")
}