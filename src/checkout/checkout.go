package checkout

import "fmt"

type Checkout struct {
	ICheckout
}

func (c Checkout) Scan(item string) {
	fmt.Println("Scan: Not implemented")
}

func (c Checkout) GetTotalPrice() int {
	fmt.Println("GetTotalPrice: Not implemented")

	return 0
}
