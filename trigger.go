package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	fmt.Println("Hello world")
	travis_api_endpoint := "https://api.travis-ci.org"
	target_username := "andrewrothstein"
	target_repo := "go-trigger"
	post_url := fmt.Sprintf("%s/repo/%s%%2F%s/requests", travis_api_endpoint, target_username, target_repo)
	fmt.Printf("posting to %s...\n", post_url)
	resp, err := http.Post(post_url, "application/json", MakeBody(target_username, target_repo))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)

}
