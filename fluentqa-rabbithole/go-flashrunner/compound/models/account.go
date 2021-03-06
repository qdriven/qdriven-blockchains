package models

import "errors"

type Account struct {
	Address      string      `json:"address"`
	BlockUpdated interface{} `json:"block_updated"`
	Health       struct {
		Value string `json:"value"`
	} `json:"health"`
	Tokens                []Token `json:"tokens"`
	TotalBorrowValueInEth struct {
		Value string `json:"value"`
	} `json:"total_borrow_value_in_eth"`
	TotalCollateralValueInEth struct {
		Value string `json:"value"`
	} `json:"total_collateral_value_in_eth"`
}

// AccountResponse is a response to a
// https://api.compound.finance/api/v2/account?addresses[]= call
type AccountResponse struct {
	Accounts             []Account   `json:"accounts"`
	CloseFactor          float64     `json:"close_factor"`
	Error                interface{} `json:"error"`
	LiquidationIncentive float64     `json:"liquidation_incentive"`
	PaginationSummary    struct {
		PageNumber   int `json:"page_number"`
		PageSize     int `json:"page_size"`
		TotalEntries int `json:"total_entries"`
		TotalPages   int `json:"total_pages"`
	} `json:"pagination_summary"`
	Request struct {
		Addresses           []string    `json:"addresses"`
		BlockNumber         int         `json:"block_number"`
		BlockTimestamp      int         `json:"block_timestamp"`
		MaxHealth           interface{} `json:"max_health"`
		MinBorrowValueInEth interface{} `json:"min_borrow_value_in_eth"`
		PageNumber          int         `json:"page_number"`
		PageSize            int         `json:"page_size"`
	} `json:"request"`
}

type ValueField struct {
	Value string `json:"value"`
}

type AccountRequest struct {
	Addresses           []string `json:"addresses,omitempty"`
	BlockNumber         int      `json:"block_number,omitempty"`
	BlockTimestamp      int      `json:"block_timestamp,omitempty"`
	MaxHealth           ValueField   `json:"max_health"`
	MinBorrowValueInEth ValueField   `json:"min_borrow_value_in_eth"`
	PageNumber          int      `json:"page_number"`
	PageSize            int      `json:"page_size"`
}


// Token is an individual token
type Token struct {
	Address                 string `json:"address"`
	BorrowBalanceUnderlying struct {
		Value string `json:"value"`
	} `json:"borrow_balance_underlying"`
	LifetimeBorrowInterestAccrued struct {
		Value string `json:"value"`
	} `json:"lifetime_borrow_interest_accrued"`
	LifetimeSupplyInterestAccrued struct {
		Value string `json:"value"`
	} `json:"lifetime_supply_interest_accrued"`
	SupplyBalanceUnderlying struct {
		Value string `json:"value"`
	} `json:"supply_balance_underlying"`
	Symbol string `json:"symbol"`
}

// GetTokenByAddress is used to retrieve a token by its address from an AccountResponse type
func GetTokenByAddress(address string, resp *AccountResponse) (Token, error) {
	for _, token := range resp.Accounts[0].Tokens {
		if token.Address == address {
			return token, nil
		}
	}
	return Token{}, errors.New("token not found")
}
