package model

type SendTransactionalRequest struct {
	From    string `validate:"required"`
	To      string `validate:"required"`
	Subject string `validate:"required"`
}

type SendTransactionalResponse struct {
	Success bool
}
