package test_builder

import (
	"toml/checkout/src/item"
	"toml/checkout/src/itemsinstore"
)

func TestCheckoutConfiguration() itemsinstore.IItemsInStore {
	store := itemsinstore.ItemsInStore{}
	store.AddItem(item.Item{
		SKU:   "A",
		Price: 50,
		Discount: &item.Discount{
			Price:  130,
			Amount: 3,
		},
	})
	store.AddItem(item.Item{
		SKU:   "B",
		Price: 30,
		Discount: &item.Discount{
			Price:  45,
			Amount: 2,
		},
	})
	store.AddItem(item.Item{
		SKU:   "C",
		Price: 20,
	})
	store.AddItem(item.Item{
		SKU:   "D",
		Price: 15,
	})

	return &store
}
