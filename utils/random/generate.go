package generator

import (
	"github.com/RianIhsan/go-topup-midtrans/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GeneratorInterface interface {
	GenerateUUID() (string, error)
	GenerateOrderID() (string, error)
}

type Generator struct {
	db *gorm.DB
}

func NewGeneratorUUID(db *gorm.DB) *Generator {
	return &Generator{
		db: db,
	}
}

func (g *Generator) GenerateUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (g *Generator) GenerateOrderID() (string, error) {
	for {
		id, err := g.GenerateUUID()
		if err != nil {
			return "", err
		}

		// Check if the ID exists
		exists, err := g.checkIDExists(id)
		if err != nil {
			return "", err
		}

		if !exists {
			return id, nil
		}
	}
}

func (g *Generator) checkIDExists(orderID string) (bool, error) {
	var count int64
	if err := g.db.Model(&entities.MstBalance{}).Where("order_id = ?", orderID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
