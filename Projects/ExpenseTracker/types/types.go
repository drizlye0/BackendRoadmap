package types

type Expense struct {
	ID          int     `json:"id"`
	Date        string  `json:"time"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

func (e *Expense) AddID(id int) {
	e.ID = id
}
