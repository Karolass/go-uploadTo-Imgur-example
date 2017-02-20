package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	imgurAPIURL = "https://api.imgur.com/3/image"
	imgurKey    = "Your Imgur Client ID"
)

func main() {

	// multipart writer
	reqBody := new(bytes.Buffer)
	writer := multipart.NewWriter(reqBody)

	// open file
	file, err := os.Open("test.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	b, _ := ioutil.ReadAll(file)

	// multipart field
	imageField, _ := writer.CreateFormField("image")
	imageField.Write(b)
	writer.Close()

	// http request
	client := new(http.Client)
	req, err := http.NewRequest("POST", imgurAPIURL, reqBody)
	if err != nil {
		log.Fatalln(err)
	}

	// Headers
	req.Header.Add("Authorization", imgurKey)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	u := map[string]interface{}{}
	err = json.Unmarshal(respBody, &u)
	bytes, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bytes))
}
