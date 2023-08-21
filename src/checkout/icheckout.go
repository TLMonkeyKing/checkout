package checkout

type ICheckout interface {
	Scan(item string) error
	GetTotalPrice() (int, error)
}
