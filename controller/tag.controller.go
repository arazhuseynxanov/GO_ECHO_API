package controller

import (
	"task25/database"
	"task25/models"
)

func tagController() {

}
func CreateTag(name string, product_ids []*models.Products) (*models.Tag, error) {
	tag := &models.Tag{
		Name:     name,
		Products: product_ids,
	}
	for _, id := range product_ids {
		product := &models.Products{}
		if err := database.DB.First(product, id).Error; err != nil {
			return nil, err
		}
		tag.Products = append(tag.Products, product)
	}
	if err := database.DB.Create(tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}
func GetTag(id int) (*models.Tag, error) {

	tag := &models.Tag{}
	if err := database.DB.Preload("tags").First(tag, id).Error; err != nil {
		return nil, err
	}
	return tag, nil
}
func UpdateTag(id int, name string, product_ids []*models.Products) (*models.Tag, error) {
	tag := &models.Tag{}
	if err := database.DB.First(tag, id).Error; err != nil {
		return nil, err
	}
	tag.Name = name
	tag.Products = product_ids
	if err := database.DB.Save(tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}
func DeleteTag(id int) error {
	tag := &models.Tag{}
	if err := database.DB.First(tag, id).Error; err != nil {
		return err
	}
	if err := database.DB.Delete(tag).Error; err != nil {
		return err
	}
	return nil
}
