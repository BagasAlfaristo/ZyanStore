package data

import (
	"ZyanStore/features/product"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	SKU         string
	Brand       string
	ProductName string
	ImageURL    string
	Duration    string
	Description string
	Category    string
	Type        string
	Tags        string
	UnitCost    float64
	RetailPrice float64
	Stock       uint
	Warranty    string
	Status      string
	Testimoni   string
}

func coreToGorm(c product.ProductCore) Product {
	return Product{
		SKU:         c.SKU,
		Brand:       c.Brand,
		ProductName: c.ProductName,
		ImageURL:    c.ImageURL,
		Duration:    c.Duration,
		Description: c.Description,
		Category:    c.Category,
		Type:        c.Type,
		Tags:        c.Tags,
		UnitCost:    c.UnitCost,
		RetailPrice: c.RetailPrice,
		Stock:       c.Stock,
		Warranty:    c.Warranty,
		Status:      c.Status,
		Testimoni:   c.Testimoni,
	}
}

func gormToCore(g Product) product.ProductCore {
	return product.ProductCore{
		ID:          g.ID,
		SKU:         g.SKU,
		Brand:       g.Brand,
		ProductName: g.ProductName,
		ImageURL:    g.ImageURL,
		Duration:    g.Duration,
		Description: g.Description,
		Category:    g.Category,
		Type:        g.Type,
		Tags:        g.Tags,
		UnitCost:    g.UnitCost,
		RetailPrice: g.RetailPrice,
		Stock:       g.Stock,
		Warranty:    g.Warranty,
		Status:      g.Status,
		Testimoni:   g.Testimoni,
	}
}
