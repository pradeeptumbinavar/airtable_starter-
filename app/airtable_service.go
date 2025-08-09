package app

import (
	"airtable_starter/gen/airtableapi"
	"context"
)

type AirtableService struct {
	client airtableapi.ClientWithResponsesInterface
}

func NewAirtableService(client airtableapi.ClientWithResponsesInterface) *AirtableService {
	return &AirtableService{client: client}
}

func (s *AirtableService) GetRecord(ctx context.Context, baseId, tableIdOrName, recordId string) (*airtableapi.GetRecordResponse, error) {
	return s.client.GetRecordWithResponse(ctx, baseId, tableIdOrName, recordId, nil)
}
