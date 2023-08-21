# checkout
This is my implementation of the checkout kata

Contained is a go struct "Checkout" which implements the given interface:
```
type ICheckout interface {
	Scan(item string) error
	GetTotalPrice() (int, error)
}
```

Price of items in store is separate from the checkout logic - the available items in store is controlled by a seperate struct `ItemsInStore`, which takes a JSON-esque config, and determines how the checkout gets it's prices and discounts for scanned items.

Contained is a series of tests that show that, when given the prices defined in the requirements, the checkout behaves as expected. Additionally, unit tests have been provided for the supporting structs `ItemsInStore` and `Item`
