package item

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ItemTestSuite struct {
	suite.Suite
}

func (suite *ItemTestSuite) SetupTest() {
}

func (suite *ItemTestSuite) Test_ItemSKUOK() {
	item := Item{
		SKU:   "Test",
		Price: 12,
	}

	assert.Equal(suite.T(), "Test", item.GetSKU())
}

func (suite *ItemTestSuite) Test_ItemPriceOK() {
	item := Item{
		SKU:   "Test",
		Price: 42,
	}

	assert.Equal(suite.T(), 42, item.GetPrice())
}

func (suite *ItemTestSuite) Test_ItemWithoutDiscount_HasDiscountAvailable_ReturnsFalse() {
	item := Item{
		SKU:   "Test",
		Price: 42,
	}

	assert.Equal(suite.T(), false, item.HasDiscountAvailable())
}

func (suite *ItemTestSuite) Test_ItemWithoutDiscount_GetDiscountAmount_ReturnsError() {
	item := Item{
		SKU:   "Test",
		Price: 42,
	}

	_, err := item.GetAmountForDiscount()
	assert.Error(suite.T(), err)
}

func (suite *ItemTestSuite) Test_ItemWithoutDiscount_GetDiscountPrice_ReturnsError() {
	item := Item{
		SKU:   "Test",
		Price: 42,
	}

	_, err := item.GetDiscountPrice()
	assert.Error(suite.T(), err)
}

func (suite *ItemTestSuite) Test_ItemWithDiscount_HasDiscountAvailable_ReturnsTrue() {
	item := Item{
		SKU:   "Test",
		Price: 42,
		Discount: &Discount{
			Price:  24,
			Amount: 3,
		},
	}

	assert.Equal(suite.T(), true, item.HasDiscountAvailable())
}

func (suite *ItemTestSuite) Test_ItemWithDiscount_GetDiscountAmountOK() {
	item := Item{
		SKU:   "Test",
		Price: 42,
		Discount: &Discount{
			Price:  24,
			Amount: 3,
		},
	}

	amount, err := item.GetAmountForDiscount()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 3, amount)
}

func (suite *ItemTestSuite) Test_ItemWithDiscount_GetDiscountPriceOK() {
	item := Item{
		SKU:   "Test",
		Price: 42,
		Discount: &Discount{
			Price:  24,
			Amount: 3,
		},
	}

	price, err := item.GetDiscountPrice()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 24, price)
}

func TestItemSuite(t *testing.T) {
	suite.Run(t, new(ItemTestSuite))
}
