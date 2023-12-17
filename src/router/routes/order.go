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
	{
		URI:           "/show_orders",
		Method:        http.MethodGet,
		Func:          controllers.ShowAllOrders,
		Authorization: true,
		OnlyAdmin:     true,
	},
	{
		URI:           "/show_user_orders",
		Method:        http.MethodGet,
		Func:          controllers.ShowUserOrders,
		Authorization: true,
		OnlyAdmin:     false,
	},
	{
		URI:           "/change_order_status",
		Method:        http.MethodPut,
		Func:          controllers.ChangeOrderStatus,
		Authorization: true,
		OnlyAdmin:     true,
	},
}
