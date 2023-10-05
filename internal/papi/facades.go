package papi

import (
	"errors"
	"github.com/Zyigh/hetic-cms/hetic-cms/facades"
	"goji.io/pat"
	"net/http"
)

type GetPageForPAPI struct {
	facades.GetPageForPAPI
}

func (p *GetPageForPAPI) Deserialize(r *http.Request) error {
	name := pat.Param(r, "name")

	if name == "" {
		return errors.New("frr tabuse")
	}

	p.Name = name

	return nil
}
