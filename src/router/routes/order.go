package routes

import (
	"api/src/controllers"
	"net/http"
)

var Orders = []Route{
	{
		URI:           "/place_order",
		Method:        http.MethodPost,
		Func:          controllers.PlaceOrder,
		Authorization: true,
		OnlyAdmin:     false,
	},
	{
		URI:           "/show_order/{orderID}",
		Method:        http.MethodGet,
		Func:          controllers.ShowOrder,
		Authorization: true,
		OnlyAdmin:     false,
	},
}
