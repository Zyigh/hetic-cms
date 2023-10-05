package page

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Zyigh/hetic-cms/internal"
	"github.com/Zyigh/hetic-cms/internal/clients"
	goji "goji.io"
	"goji.io/pat"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	service Service
}

func NewHandler(clts clients.Clients) Handler {
	return Handler{
		service: NewService(clts),
	}
}

func (h Handler) Register(mux *goji.Mux) {
	mux.HandleFunc(pat.Get("/"), h.ListPages)
	mux.HandleFunc(pat.Get("/:id"), h.GetOnePage)
}

func (h Handler) Prefix() string {
	return "/page"
}

func (h Handler) ListPages(w http.ResponseWriter, r *http.Request) {
	pages, err := h.service.ListPages(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		fmt.Fprintf(w, `{"err": "%s"}`, strings.ReplaceAll(err.Error(), "\"", "'"))
		return
	}

	if err := json.NewEncoder(w).Encode(pages); err != nil {
		log.Println(err)
	}
}

func (h Handler) GetOnePage(w http.ResponseWriter, r *http.Request) {
	facade := GetOnePage{}

	if err := facade.Deserializer(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"err": "%s"}`, strings.ReplaceAll(err.Error(), "\"", "'"))
		return
	}

	page, err := h.service.GetOnePage(r.Context(), facade.GetOnePage)

	if err != nil {
		if errors.Is(err, internal.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		fmt.Fprintf(w, `{"err": "%s"}`, strings.ReplaceAll(err.Error(), "\"", "'"))
		return
	}

	if err := json.NewEncoder(w).Encode(page); err != nil {
		log.Println(err)
	}
}
