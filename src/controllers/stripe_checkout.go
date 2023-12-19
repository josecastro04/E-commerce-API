package controllers

import (
	"api/src/models"
	"fmt"
	"github.com/stripe/stripe-go/v76"
	session2 "github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/price"
	"time"
)

func CreateNewSessionCheckOut(order *models.Order) (string, error) {
	searchCustomerParams := &stripe.CustomerSearchParams{
		SearchParams: stripe.SearchParams{
			Query: fmt.Sprintf("id:'%s'", order.UserID),
		},
	}

	result := customer.Search(searchCustomerParams)

	var userEmail string
	for result.Next() {
		userEmail = result.Customer().Email
	}

	var items []*stripe.CheckoutSessionLineItemParams
	for _, orderitem := range order.OrderItems {
		var item *stripe.CheckoutSessionLineItemParams

		searchPriceParams := &stripe.PriceSearchParams{
			SearchParams: stripe.SearchParams{
				Query: fmt.Sprintf("active:'true' AND product:'%s'", orderitem.Product.ID),
			},
		}

		searchResult := price.Search(searchPriceParams)

		var currentPrice *stripe.Price
		for searchResult.Next() {
			currentPrice = searchResult.Price()
		}

		item = &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(currentPrice.ID),
			Quantity: stripe.Int64(int64(orderitem.Amount)),
		}

		items = append(items, item)
	}

	sessionParams := &stripe.CheckoutSessionParams{
		CustomerEmail:      stripe.String(userEmail),
		SuccessURL:         stripe.String("https://example.com/sucess"),
		LineItems:          items,
		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
		PaymentMethodTypes: stripe.StringSlice([]string{string(stripe.PaymentMethodTypePaypal), string(stripe.PaymentMethodTypeCard), string(stripe.PaymentMethodCardWalletTypeApplePay)}),
		PaymentMethodOptions: &stripe.CheckoutSessionPaymentMethodOptionsParams{
			Paypal: &stripe.CheckoutSessionPaymentMethodOptionsPaypalParams{
				CaptureMethod:   stripe.String("manual"),
				PreferredLocale: stripe.String("en-GB"),
			},
			Card: &stripe.CheckoutSessionPaymentMethodOptionsCardParams{
				Installments: &stripe.CheckoutSessionPaymentMethodOptionsCardInstallmentsParams{
					Enabled: stripe.Bool(false),
				},
			},
		},
		ExpiresAt: stripe.Int64(time.Now().Add(30 * time.Minute).Unix()),
	}

	checkout, err := session2.New(sessionParams)

	if err != nil {
		return "", err
	}

	order.OrderID = checkout.ID
	return checkout.URL, nil
}
