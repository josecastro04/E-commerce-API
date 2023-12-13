package routes

import (
	"api/src/controllers"
	"net/http"
)

var Product = []Route{
	{
		URI:           "/insert_product",
		Method:        http.MethodPost,
		Func:          controllers.InsertProduct,
		Authorization: true,
		OnlyAdmin:     true,
	},
	{
		URI:           "/show_product/{productID}",
		Method:        http.MethodGet,
		Func:          controllers.ShowProduct,
		Authorization: false,
		OnlyAdmin:     false,
	},
	{
		URI:           "/change_price/{productID}",
		Method:        http.MethodPut,
		Func:          controllers.ChangePrice,
		Authorization: true,
		OnlyAdmin:     true,
	},
	{
		URI:           "/delete_product/{productID}",
		Method:        http.MethodDelete,
		Func:          controllers.DeleteProduct,
		Authorization: true,
		OnlyAdmin:     true,
	},
	{
		URI:           "/update_product_image",
		Method:        http.MethodPut,
		Func:          controllers.UpdateImage,
		Authorization: true,
		OnlyAdmin:     true,
	},
}
