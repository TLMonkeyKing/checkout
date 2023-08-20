package item

type IItem interface {
	GetSKU() string
	GetPrice() int
	GetDiscountPrice() (int, error)
	GetAmountForDiscount() (int, error)
	HasDiscountAvailable() bool
}
