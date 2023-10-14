package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Define the Api key to use in query

func main() {
	// Define command line arguments
	word := flag.String("word", "", "Word to find meaning of")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	apiKey := os.Getenv("API_KEY")
	meaning, err := GetWordMeaning(word, apiKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(meaning)
}

func GetWordMeaning(word *string, apiKey string) (string, error) {
	if apiKey == "" {
		return "", errors.New("Environment variable API_KEY is mandatory")
	}
	if *word == "" {
		return "", errors.New("Please enter the word that you want to find the meaning of")
	}
	url := "https://www.dictionaryapi.com/api/v3/references/collegiate/json/%s?key=%s"
	// Send a get request
	resp, err := http.Get(fmt.Sprintf(url, *word, apiKey))
	if err != nil {
		return "", errors.New("Failure calling the webster API")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonRes := []map[string]interface{}{}
	// Parse JSON into an array of map
	err = json.Unmarshal(body, &jsonRes)
	if err != nil || len(jsonRes) == 0 {
		return "", errors.New("Error: No meaning found")
	}
	mw := jsonRes[0]["hwi"].(map[string]interface{})["prs"].([]interface{})[0].(map[string]interface{})["mw"]
	wordType := jsonRes[0]["fl"]
	shortdef := jsonRes[0]["shortdef"].([]interface{})[0]
	return fmt.Sprintf("%s (%s): %s", mw, wordType, shortdef), nil
}
