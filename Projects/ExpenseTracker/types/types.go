package types

type Expense struct {
	ID          int     `json:"id"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Month       int     `json:"month"`
}

func (e *Expense) AddID(id int) {
	e.ID = id
}
