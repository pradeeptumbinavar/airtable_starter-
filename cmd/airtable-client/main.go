package main

import (
	"context"
	"fmt"
	"log"

	"airtable_starter/gen/airtableapi"
	"airtable_starter/pkg/httpclient"
)

func main() {
	token := "YOUR_AIRTABLE_TOKEN"
	client := httpclient.NewBearerClient(token)

	airtableClient, err := airtableapi.NewClientWithResponses("https://api.airtable.com/v0", airtableapi.WithHTTPClient(client))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := airtableClient.GetRecordWithResponse(
		context.Background(),
		"baseId",
		"tableName",
		"recordId",
		&airtableapi.GetRecordParams{}, // provide empty params or set fields as needed
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resp.Body))
}
