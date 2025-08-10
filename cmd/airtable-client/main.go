package main

import (
	"airtable_starter/gen/airtableapi"
	"airtable_starter/pkg/httpclient"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read variables from environment
	token := os.Getenv("AIRTABLE_TOKEN")
	baseID := os.Getenv("AIRTABLE_BASE_ID")
	tableID := os.Getenv("AIRTABLE_TABLE_ID")
	recordID := os.Getenv("AIRTABLE_RECORD_ID_2") // Use the first record ID for this example

	// Create HTTP client with Bearer token
	client := httpclient.NewBearerClient(token)

	// Create Airtable API client
	airtableClient, err := airtableapi.NewClientWithResponses(
		"https://api.airtable.com/v0",
		airtableapi.WithHTTPClient(client),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Call the API to get a record
	resp, err := airtableClient.GetRecordWithResponse(
		context.Background(),
		baseID,
		tableID,
		recordID,
		&airtableapi.GetRecordParams{}, // keep empty for now
	)
	if err != nil {
		log.Fatal(err)
	}

	// Print raw response body
	fmt.Println(string(resp.Body))
}
