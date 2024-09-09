package types

type Expense struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Amount      float32 `json:"amount"`
}

func (e *Expense) AddID(id int) {
	e.ID = id
}
