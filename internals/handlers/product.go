package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/abefiker/go_ecommerce/internals/config"
	"github.com/abefiker/go_ecommerce/internals/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Product(c echo.Context) error {
	return c.String(http.StatusOK, "here is product")
}
func CreateProduct(c echo.Context) error {
	product := &models.Product{}
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product data")
	}

	result := config.DB().Create(&product) // Persist product to database
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Error creating product")
	}

	// Optionally return the newly created product with its ID for confirmation
	return c.JSON(http.StatusCreated, product)
}

func GetProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product ID")
	}
	var product models.Product
	result := config.DB().First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching product")
	}

	return c.JSON(http.StatusOK, product)
}

func GetAllUsers(c echo.Context) error {
	var products []models.Product
	result := config.DB().Find(&products)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching products")
	}

	return c.JSON(http.StatusOK, products)
}

func UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product ID")
	}

	// Fetch product by ID (optional for validation)
	var existingProduct models.Product
	result := config.DB().First(&existingProduct, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "Product not found")
		}
		return c.JSON(http.StatusInternalServerError, "Error fetching product")
	}

	// Bind updated data from request body
	var updateProduct models.Product
	if err := c.Bind(&updateProduct); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid update data")
	}

	// Update specific fields based on non-zero values in updateProduct
	if updateProduct.Name != "" {
		existingProduct.Name = updateProduct.Name
	}
	if updateProduct.Description != "" {
		existingProduct.Description = updateProduct.Description
	}
	if updateProduct.Price != 0 {
		existingProduct.Price = updateProduct.Price
	}
	if updateProduct.Stock != 0 {
		existingProduct.Stock = updateProduct.Stock
	}
	if updateProduct.Images != nil {
		existingProduct.Images = updateProduct.Images
	}
	// ... Update other fields as needed

	// Update the product in the database
	result = config.DB().Updates(&existingProduct) // Update existing product record
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Error updating product")
	}

	return c.JSON(http.StatusOK, existingProduct) // Optionally return the updated product
}

func DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	  return c.JSON(http.StatusBadRequest, "Invalid product ID")
	}
  
	result := config.DB().Delete(&models.Product{}, id) // Delete product by ID
	if result.Error != nil {
	  if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, "Product not found")
	  }
	  return c.JSON(http.StatusInternalServerError, "Error deleting product")
	}
  
	return c.NoContent(http.StatusNoContent) // No content to return upon successful deletion
  }
  