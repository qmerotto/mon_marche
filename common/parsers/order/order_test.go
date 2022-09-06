package order

import (
	"github.com/stretchr/testify/assert"
	"test_tech/common/database/models"
	"test_tech/common/parsers/product"
	"testing"
)

func TestParseOrder(t *testing.T) {
	orderPayload := []byte(`Order: 123456
VAT: 3.10
Total: 16.90

product,product_id,price
Formule(s) midi,aZde,14.90
Café,IZ8z,2`)

	expectedOrder := models.Order{
		ID:    123456,
		VAT:   3.10,
		Total: 16.90,
		Products: []models.Product{
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
		},
	}

	order, err := (&parser{ProductParser: product.GetParser()}).Parse(orderPayload)
	assert.Equal(t, expectedOrder, *order)
	assert.Nil(t, err)
}
