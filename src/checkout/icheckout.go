package checkout

type ICheckout interface {
	Scan(item string)
	GetTotalPrice() int
}
