package handler

import (
	"ZyanStore/features/product"
	"ZyanStore/utils/responses"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productService product.ProductService
}

func New(ps product.ProductService) *productHandler {
	return &productHandler{
		productService: ps,
	}
}

func (ph *productHandler) AddProduct(c echo.Context) error {
	var newProductReq ProductRequest
	errBind := c.Bind(&newProductReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Err bind data"+errBind.Error(), nil))
	}

	inputCore := reqToCore(newProductReq)
	errIns := ph.productService.AddProduct(inputCore)
	if errIns != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(errIns.Error(), nil))
	} else if errIns == nil {
		return c.JSON(http.StatusCreated, responses.JSONWebResponse("Add product successfull!", nil))
	} else {
		return c.JSON(http.StatusBadGateway, responses.JSONWebResponse("Unknown err!!", nil))
	}
}

func (ph *productHandler) GetProduct(c echo.Context) error {
	category := c.QueryParam("category")
	log.Println("category", category)

	products, err := ph.productService.GetProduct(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error retrieving products: "+err.Error(), nil))
	}

	var response []productResponse
	for _, v := range products {
		response = append(response, allGormToCore(v))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("Product retrive successfully", response))
}

func (ph *productHandler) UpdatebyID(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(errConv.Error()+"Error parsing ID", idConv))
	}

	updateData := ProductRequest{}
	errBind := c.Bind(&updateData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Err bind data"+errBind.Error(), nil))
	}

	inputCore := reqToCore(updateData)
	errUpdate := ph.productService.Update(uint(idConv), inputCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to update product", errUpdate))
	} else {
		return c.JSON(http.StatusOK, responses.JSONWebResponse("Product update successfull", nil))
	}
}
