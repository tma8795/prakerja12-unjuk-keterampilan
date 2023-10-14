package repositories

import (
	"ecommerce/configs"
	"ecommerce/models"
	"fmt"
)

func AddProducts(productsDB *models.Products) error {
	result := configs.DB.Create(&productsDB)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetProducts(productsList *[]models.Products) error {
	result := configs.DB.Find(&productsList)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetProduct(product *models.Products) error {
	result := configs.DB.First(&product, product.Id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteProduct(product *models.Products) error {
	result := configs.DB.Where("id = ?", product.Id).Delete(&product)
	if result.Error != nil {
		return fmt.Errorf("error")
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	} else {
		return nil
	}
}
