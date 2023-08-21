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

![image](https://github.com/TLMonkeyKing/checkout/assets/67118768/db5bd853-5690-4797-8d41-bfd537bc9fd2)

Edit: I later ran code coverage on the project and got the following results:
![image](https://github.com/TLMonkeyKing/checkout/assets/67118768/71252f15-87c0-4096-9d1c-1bd53be9aa85)
![image](https://github.com/TLMonkeyKing/checkout/assets/67118768/4cff1885-8788-460e-8439-9ce5c304ccaa)
![image](https://github.com/TLMonkeyKing/checkout/assets/67118768/f10aaa5a-c577-42dd-84f7-6318a391eb17)
