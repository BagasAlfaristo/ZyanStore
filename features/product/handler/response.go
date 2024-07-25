package handler

import "ZyanStore/features/product"

type productResponse struct {
	ID          uint    `json:"id" form:"id"`
	SKU         string  `json:"sku" form:"sku"`
	Brand       string  `json:"brand" form:"brand"`
	ProductName string  `json:"product_name" form:"product_name"`
	ImageURL    string  `json:"image_url" form:"image_url"`
	Duration    string  `json:"duration" form:"duration"`
	Description string  `json:"description" form:"description"`
	Category    string  `json:"category" form:"category"`
	Type        string  `json:"type" form:"type"`
	Tags        string  `json:"tags" form:"tags"`
	UnitCost    float64 `json:"unit_cost" form:"unit_cost"`
	RetailPrice float64 `json:"retail_price" form:"retail_price"`
	Stock       uint    `json:"stock" form:"stock"`
	Warranty    string  `json:"warranty" form:"warranty"`
	Status      string  `json:"status" form:"status"`
	Testimoni   string  `json:"testimony" form:"testimony"`
}

func allGormToCore(c product.ProductCore) productResponse {
	return productResponse{
		ID:          c.ID,
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
