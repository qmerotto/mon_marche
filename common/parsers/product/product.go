package product

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"test_tech/common/database/models"
)

type Parser interface {
	Parse(message []byte) ([]models.Product, error)
}

type parser struct{}

func GetParser() *parser {
	return &parser{}
}

func (p *parser) Parse(message []byte) ([]models.Product, error) {
	productsCount := len(bytes.Split(message, []byte("\n"))) - 1
	products := make([]models.Product, productsCount)

	reader := bytes.NewReader(message)
	csvReader := csv.NewReader(reader)

	header, err := csvReader.Read()
	if err != nil {
		return nil, err
	}

	if err = p.validateHeader(header); err != nil {
		return nil, err
	}

	for i := 0; i < productsCount; i++ {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if err = p.validateRecord(record); err != nil {
			return nil, err
		}

		if err = p.unmarshallCSV(record, &products[i]); err != nil {
			return nil, err
		}
	}

	return products, nil
}

func (p *parser) unmarshallCSV(record []string, product *models.Product) error {
	if product == nil {
		product = &models.Product{}
	}

	price, err := strconv.ParseFloat(record[2], 32)
	if err != nil {
		return err
	}

	product.ID = record[1]
	product.Name = record[0]
	product.Price = float32(price)

	return nil
}

func (p *parser) validateHeader(headers []string) error {
	if len(headers) != 3 {
		return fmt.Errorf("invalid header format")
	}

	return nil
}

func (p *parser) validateRecord(record []string) error {
	if len(record) != 3 {
		return fmt.Errorf("invalid record format")
	}

	return nil
}
