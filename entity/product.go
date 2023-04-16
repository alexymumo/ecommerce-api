package entity

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID           uint      `gorm:"primary_key;auto_increment" json:"id"`
	ProductName  string    `gorm:"size:255;not null" json:"productname"`
	ProductImage string    `gorm:"size:255" json:"productimage"`
	Price        float64   `gorm:"size:255" json:"price"`
	Category     string    `gorm:"size:255" json:"category"`
	Description  string    `gorm:"size:255" json:"description"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (p *Product) Prepare() {
	p.ID = 0
	p.ProductName = html.EscapeString(strings.TrimSpace(p.ProductName))
	p.ProductImage = html.EscapeString(strings.TrimSpace(p.ProductImage))
	p.Category = html.EscapeString(html.EscapeString(p.Category))
	p.Description = html.EscapeString(html.EscapeString(p.Description))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Product) Validate() error {
	if p.ProductName == "" {
		return errors.New("ProductName Required")
	}
	if p.Category == "" {
		return errors.New("Category Required")
	}
	return nil
}

// save product
func (p *Product) SaveProduct(db *gorm.DB) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}
