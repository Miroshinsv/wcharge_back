package entity

// Тут пока вопросы
type PaymentHistory struct {
	ID          int     `json:"id"`
	UserID      int     `json:"userId"`
	PowerbankID int     `json:"powerbankId"`
	Amount      float64 `json:"amount"`
}
