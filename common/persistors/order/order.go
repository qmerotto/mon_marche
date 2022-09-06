package order

import (
	"gorm.io/gorm"
	"test_tech/common/database"
	"test_tech/common/database/models"
)

type Persistor interface {
	Create(order *models.Order) error
}

type order struct {
	Conn *gorm.DB
}

func GetPersistor() *order {
	return &order{Conn: database.DB}
}

func (o *order) Create(order *models.Order) error {
	return o.Conn.Create(order).Error
}
