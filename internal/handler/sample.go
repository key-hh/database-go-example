package handler

import (
	"go-database/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type SampleHandler struct {
	srv *service.SampleService
}

func NewSampleHandler(srv *service.SampleService) *SampleHandler {
	return &SampleHandler{
		srv: srv,
	}
}

func (h *SampleHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := h.srv.Create(ctx); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	resp := struct {
		Result string `json:"result""`
	}{
		Result: "ok",
	}
	writeResponse(w, resp)
}

func (h *SampleHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	resp, err := h.srv.Get(ctx, id)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeResponse(w, resp)
}

func (h *SampleHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp, err := h.srv.List(ctx)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeResponse(w, resp)
}
