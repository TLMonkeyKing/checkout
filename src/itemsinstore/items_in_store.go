package itemsinstore

import (
	"fmt"
	_item "toml/checkout/src/item"
)

type ItemsInStore struct {
	IItemsInStore
	Items []_item.IItem
}

func (i *ItemsInStore) GetItem(item string) (_item.IItem, error) {
	for _, it := range i.Items {
		if (it).GetSKU() == item {
			// Item is in store, return it
			return it, nil
		}
	}

	// Item not in store, return error
	return _item.Item{}, fmt.Errorf("No item in store for '%s'", item)
}

func (i *ItemsInStore) AddItem(item _item.IItem) {
	i.Items = append(i.Items, item)
}
