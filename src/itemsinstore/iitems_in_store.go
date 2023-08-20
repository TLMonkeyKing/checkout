package itemsinstore

import _item "toml/checkout/src/item"

type IItemsInStore interface {
	GetItem(item string) (_item.IItem, error)
	AddItem(item _item.IItem)
}
