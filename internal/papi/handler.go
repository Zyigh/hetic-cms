package papi

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

func (h Handler) Prefix() string {
	return "/papi"
}

func (h Handler) Register(mux *goji.Mux) {
	mux.HandleFunc(pat.Get("page/:name"), h.GetPage)
}

func (h Handler) GetPage(w http.ResponseWriter, r *http.Request) {
	facade := GetPageForPAPI{}

	if err := facade.Deserialize(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"err": "%s"}`, strings.ReplaceAll(err.Error(), "\"", "'"))
		return
	}

	page, err := h.service.GetOnePage(r.Context(), facade.GetPageForPAPI)

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
