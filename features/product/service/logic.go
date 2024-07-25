package service

import (
	"ZyanStore/features/product"
)

type productService struct {
	productModel product.ProductModel
}

func New(pm product.ProductModel) product.ProductService {
	return &productService{
		productModel: pm,
	}
}

func (ps *productService) AddProduct(productCore product.ProductCore) error {
	return ps.productModel.AddProduct(productCore)
}

func (pm *productService) GetProduct(category string) ([]product.ProductCore, error) {
	return pm.productModel.GetProduct(category)
}

func (pm *productService) Update(idProduct uint, updateData product.ProductCore) error {
	return pm.productModel.Update(idProduct, updateData)
}
