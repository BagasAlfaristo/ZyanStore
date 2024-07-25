package handler

import "ZyanStore/features/product"

type ProductRequest struct {
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

func reqToCore(input ProductRequest) product.ProductCore {
	return product.ProductCore{
		SKU:         input.SKU,
		Brand:       input.Brand,
		ProductName: input.ProductName,
		ImageURL:    input.ImageURL,
		Duration:    input.Duration,
		Description: input.Description,
		Category:    input.Category,
		Type:        input.Type,
		Tags:        input.Tags,
		UnitCost:    input.UnitCost,
		RetailPrice: input.RetailPrice,
		Stock:       input.Stock,
		Warranty:    input.Warranty,
		Status:      input.Status,
		Testimoni:   input.Testimoni,
	}
}
