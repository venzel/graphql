package model

type Transaction struct {
	ID        string  `json:"id"`
	Amount    float64 `json:"amount"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}
