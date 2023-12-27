package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	apiKey := "AIzaSyA9ri_3bC-QukUP4jk-8GhTV7RUXD18OMM"
	geminiEndpoint := "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent"

	requestBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"role": "USER",
				"parts": []map[string]interface{}{
					{
						"text": "用python画Mandelbrot集合",
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", geminiEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	req.Header.Set("x-goog-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")


	proxyURL,_:= url.Parse("http://localhost:7890")
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	// Make the HTTP request
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer response.Body.Close()

	// Read response body
	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	fmt.Println(result)
}
