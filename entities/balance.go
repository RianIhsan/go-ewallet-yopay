package entities

import "time"

type MstBalance struct {
	Id        int       `gorm:"primaryKey"`
	UserId    int       `gorm:"not null"`
	Amount    float64   `gorm:"not null"`
	OrderID   string    `gorm:"not null"`
	Status    string    `gorm:"type:varchar(50);default:'pending'"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (MstBalance) TableName() string {
	return "mst_balance"
}
