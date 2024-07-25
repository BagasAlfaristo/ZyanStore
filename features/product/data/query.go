package data

import (
	"ZyanStore/features/product"

	"gorm.io/gorm"
)

type productModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.ProductModel {
	return &productModel{
		db: db,
	}
}

func (pm *productModel) AddProduct(product product.ProductCore) error {
	productGorm := coreToGorm(product)
	tx := pm.db.Create(&productGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (pm *productModel) GetProduct(category string) ([]product.ProductCore, error) {
	var productGorm []Product
	if category == "" {
		if err := pm.db.Find(&productGorm).Error; err != nil {
			return nil, err
		}
	} else {
		if err := pm.db.Where("category = ?", category).Find(&productGorm).Error; err != nil {
			return nil, err
		}
	}

	var productCore []product.ProductCore
	for _, v := range productGorm {
		productCore = append(productCore, gormToCore(v))
	}
	return productCore, nil
}

func (pm *productModel) Update(idProduct uint, updateData product.ProductCore) error {
	productGorm := coreToGorm(updateData)
	if err := pm.db.Model(&Product{}).Where("id = ?", idProduct).Updates(&productGorm).Error; err != nil {
		return err
	}
	return nil
}
