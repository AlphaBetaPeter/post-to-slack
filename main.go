package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"os"
	"flag"
)

func main() {
	url := flag.String("url", "url", "The slack webhook url");
	text := flag.String("text", "text", "The slack message you want to post");

	flag.Parse()

	if *url == "url" {
		fmt.Println("Please provide a webhook url")
		os.Exit(0)
	}

	if *text == "text" {
		fmt.Println("Please provide text")
		os.Exit(0)
	}

	fmt.Println("Webhook URL: ", *url)

	jsonValue, _ := json.Marshal(map[string]string{
		"text": *text,
		"username": "Slackr",
		"icon_emoji": ":monkey:",
	})
	jsonString := bytes.NewBuffer(jsonValue)
	fmt.Println("Sending payload", jsonString)

	req, err := http.NewRequest("POST", *url, jsonString)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
		fmt.Printf("Could not post %s to %s\n", jsonValue, *url)
        panic(err)
    }
    defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response body:", string(body))

	if resp.Status != "200 OK" {
		os.Exit(1);
	}
}
