package tests

import (
	"testing"
	"toml/checkout/src/checkout"
	"toml/checkout/src/item"
	"toml/checkout/src/itemsinstore"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CheckoutTestSuite struct {
	suite.Suite
	Checkout checkout.ICheckout
}

func (suite *CheckoutTestSuite) SetupTest() {
	checkout := checkout.Checkout{}

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
	checkout.WithItemStore(&store)

	suite.Checkout = &checkout
}

func (suite *CheckoutTestSuite) Test_EmptyBasketPriceIsZero() {
	assert.Equal(suite.T(), 0, suite.Checkout.GetTotalPrice())
}

func (suite *CheckoutTestSuite) Test_ItemAPrices50() {
	suite.Checkout.Scan("A")
	assert.Equal(suite.T(), 50, suite.Checkout.GetTotalPrice())
}

func (suite *CheckoutTestSuite) Test_ItemADiscountPrices130() {
	// Scan A three times
	for i := 0; i < 3; i++ {
		suite.Checkout.Scan("A")
	}

	// Three As gives us a discount
	assert.Equal(suite.T(), 130, suite.Checkout.GetTotalPrice())
}

func (suite *CheckoutTestSuite) Test_ItemADiscountAppliesMultipleTimes() {
	// Scan A six times
	for i := 0; i < 6; i++ {
		suite.Checkout.Scan("A")
	}

	// We should get two discounts
	assert.Equal(suite.T(), 260, suite.Checkout.GetTotalPrice())
}

func (suite *CheckoutTestSuite) Test_LotsOfItemA() {
	// Scan A eight times
	for i := 0; i < 8; i++ {
		suite.Checkout.Scan("A")
	}

	// We should get two discounts, and two individual prices
	assert.Equal(suite.T(), 360, suite.Checkout.GetTotalPrice())
}

func (suite *CheckoutTestSuite) Test_ItemBPrices30() {
	suite.Checkout.Scan("B")
	assert.Equal(suite.T(), 30, suite.Checkout.GetTotalPrice())
}

func (suite *CheckoutTestSuite) Test_ItemBDiscountPrices45() {
	// Scan B twice
	for i := 0; i < 2; i++ {
		suite.Checkout.Scan("B")
	}

	// Two Bs gives us a discount
	assert.Equal(suite.T(), 45, suite.Checkout.GetTotalPrice())
}

func (suite *CheckoutTestSuite) Test_ItemBDiscountAppliesMultipleTimes() {
	// Scan B four times
	for i := 0; i < 4; i++ {
		suite.Checkout.Scan("B")
	}

	// We should get two discounts
	assert.Equal(suite.T(), 90, suite.Checkout.GetTotalPrice())
}

func (suite *CheckoutTestSuite) Test_LotsOfItemB() {
	// Scan B 5 times
	for i := 0; i < 5; i++ {
		suite.Checkout.Scan("B")
	}

	// We should get two discounts, and one individual price
	assert.Equal(suite.T(), 120, suite.Checkout.GetTotalPrice())
}

func (suite *CheckoutTestSuite) Test_ItemCPrices20() {
	suite.Checkout.Scan("C")
	assert.Equal(suite.T(), 20, suite.Checkout.GetTotalPrice())
}

func (suite *CheckoutTestSuite) Test_ItemCHasNoDiscount() {
	// Scan C four times
	for i := 0; i < 4; i++ {
		suite.Checkout.Scan("C")
	}

	// No discount applied; instead we get 4 individual prices
	assert.Equal(suite.T(), 80, suite.Checkout.GetTotalPrice())
}

func (suite *CheckoutTestSuite) Test_ItemDPrices15() {
	suite.Checkout.Scan("D")
	assert.Equal(suite.T(), 15, suite.Checkout.GetTotalPrice())
}

func (suite *CheckoutTestSuite) Test_ItemDHasNoDiscount() {
	// Scan D four times
	for i := 0; i < 4; i++ {
		suite.Checkout.Scan("D")
	}

	// No discount applied; instead we get 4 individual prices
	assert.Equal(suite.T(), 60, suite.Checkout.GetTotalPrice())
}

func (suite *CheckoutTestSuite) Test_ItemDiscountsApplyBetweenItems() {
	// Given we scan A six times, B three times, C twice and D five times, in an arbitrary order
	suite.Checkout.Scan("D")
	suite.Checkout.Scan("A")
	suite.Checkout.Scan("B")
	suite.Checkout.Scan("A")
	suite.Checkout.Scan("C")
	suite.Checkout.Scan("D")
	suite.Checkout.Scan("A")
	suite.Checkout.Scan("B")
	suite.Checkout.Scan("A")
	suite.Checkout.Scan("A")
	suite.Checkout.Scan("D")
	suite.Checkout.Scan("B")
	suite.Checkout.Scan("D")
	suite.Checkout.Scan("A")
	suite.Checkout.Scan("D")
	suite.Checkout.Scan("C")

	// SIX As - two discounts: 260
	// THREE Bs - one discount, one individual: 75
	// FOUR Cs - no discounts: 40
	// FIVE Ds - no discounts: 75
	assert.Equal(suite.T(), (260 + 75 + 40 + 75), suite.Checkout.GetTotalPrice())
}

func TestCheckoutSuite(t *testing.T) {
	suite.Run(t, new(CheckoutTestSuite))
}
