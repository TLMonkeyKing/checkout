package item

import "fmt"

type Discount struct {
	Price  int
	Amount int
}

type Item struct {
	IItem
	SKU      string
	Price    int
	Discount *Discount
}

func (i Item) GetSKU() string {
	return i.SKU
}

func (i Item) GetPrice() int {
	return i.Price
}

func (i Item) GetDiscountPrice() (int, error) {
	if i.HasDiscountAvailable() {
		return i.Discount.Price, nil
	}

	return 0, fmt.Errorf("No discount for item %s", i.GetSKU())
}

func (i Item) GetAmountForDiscount() (int, error) {
	if i.HasDiscountAvailable() {
		return i.Discount.Amount, nil
	}

	return 0, fmt.Errorf("No discount for item %s", i.GetSKU())
}

func (i Item) HasDiscountAvailable() bool {
	return i.Discount != nil
}
