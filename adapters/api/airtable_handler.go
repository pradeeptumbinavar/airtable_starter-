package api

import (
    "context"
    "net/http"

    "github.com/go-chi/chi/v5"
    "airtable_starter/app"
    
)

type AirtableHandler struct {
    service *app.AirtableService
}

func NewAirtableHandler(service *app.AirtableService) *AirtableHandler {
    return &AirtableHandler{service: service}
}

func (h *AirtableHandler) GetRecord(w http.ResponseWriter, r *http.Request) {
    baseId := chi.URLParam(r, "baseId")
    table := chi.URLParam(r, "tableIdOrName")
    recordId := chi.URLParam(r, "recordId")

    resp, err := h.service.GetRecord(context.Background(), baseId, table, recordId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(resp.Body)
}
