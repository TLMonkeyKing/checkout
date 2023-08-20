package checkout

import (
	"fmt"
	"toml/checkout/src/itemsinstore"
)

type Checkout struct {
	ICheckout
	itemsInStore itemsinstore.IItemsInStore
	items        []string
}

func (c *Checkout) Scan(item string) {
	c.items = append(c.items, item)
}

func (c *Checkout) WithItemStore(store itemsinstore.IItemsInStore) {
	c.itemsInStore = store
}

func (c *Checkout) GetTotalPrice() int {
	total := 0

	// Determine counts of items in item list
	var countItems map[string]int = make(map[string]int)

	for _, i := range c.items {
		if _, ok := countItems[i]; !ok {
			countItems[i] = 0
		}

		countItems[i]++
	}

	// For each type of item in cart
	for key := range countItems {
		itemInStore, err := c.itemsInStore.GetItem(key)

		if err != nil {
			// We have scanned an erroneous item - ignore it
			continue
		}

		// Work out how many discounts we have for the item(s) that have been scanned
		var individual, discounts int
		if itemInStore.HasDiscountAvailable() {
			discountAmount, _ := itemInStore.GetAmountForDiscount()

			individual = countItems[key] % discountAmount
			discounts = countItems[key] / discountAmount
		} else {
			// Item has no discounts availble, just do them all individually
			individual = countItems[key]
			discounts = 0
		}

		// Add to running total
		if individual > 0 {
			total += individual * itemInStore.GetPrice()
		}

		if discounts > 0 {
			price, err := itemInStore.GetDiscountPrice()
			if err != nil {
				fmt.Println("Warning: Error redeeming item discount")
			}
			total += discounts * price
		}
	}

	return total
}
