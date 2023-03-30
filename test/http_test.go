package test

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
	var (
		name = "test.jpg"
		url  = "https://api.ma"
	)
	client := &http.Client{}
	data, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(string(content))
}
