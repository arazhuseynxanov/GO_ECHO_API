package models

type Tag struct {
	//gorm.Model
	ID       int         `json:"id" form:"id" gorm:"primaryKey"`
	Name     string      `json:"name" form:"name" gorm:"not null"`
	Products []*Products `json:"products" form:"products" gorm:"many2many:product_tags;"`
}
