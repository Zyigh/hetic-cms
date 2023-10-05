package entities

type Menu struct {
	Name  string `db:"name"`
	Items []Item
}

type FullMenu struct {
	Menu Menu `db:"menu"`
	Item Item `db:"item"`
}

type FullMenuWithRows []FullMenu

func (f FullMenuWithRows) AsMenu() Menu {
	menu := Menu{}
	if len(f) < 1 {
		return menu
	}

	for _, row := range f {
		menu.Name = row.Menu.Name
		menu.Items = append(menu.Items, row.Item)
	}

	return menu
}
