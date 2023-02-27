package controller

import (
	"task25/database"
	"task25/models"
)

func productController() {

}
func CreateProduct(name string, price float64, discountAmount float64, isNew bool, count int, desc string, code string, costPrice float64, tag_ids []*models.Tag) (*models.Products, error) {
	product := &models.Products{
		Name:           name,
		Price:          price,
		DiscountAmount: discountAmount,
		IsNew:          isNew,
		Count:          count,
		Desc:           desc,
		Code:           code,
		CostPrice:      costPrice,
		Tags:           tag_ids,
	}
	for _, id := range tag_ids {
		tag := &models.Tag{}
		if err := database.DB.First(tag, id).Error; err != nil {
			return nil, err
		}
		product.Tags = append(product.Tags, tag)
	}
	if err := database.DB.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func GetProduct(id int) (*models.Products, error) {

	product := &models.Products{}
	if err := database.DB.Preload("tags").First(product, id).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func UpdateProduct(id int, name string, price float64, discountAmount float64, isNew bool, count int, desc string, code string, costPrice float64, tag_ids []*models.Tag) (*models.Products, error) {
	product := &models.Products{}
	if err := database.DB.First(product, id).Error; err != nil {
		return nil, err
	}

	product.Name = name
	product.Price = price
	product.DiscountAmount = discountAmount
	product.IsNew = isNew
	product.Count = count
	product.Desc = desc
	product.Code = code
	product.CostPrice = costPrice
	product.Tags = tag_ids

	if err := database.DB.Save(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}
func DeleteProduct(id int) error {
	product := &models.Products{}
	if err := database.DB.First(product, id).Error; err != nil {
		return err
	}
	if err := database.DB.Delete(product).Error; err != nil {
		return err
	}
	return nil
}
