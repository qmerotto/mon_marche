package product

import (
	"github.com/stretchr/testify/assert"
	"test_tech/common/database/models"
	"testing"
)

func TestParseProduct(t *testing.T) {
	productsPayload := []byte(`product,product_id,price
Formule(s) midi,aZde,14.90
Café,IZ8z,2`)

	expectedProducts := []models.Product{
		{
			ID:    "aZde",
			Name:  "Formule(s) midi",
			Price: 14.90,
		},
		{
			ID:    "IZ8z",
			Name:  "Café",
			Price: 2.0,
		},
	}

	products, err := (&parser{}).Parse(productsPayload)
	assert.Equal(t, 2, len(products))
	assert.Equal(t, expectedProducts, products)
	assert.Nil(t, err)
}
