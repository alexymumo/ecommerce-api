package entity

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	ID           uint      `gorm:"primary_key;auto_increment" json:"id"`
	ProductName  string    `gorm:"size:255;not null" json:"productname"`
	ProductImage string    `gorm:"size:255" json:"productimage"`
	Price        float64   `gorm:"size:255" json:"price"`
	Category     string    `gorm:"size:255" json:"category"`
	Description  string    `gorm:"size:255" json:"description"`
	User         User      `json:"user"`
	Rating       Rating    `json:"rating"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (p *Product) Prepare() {
	p.ID = 0
	p.ProductName = html.EscapeString(strings.TrimSpace(p.ProductName))
	p.ProductImage = html.EscapeString(strings.TrimSpace(p.ProductImage))
	p.Category = html.EscapeString(html.EscapeString(p.Category))
	p.Description = html.EscapeString(html.EscapeString(p.Description))
	p.User = User{}
	p.Rating = Rating{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Product) Validate() error {
	if p.ProductName == "" {
		return errors.New("ProductName Required")
	}
	if p.Category == "" {
		return errors.New("category required")
	}
	if p.Description == "" {
		return errors.New("description required")

	}
	return nil
}

// save product
func (p *Product) SaveProduct(db *gorm.DB) (*Product, error) {
	//var err error
	err := db.Debug().Model(&Product{}).Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) DeleteProduct(db *gorm.DB, productID uint64) (int64, error) {
	db = db.Debug().Model(&Product{}).Where("id = ?", productID).Take(&Product{}).Delete(&Product{})

	if db != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("not found")
		}
	}
	return db.RowsAffected, nil
}

func (p *Product) GetProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	return &products, nil
}

func (p *Product) SearchProductsByName(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Find(&products).Where("name LIKE ?").Error
	if err != nil {
		return &[]Product{}, err
	}
	return &products, nil
}

func (p *Product) UpdateProduct(db *gorm.DB) (*Product, error) {
	//var err error
	err := db.Debug().Model(&Product{}).Where("id = ?").Updates(Product{ProductName: p.ProductName, Price: p.Price}).Error
	if err != nil {
		return &Product{}, err
	}
	return p, err
}

func (p *Product) GetProductById(db *gorm.DB, productId uint64) (*Product, error) {
	//var err error
	err := db.Debug().Model(&Product{}).Where("id=?", productId).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}
	/*if p.ID != 0 {

	}*/
	return p, nil

}
