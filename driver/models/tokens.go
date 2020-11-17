package models


// PlaidIntegration Table that stores plaid access info needed for requests to linked bank accounts
type PlaidIntegration struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	UserID        string `json:"userid"`
	ItemID      string `json:"itemid"`
	AccessToken string `json:"accesstoken"`
	PaymentID   string `json:"paymentid"`
}

// CreatePlaidIntegration create new entries in table
type CreatePlaidIntegration struct {
	UserID        string `json:"userid" binding:"required"`
	ItemID      string `json:"itemid" binding:"required"`
	AccessToken string `json:"accesstoken" binding:"required"`
	PaymentID   string `json:"paymentid" binding:"required"`
}