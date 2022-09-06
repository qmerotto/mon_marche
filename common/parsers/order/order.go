package order

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"test_tech/common/database/models"
	"test_tech/common/parsers/product"
)

type Parser interface {
	Parse(message []byte) (*models.Order, error)
}

type parser struct {
	ProductParser product.Parser
}

func GetParser() *parser {
	return &parser{ProductParser: product.GetParser()}
}

func (p *parser) Parse(message []byte) (*models.Order, error) {
	msgStrs := bytes.Split(message, []byte("\n\n"))
	if len(msgStrs) != 2 {
		return nil, fmt.Errorf("invalid message format")
	}

	//TODO Parse each part in a goroutine
	dbOrder, err := p.parseHeader(msgStrs[0])
	if err != nil {
		return nil, err
	}

	products, err := p.ProductParser.Parse(msgStrs[1])
	if err != nil {
		return nil, err
	}

	dbOrder.Products = products

	return dbOrder, nil
}

func (p *parser) unmarshall(record []byte, order *models.Order) error {
	strs := strings.Split(string(record), "\n")

	id, err := strconv.ParseInt(strings.Split(strs[0], ": ")[1], 10, 32)
	if err != nil {
		return err
	}

	vat, err := strconv.ParseFloat(strings.Split(strs[1], ": ")[1], 32)
	if err != nil {
		return err
	}

	total, err := strconv.ParseFloat(strings.Split(strs[2], ": ")[1], 32)
	if err != nil {
		return err
	}

	order.ID = id
	order.VAT = float32(vat)
	order.Total = float32(total)

	return nil
}

func (p *parser) parseHeader(message []byte) (*models.Order, error) {
	order := models.Order{}

	if err := p.unmarshall(message, &order); err != nil {
		return nil, err
	}

	return &order, nil
}
