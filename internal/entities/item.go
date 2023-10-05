package entities

type Item struct {
	Name string `db:"name"`
	Page Page
}

type FullItem struct {
	Item Item `db:"item"`
	Page Page `db:"page"`
}

func (f FullItem) AsItem() Item {
	return Item{
		Name: f.Item.Name,
		Page: f.Page,
	}
}
