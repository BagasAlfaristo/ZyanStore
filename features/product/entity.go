package product

import "gorm.io/gorm"

type ProductCore struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey"`
	SKU         string  `gorm:"type:varchar(50);not null"`
	Brand       string  `gorm:"type:varchar(100);not null"`
	ProductName string  `gorm:"type:varchar(100);not null"`
	ImageURL    string  `gorm:"type:varchar(255)"`
	Duration    string  `gorm:"type:varchar(50)"`
	Description string  `gorm:"type:text"`
	Category    string  `gorm:"type:varchar(100);not null"`
	Type        string  `gorm:"type:varchar(50);not null"`
	Tags        string  `gorm:"type:varchar(255)"`
	UnitCost    float64 `gorm:"type:decimal(10,2);not null"`
	RetailPrice float64 `gorm:"type:decimal(10,2);not null"`
	Stock       uint    `gorm:"not null"`
	Warranty    string  `gorm:"type:varchar(50)"`
	Status      string  `gorm:"type:enum('Active', 'Inactive');not null"`
	Testimoni   string  `gorm:"type:text"`
}

type ProductModel interface {
	AddProduct(product ProductCore) error
	GetProduct(category string) ([]ProductCore, error)
	Update(idProduct uint, updateData ProductCore) error
	// GetProductByName(productName string) ([]ProductCore, error)
}

type ProductService interface {
	AddProduct(product ProductCore) error
	GetProduct(category string) ([]ProductCore, error)
	Update(idProduct uint, updateData ProductCore) error
	// GetProductByName(productName string) ([]ProductCore, error)
}
