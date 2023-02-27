package models

type ProductTag struct {
	ID         int `json:"id" gorm:"primaryKey"`
	ProductsID int `json:"products_id"`
	TagID      int `json:"tag_id"`
}
