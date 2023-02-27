package models

type Products struct {
	//gorm.Model
	ID             int     `json:"id" from:"id" gorm:"primaryKey"`
	Name           string  `json:"name" form:"name" gorm:"not null"`
	Price          float64 `json:"price" form:"price" gorm:"not null" validate:"required"`
	DiscountAmount float64 `json:"discountAmount" form:"discountAmount" gorm:"not null" validate:"required"`
	IsNew          bool    `json:"isNew" form:"isNew"`
	Count          int     `json:"count" form:"count"`
	Desc           string  `json:"desc" form:"desc" gorm:"not null"`
	Code           string  `json:"code" form:"code" gorm:"not null"`
	CostPrice      float64 `json:"costPrice" form:"costPrice" gorm:"not null" validate:"required"`
	Tags           []*Tag  `json:"tags" form:"tags" gorm:"many2many:product_tags"`
}
