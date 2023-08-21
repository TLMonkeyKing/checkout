package checkout

import (
	"testing"
	test_builder "toml/checkout/test/builder"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CheckoutTestSuite struct {
	suite.Suite
	Checkout ICheckout
}

func (suite *CheckoutTestSuite) SetupTest() {
	checkout := Checkout{}

	checkout.WithItemStore(test_builder.TestCheckoutConfiguration())
	suite.Checkout = &checkout
}

func (suite *CheckoutTestSuite) Test_EmptyBasketPriceIsZero() {
	total, err := suite.Checkout.GetTotalPrice()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 0, total)
}

func (suite *CheckoutTestSuite) Test_ItemAPrices50() {
	err := suite.Checkout.Scan("A")
	assert.NoError(suite.T(), err)

	total, err := suite.Checkout.GetTotalPrice()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 50, total)
}

func (suite *CheckoutTestSuite) Test_ItemADiscountPrices130() {
	// Scan A three times
	for i := 0; i < 3; i++ {
		suite.Checkout.Scan("A")
	}

	// Three As gives us a discount
	total, err := suite.Checkout.GetTotalPrice()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 130, total)
}

func (suite *CheckoutTestSuite) Test_ItemADiscountAppliesMultipleTimes() {
	// Scan A six times
	for i := 0; i < 6; i++ {
		suite.Checkout.Scan("A")
	}

	// We should get two discounts
	total, err := suite.Checkout.GetTotalPrice()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 260, total)
}

func (suite *CheckoutTestSuite) Test_LotsOfItemA() {
	// Scan A eight times
	for i := 0; i < 8; i++ {
		suite.Checkout.Scan("A")
	}

	// We should get two discounts, and two individual prices
	total, err := suite.Checkout.GetTotalPrice()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 360, total)
}

func (suite *CheckoutTestSuite) Test_ItemBPrices30() {
	err := suite.Checkout.Scan("B")
	assert.NoError(suite.T(), err)

	total, err := suite.Checkout.GetTotalPrice()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 30, total)
}

func (suite *CheckoutTestSuite) Test_ItemBDiscountPrices45() {
	// Scan B twice
	for i := 0; i < 2; i++ {
		suite.Checkout.Scan("B")
	}

	// Two Bs gives us a discount
	total, err := suite.Checkout.GetTotalPrice()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 45, total)
}

func (suite *CheckoutTestSuite) Test_ItemBDiscountAppliesMultipleTimes() {
	// Scan B four times
	for i := 0; i < 4; i++ {
		suite.Checkout.Scan("B")
	}

	// We should get two discounts
	total, err := suite.Checkout.GetTotalPrice()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 90, total)
}

func (suite *CheckoutTestSuite) Test_LotsOfItemB() {
	// Scan B 5 times
	for i := 0; i < 5; i++ {
		suite.Checkout.Scan("B")
	}

	// We should get two discounts, and one individual price
	total, err := suite.Checkout.GetTotalPrice()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 120, total)
}

func (suite *CheckoutTestSuite) Test_ItemCPrices20() {
	err := suite.Checkout.Scan("C")
	assert.NoError(suite.T(), err)

	total, err := suite.Checkout.GetTotalPrice()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 20, total)
}

func (suite *CheckoutTestSuite) Test_ItemCHasNoDiscount() {
	// Scan C four times
	for i := 0; i < 4; i++ {
		suite.Checkout.Scan("C")
	}

	// No discount applied; instead we get 4 individual prices
	total, err := suite.Checkout.GetTotalPrice()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 80, total)
}

func (suite *CheckoutTestSuite) Test_ItemDPrices15() {
	err := suite.Checkout.Scan("D")
	assert.NoError(suite.T(), err)

	total, err := suite.Checkout.GetTotalPrice()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 15, total)
}

func (suite *CheckoutTestSuite) Test_ItemDHasNoDiscount() {
	// Scan D four times
	for i := 0; i < 4; i++ {
		suite.Checkout.Scan("D")
	}

	// No discount applied; instead we get 4 individual prices
	total, err := suite.Checkout.GetTotalPrice()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 60, total)
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
	total, err := suite.Checkout.GetTotalPrice()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), (260 + 75 + 40 + 75), total)
}

func (suite *CheckoutTestSuite) Test_ItemScannedThatDoesntExist() {
	// Given we scan A, B, C and D and have a running total
	suite.Checkout.Scan("A")
	suite.Checkout.Scan("B")
	suite.Checkout.Scan("C")
	suite.Checkout.Scan("D")

	expectedTotal := (50 + 30 + 20 + 15)
	total, err := suite.Checkout.GetTotalPrice()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedTotal, total)

	// When we scan an item not in store
	err = suite.Checkout.Scan("E")

	// Then we get an error
	assert.Error(suite.T(), err)
}

func (suite *CheckoutTestSuite) Test_CheckoutWithEmptyStock() {
	// Given we create a new checkout and don't give it a stock
	checkout := Checkout{}

	// When we scan an item
	err := checkout.Scan("A")

	// Then we get an error
	assert.Error(suite.T(), err)
}

func TestCheckoutSuite(t *testing.T) {
	suite.Run(t, new(CheckoutTestSuite))
}
