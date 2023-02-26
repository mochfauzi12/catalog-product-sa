package entity

type Item struct {
	name     string
	codeItem string
}

func NewItem(nameItem string, codeItem string) *Item {
	return &Item{
		name:     nameItem,
		codeItem: codeItem,
	}
}

func (item *Item) GetName() string {
	return item.name
}

func (item *Item) GetCodeItem() string {
	return item.codeItem
}
