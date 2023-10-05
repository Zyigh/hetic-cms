package page

import (
	"fmt"
	"github.com/Zyigh/hetic-cms/hetic-cms/facades"
	"github.com/google/uuid"
	"goji.io/pat"
	"net/http"
)

type GetOnePage struct {
	facades.GetOnePage
}

func (g *GetOnePage) Deserializer(r *http.Request) error {
	idRaw := pat.Param(r, "id")

	id, err := uuid.Parse(idRaw)

	if err != nil {
		return fmt.Errorf("frr tabuse, on a dit un uuid: %w", err)
	}

	g.ID = id

	return nil
}
