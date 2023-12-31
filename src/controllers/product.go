package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	var product models.Product

	if err = json.Unmarshal(bodyRequest, &product); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = AddProductStripe(&product); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.ConnectWithDatabase()

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	productRepository := repositories.NewRepositoryProduct(db)

	imageID, err := productRepository.InsertImage(product.Image)

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	product.Image.ImageID = uint64(imageID)

	if err = productRepository.InsertNewProduct(product); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, "Inserted")
}

func ShowProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db, err := database.ConnectWithDatabase()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	productRepository := repositories.NewRepositoryProduct(db)

	product, err := productRepository.SearchProductByID(params["productID"])
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, product)
}

func ChangePrice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	var price *float64

	if err = json.Unmarshal(bodyRequest, &price); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = ChangePriceProductStripe(params["productID"], *price); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.ConnectWithDatabase()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	productRepository := repositories.NewRepositoryProduct(db)

	if err = productRepository.ChangeProductPrice(params["productID"], *price); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, "Price changed")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if err := DeleteProductStripe(params["productID"]); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.ConnectWithDatabase()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	productRepository := repositories.NewRepositoryProduct(db)

	if err = productRepository.Delete(params["productID"]); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, "Deleted")
}

func UpdateImage(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	var image models.Images
	if err = json.Unmarshal(bodyRequest, &image); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectWithDatabase()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	productRepository := repositories.NewRepositoryProduct(db)

	if err = productRepository.UpdateImage(image); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, "Updated")
}
