package controllers

import (
	"api/src/models"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/product"
)

func AddProductStripe(productInfo *models.Product) error {

	productParams := &stripe.ProductParams{
		Name:        stripe.String(productInfo.Name),
		Description: stripe.String(productInfo.Description),
		DefaultPriceData: &stripe.ProductDefaultPriceDataParams{
			Currency:   stripe.String(string(stripe.CurrencyEUR)),
			UnitAmount: stripe.Int64(int64(productInfo.Price * 100)),
		},
	}

	createdProduct, err := product.New(productParams)
	if err != nil {
		return err
	}

	productInfo.ID = createdProduct.ID

	return nil
}

func UpdateProductAvailabilityStripe(orderItem models.OrderItem, value bool) error {
	updateParams := &stripe.ProductParams{Active: stripe.Bool(value)}
	if _, err := product.Update(orderItem.Product.ID, updateParams); err != nil {
		return err
	}
	return nil
}

func DeleteProductStripe(productID string) error {
	deleteParams := &stripe.ProductParams{}
	if _, err := product.Del(productID, deleteParams); err != nil {
		return err
	}

	return nil
}

func ChangePriceProductStripe(productID string, price float64) error {
	updateParams := &stripe.ProductParams{
		DefaultPriceData: &stripe.ProductDefaultPriceDataParams{
			UnitAmount: stripe.Int64(int64(price * 100)),
		},
	}

	if _, err := product.Update(productID, updateParams); err != nil {
		return err
	}
	return nil
}
