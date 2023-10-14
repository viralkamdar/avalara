package main

import (
	"os"
	"testing"
)

func TestGetWordMeaning(t *testing.T) {
	apiKey := os.Getenv("API_KEY")
	word := "exercise"
	expectedOutput := "ˈek-sər-ˌsīz (noun): the act of bringing into play or realizing in action : use"
	actualOutput, err := GetWordMeaning(&word, apiKey)
	if err != nil {
		t.Error("TestGetWordMeaning1 failed")
	}
	if expectedOutput != actualOutput {
		t.Errorf("Expected output \"%s\" doesn't match the actual output \"%s\"\n", expectedOutput, actualOutput)
	}
}

func TestGetWordMeaningNoApiKey(t *testing.T) {
	word := "exercise"
	expectedError := "Environment variable API_KEY is mandatory"
	output, actualError := GetWordMeaning(&word, "")
	if output != "" {
		t.Error("Expected blank output")
	}
	if actualError.Error() != expectedError {
		t.Errorf("Expected error \"%s\" but got \"%s\"\n", expectedError, actualError.Error())
	}
}

func TestGetWordMeaningNoWord(t *testing.T) {
	word := ""
	apiKey := os.Getenv("API_KEY")
	expectedError := "Please enter the word that you want to find the meaning of"
	output, actualError := GetWordMeaning(&word, apiKey)
	if output != "" {
		t.Error("Expected blank output")
	}
	if actualError.Error() != expectedError {
		t.Errorf("Expected error \"%s\" but got \"%s\"\n", expectedError, actualError.Error())
	}
}

func TestGetWordMeaningBadWord(t *testing.T) {
	word := "abcd1234"
	apiKey := os.Getenv("API_KEY")
	expectedError := "Error: No meaning found"
	output, actualError := GetWordMeaning(&word, apiKey)
	if output != "" {
		t.Error("Expected blank output")
	}
	if actualError.Error() != expectedError {
		t.Errorf("Expected error \"%s\" but got \"%s\"\n", expectedError, actualError.Error())
	}
}
