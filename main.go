package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"task25/controller"
	"task25/database"
	"task25/models"
)

func main() {
	//Connected Database
	database.DatabaseInit()

	//controller.DeleteProduct(1)
	//end
	//controller.DeleteTag(4)
	//controller.CreateTag("Araz", nil)
	//controller.CreateTag("Emil", nil)
	//controller.CreateTag("Rafael", nil)
	//controller.CreateTag("Seymur", nil)

	//controller.CreateProduct("Araz", 3.4, 2, true, 1, "saalam", "wdwqd", 43.5, nil)
	//controller.CreateProduct("rafael", 3.4, 2, true, 1, "saalam", "wdwqd", 43.5, nil)
	//controller.CreateProduct("Seymur", 3.4, 2, true, 1, "saalam", "wdwqd", 43.5, nil)
	//controller.CreateProduct("Emil", 3.4, 2, true, 1, "saalam", "wdwqd", 43.5, nil)

	//controller.UpdateTag(2, "rafael", nil)
	//fmt.Println(controller.GetTag(2))
	//Migration
	//migrations.Migration()
	//end

	e := echo.New()
	e.GET("/products", func(c echo.Context) error {
		var product []models.Products
		if err := database.DB.Find(&product).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, "Veritabanı hatası")
		}
		return c.JSON(http.StatusOK, product)
	})
	e.GET("/tags", func(c echo.Context) error {
		var tag []models.Tag
		if err := database.DB.Find(&tag).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, "Veritabanı hatası")
		}
		return c.JSON(http.StatusOK, tag)
	})

	e.DELETE("/tags/:id", func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))
		if controller.DeleteTag(id) != nil {
			return c.JSON(http.StatusNotFound, "Not Found")
		} else {
			return c.JSON(http.StatusNoContent, controller.DeleteTag(id))
		}

	})

	e.DELETE("/products/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		if controller.DeleteProduct(id) != nil {
			return c.JSON(http.StatusNotFound, "Not Found")
		} else {
			return c.JSON(http.StatusNoContent, controller.DeleteProduct(id))
		}
	})
	e.POST("/products", createProductHandler)
	e.POST("/tags", createTagHandler)
	e.PUT("/tags/:id", updateTag)
	e.PUT("/products/:id", updateProduct)

	e.Logger.Fatal(e.Start(":1323"))
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func createProductHandler(c echo.Context) error {
	var productInfo models.Products
	if err := c.Bind(&productInfo); err != nil {
		return err
	}
	fmt.Println("--> ", productInfo)
	product, err := controller.CreateProduct(
		productInfo.Name,
		productInfo.Price,
		productInfo.DiscountAmount,
		productInfo.IsNew,
		productInfo.Count,
		productInfo.Desc,
		productInfo.Code,
		productInfo.CostPrice,
		productInfo.Tags,
	)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, product)
}
func createTagHandler(c echo.Context) error {
	var tagInfo models.Tag
	if err := c.Bind(&tagInfo); err != nil {
		return err
	}
	tag, err := controller.CreateTag(
		tagInfo.Name,
		tagInfo.Products,
	)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, tag)
}

func updateTag(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid tag ID")
	}
	var tagInfo models.Tag
	if err := c.Bind(&tagInfo); err != nil {
		return err
	}
	if err := c.Bind(&tagInfo); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	updatedTag, err := controller.UpdateTag(id, tagInfo.Name, tagInfo.Products)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to update tag")
	}
	return c.JSON(http.StatusOK, updatedTag)
}
func updateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid tag ID")
	}
	var productInfo models.Products
	if err := c.Bind(&productInfo); err != nil {
		return err
	}
	if err := c.Bind(&productInfo); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	updatedProduct, err := controller.UpdateProduct(id, productInfo.Name, productInfo.Price, productInfo.DiscountAmount, productInfo.IsNew, productInfo.Count, productInfo.Desc, productInfo.Code, productInfo.CostPrice, productInfo.Tags)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to update tag")
	}
	return c.JSON(http.StatusOK, updatedProduct)
}
