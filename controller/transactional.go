package controller

import (
	"context"

	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/common/log"
	"github.com/emitra-labs/common/validator"
	"github.com/emitra-labs/mail-service/model"
	"github.com/emitra-labs/mail-service/smtp"

	mail "github.com/xhit/go-simple-mail/v2"
)

func SendTransactional(ctx context.Context, req *model.SendTransactionalRequest) (*model.SendTransactionalResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	email := mail.NewMSG()
	email.SetFrom(req.From)
	email.AddTo(req.To)
	email.SetSubject(req.Subject)
	email.SetBody(mail.TextHTML, "Hello, World!")

	if err := email.Send(smtp.Client); err != nil {
		log.Errorf("Failed to send email: %s", err)
		return nil, errors.Internal()
	}

	return &model.SendTransactionalResponse{
		Success: true,
	}, nil
}
