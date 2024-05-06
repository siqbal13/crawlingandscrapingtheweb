package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestCreateJson(t *testing.T) {
	// Create a temporary file for testing
	file, err := os.CreateTemp("", "output*.jl")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(file.Name()) // Clean up the temporary file

	// Create test data
	data := PageContent{
		URL:   "https://example.com",
		Title: "Test Title",
		Text:  "Lorem ipsum dolor sit amet",
	}

	// Call the CreateJson function
	err = CreateJson(file, data)
	if err != nil {
		t.Fatalf("CreateJson returned an error: %v", err)
	}

	// Read the content of the file
	file.Seek(0, 0)
	decoder := json.NewDecoder(file)
	var decodedData PageContent
	err = decoder.Decode(&decodedData)
	if err != nil {
		t.Fatalf("Error decoding JSON from file: %v", err)
	}

	// Check if the decoded data matches the original data
	if decodedData.URL != data.URL || decodedData.Title != data.Title || decodedData.Text != data.Text {
		t.Errorf("Decoded data does not match original data:\nExpected: %v\nGot: %v", data, decodedData)
	}
}
