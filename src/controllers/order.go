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

	repository := repositories.NewRepositoryOrder(db)

	if err = repository.CreateNewOrder(order); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	for _, orderProducts := range order.OrderItems {
		if err = repository.InsertOrderProducts(order.OrderID, orderProducts); err != nil {
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

	repository := repositories.NewRepositoryOrder(db)

	order, err := repository.SearchOrderByID(orderID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if order.UserID != userID {
		responses.Erro(w, http.StatusForbidden, errors.New("can't see this order"))
		return
	}

	order, err = repository.SearchOrderItens(order)
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

	repository := repositories.NewRepositoryOrder(db)

	orders, err := repository.ShowOrders()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	for i := 0; i < len(orders); i++ {
		orderProducts, err := repository.SearchOrderItens(orders[i])
		if err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
			return
		}

		orders[i].OrderItems = append(orders[i].OrderItems, orderProducts.OrderItems...)
	}

	responses.JSON(w, http.StatusOK, orders)
}
