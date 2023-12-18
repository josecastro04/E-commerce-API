package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	var order models.Order

	if err = json.Unmarshal(bodyRequest, &order); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	order.UserID = userID

	db, err := database.ConnectWithDatabase()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	orderRepository := repositories.NewRepositoryOrder(db)

	if err = orderRepository.CreateNewOrder(order); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	productRepository := repositories.NewRepositoryProduct(db)

	for _, orderProducts := range order.OrderItems {
		if err = orderRepository.InsertOrderProducts(order.OrderID, orderProducts); err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
			return
		}

		if orderProducts.Product.Stock == orderProducts.Amount {
			if err = UpdateProductAvailabilityStripe(orderProducts, false); err != nil {
				responses.Erro(w, http.StatusInternalServerError, err)
				return
			}
		}

		if err = productRepository.DecrementProductStock(orderProducts); err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
			return
		}
	}

	responses.JSON(w, http.StatusOK, "Order Placed")
}

func ShowOrder(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	role, err := authentication.ExtractRoleFromToken(r)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	params := mux.Vars(r)
	orderID, err := strconv.ParseUint(params["orderID"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectWithDatabase()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	orderRepository := repositories.NewRepositoryOrder(db)

	order, err := orderRepository.SearchOrderByID(orderID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if order.UserID != userID && role != "admin" {
		responses.Erro(w, http.StatusForbidden, errors.New("can't see this order"))
		return
	}

	err = orderRepository.SearchOrderItens(&order)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, order)
}

func ShowAllOrders(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectWithDatabase()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	orderRepository := repositories.NewRepositoryOrder(db)

	orders, err := orderRepository.ShowOrders()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	for i := 0; i < len(orders); i++ {
		err := orderRepository.SearchOrderItens(&orders[i])
		if err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
			return
		}
	}

	responses.JSON(w, http.StatusOK, orders)
}

func ShowUserOrders(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectWithDatabase()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	orderRepository := repositories.NewRepositoryOrder(db)

	orders, err := orderRepository.SearchOrderByUserID(userID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	for i := 0; i < len(orders); i++ {
		err := orderRepository.SearchOrderItens(&orders[i])
		if err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
			return
		}
	}

	responses.JSON(w, http.StatusOK, orders)
}

func ChangeOrderStatus(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	var order models.Order

	if err = json.Unmarshal(bodyRequest, &order); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectWithDatabase()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	orderRepository := repositories.NewRepositoryOrder(db)

	if err = orderRepository.ChangeStatus(order); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, "Updated")
}
