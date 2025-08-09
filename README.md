# Airtable Get Record Integration

This project demonstrates how to call the Airtable "Get record" endpoint using Go and oapi-codegen.

## Steps

1. Install `oapi-codegen`:
   ```bash
   go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
   ```

2. Generate code:
   ```bash
   oapi-codegen --config=oapi_config.yaml api/airtable.yaml
   ```

3. Run the demo client:
   ```bash
   go run cmd/airtable-client/main.go
   ```
