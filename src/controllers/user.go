package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectWithDatabase()

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.NewRepositoryUser(db)

	userInfo, err := userRepository.SearchUserByEmail(user.Email)

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CompareHashWithPassword(user.Password, userInfo.Password); err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(userInfo.ID, userInfo.RoleType)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, token)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)

	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("sign"); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = CreateUserStripe(&user); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if user.RoleType == "" {
		user.RoleType = "user"
	}

	db, err := database.ConnectWithDatabase()

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.NewRepositoryUser(db)

	if err = userRepository.InsertNewUser(user); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = userRepository.InsertDefaultAddress(user.ID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, "User created")
}

func ShowUserInfo(w http.ResponseWriter, r *http.Request) {
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

	userRepository := repositories.NewRepositoryUser(db)

	user, err := userRepository.SearchUserByID(userID)

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
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

	var password models.Password

	if err = json.Unmarshal(bodyRequest, &password); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectWithDatabase()

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.NewRepositoryUser(db)

	hash, err := userRepository.SearchPasswordFromUser(userID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
	}

	if err = security.CompareHashWithPassword(password.Current, hash); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = userRepository.UpdateUserPassword(userID, password); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, "Updated")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
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

	if err = DeleteUserStripe(userID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.NewRepositoryUser(db)

	if err = userRepository.Delete(userID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, "Deleted")
}

func UpdateUserAddress(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	var address models.Address
	address.UserID = userID

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = json.Unmarshal(bodyRequest, &address); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = ChangeAddressUserStripe(address); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.ConnectWithDatabase()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.NewRepositoryUser(db)

	if err = userRepository.UpdateUserAddress(address); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, "Updated")
}
