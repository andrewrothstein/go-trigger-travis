package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func MakeBody(target_username, target_repo string) io.Reader {

	message := map[string]interface{}{
		"request": map[string]string{
			"message": "trigger from here",
			"branch":  "master",
		},
	}

	bytesRep, err := json.Marshal(&message)
	if err != nil {
		log.Fatalln(err)
	}

	return bytes.NewBuffer(bytesRep)
}

func main() {
	travis_api_endpoint := "https://api.travis-ci.org"
	target_username := "andrewrothstein"
	target_repo := "go-trigger"
	post_body := MakeBody(target_username, target_repo)
	post_url := fmt.Sprintf("%s/repo/%s%%2F%s/requests", travis_api_endpoint, target_username, target_repo)
	log.Printf("posting to %s...\n", post_url)
	req, err := http.NewRequest("POST", post_url, post_body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Travis-API-Version", "3")

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)

}
