package controllers

import (
	"api/src/models"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/customer"
)

func CreateUserStripe(user *models.User) error {
	userParams := &stripe.CustomerParams{
		Name:  stripe.String(user.Name),
		Email: stripe.String(user.Email),
		Phone: stripe.String(user.Phone),
	}

	createdUser, err := customer.New(userParams)

	if err != nil {
		return err
	}

	user.ID = createdUser.ID
	return nil
}

func ChangeAddressUserStripe(address models.Address) error {
	addressParams := &stripe.CustomerParams{
		Address: &stripe.AddressParams{
			City:       stripe.String(address.City),
			Country:    stripe.String(address.Country),
			Line1:      stripe.String(address.Address),
			PostalCode: stripe.String(address.Postal_Code),
			State:      stripe.String(address.State),
		},
	}

	if _, err := customer.Update(address.UserID, addressParams); err != nil {
		return err
	}

	return nil
}

func DeleteUserStripe(userID string) error {
	userParams := &stripe.CustomerParams{}

	if _, err := customer.Del(userID, userParams); err != nil {
		return err
	}

	return nil
}
