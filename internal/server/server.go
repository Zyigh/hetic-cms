package server

import (
	"fmt"
	"github.com/Zyigh/hetic-cms/internal/config"
	"github.com/Zyigh/hetic-cms/internal/server/middlewares"
	goji "goji.io"
	"goji.io/pat"
	"log"
	"net/http"
)

type HTTPServer struct {
	root *http.Server
}

func New(conf config.App, handlers []Handler) HTTPServer {
	mux := goji.NewMux()
	mux.Use(middlewares.ContentTypeJSON)

	for _, handler := range handlers {
		subMux := goji.SubMux()
		handler.Register(subMux)
		mux.Handle(
			pat.New(fmt.Sprintf("%s/*", handler.Prefix())),
			subMux,
		)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: mux,
	}

	return HTTPServer{
		root: srv,
	}
}

func (s HTTPServer) Run() error {
	log.Println("Server running at 127.0.0.1:8080")

	return s.root.ListenAndServe()
}
