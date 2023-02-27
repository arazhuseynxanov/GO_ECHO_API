package migrations

import (
	"fmt"
	"task25/database"
	"task25/models"
)

func Migration() {
	err := database.DB.AutoMigrate(
		&models.Products{},
		&models.Tag{},
	)
	database.DB.Model(&models.Products{}).Association("Products")
	database.DB.Model(&models.Tag{}).Association("Tag")
	//create
	//products := models.Products{
	//	Name:           "Anar",
	//	Price:          4.5,
	//	DiscountAmount: 1.01,
	//	IsNew:          true,
	//	Count:          1,
	//	Desc:           "araz",
	//	Code:           "3487284923",
	//	CostPrice:      1,
	//	Tags:           nil,
	//}
	//tags := models.Tag{
	//	Name: "Anar",
	//}
	//database.DB.Create(&products)
	//database.DB.Create(&tags)
	//end
	if err != nil {
		fmt.Println("Cannot Migration")
	}
	fmt.Println("Migrated Successfull")
}
