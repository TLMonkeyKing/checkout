package checkout_test

import (
	"testing"
	"toml/checkout/src/checkout"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CheckoutTestSuite struct {
	suite.Suite
	Checkout checkout.Checkout
}

func (suite *CheckoutTestSuite) SetupTest() {
	suite.Checkout = checkout.Checkout{}
}

func (suite *CheckoutTestSuite) Test_ItWorks() {
	assert.Equal(suite.T(), 0, suite.Checkout.GetTotalCost())
}

func TestCheckoutSuite(t *testing.T) {
	suite.Run(t, new(CheckoutTestSuite))
}
