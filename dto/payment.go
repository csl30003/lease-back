package dto

type Payment struct {
	ID              int     `json:"id"`
	CreatedAt       string  `json:"created_at"`
	Type            int     `json:"type"`
	Money           float64 `json:"money"`
	UserID          int     `json:"user_id"`
	OrderID         int     `json:"order_id"`
	OrderIdentifier string  `json:"order_identifier"`
}
