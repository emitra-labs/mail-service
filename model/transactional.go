package model

type TransactionalAction struct {
	Text  string `validate:"required"`
	Link  string `validate:"required"`
	Color string
}

type TransactionalBody struct {
	Name    string                 `validate:"required"`
	Intros  []string               `validate:"required"`
	Actions []*TransactionalAction `validate:"required"`
	Outros  []string
}

type SendTransactionalRequest struct {
	From    string             `validate:"required"`
	To      string             `validate:"required"`
	Subject string             `validate:"required"`
	Body    *TransactionalBody `validate:"required"`
}

type SendTransactionalResponse struct {
	Success bool
}
