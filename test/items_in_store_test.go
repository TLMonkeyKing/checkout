package tests

import (
	"testing"
	_item "toml/checkout/src/item"
	"toml/checkout/src/itemsinstore"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ItemsInStoreTestSuite struct {
	suite.Suite
}

func (suite *ItemsInStoreTestSuite) SetupTest() {
}

func (suite *ItemsInStoreTestSuite) Test_ItemsInStore_GetNonexistentItem_ReturnsError() {
	items := itemsinstore.ItemsInStore{}
	_, err := items.GetItem("Test")

	assert.Error(suite.T(), err)
}

func (suite *ItemsInStoreTestSuite) Test_ItemsInStore_AddAndGetItemOK() {
	items := itemsinstore.ItemsInStore{}

	itemToAdd := _item.Item{
		SKU:   "Test",
		Price: 99,
	}

	items.AddItem(itemToAdd)

	returnedItem, err := items.GetItem("Test")

	assert.Equal(suite.T(), "Test", returnedItem.GetSKU())
	assert.Equal(suite.T(), 99, returnedItem.GetPrice())
	assert.NoError(suite.T(), err)
}

func TestItemsInStoreSuite(t *testing.T) {
	suite.Run(t, new(ItemsInStoreTestSuite))
}
