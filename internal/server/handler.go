package server

import goji "goji.io"

type Handler interface {
	Register(mux *goji.Mux)
	Prefix() string
}
