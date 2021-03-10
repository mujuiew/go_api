package rest

// Rate ...
type Rate struct {
	Rate          string  `json:"rate"`
	InterestRate  float64 `json:"interest_rate"`
	PromotionName string  `json:"promotion_name"`
}

// Output ...
type Output struct {
	Promotionname string  `json:"promotion name"`
	Interestrate  float64 `json:"interest rate"`
	Number        float64 `json:"account number"`
	Pmt           float64 `json:"Payment Amount with Nember of payment"`
}

// Promotion ...
type Promotion struct {
	PromotionName string `json:"promotion_name"`
}

// Input ...
type Input struct {
	DisbursementAmount float64 `json:"disbursement_amount"`
	NumberOfPayment    float64 `json:"number_of_payment"`
	CalDate            string  `json:"cal_date"`
	PaymentFrequency   float64 `json:"payment_frequency"`
	PaymentUnit        string  `json:"payment_unit"`
	AccountNumber      float64 `json:"account_number"`
}
