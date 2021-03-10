package rest

// Rate ...
type Rate struct {
	Rate          string  `json:"rate"`
	InterestRate  float64 `json:"interest_rate"`
	PromotionName string  `json:"promotion_name"`
}

// Output ...
type Output struct {
	Promotionname string `json:"promotion name"`
	Number        int64  `json:"account number"`
}

// Promotion ...
type Promotion struct {
	PromotionName string `json:"promotion_name"`
}

// Input ...
type Input struct {
	CalDate            string  `json:"cal_date"`
	AccountNumber      int64   `json:"account_number"`
}
