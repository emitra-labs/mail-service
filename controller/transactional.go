package controller

import (
	"context"
	"os"

	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/common/log"
	"github.com/emitra-labs/common/validator"
	"github.com/emitra-labs/mail-service/model"
	"github.com/emitra-labs/mail-service/smtp"
	"github.com/go-hermes/hermes/v2"
	lo "github.com/samber/lo/parallel"
	mail "github.com/xhit/go-simple-mail/v2"
)

func SendTransactional(ctx context.Context, req *model.SendTransactionalRequest) (*model.SendTransactionalResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	h := hermes.Hermes{
		Theme: new(hermes.Default),
		Product: hermes.Product{
			Name:      os.Getenv("PRODUCT_NAME"),
			Link:      os.Getenv("PRODUCT_LINK"),
			Logo:      os.Getenv("PRODUCT_LOGO"),
			Copyright: os.Getenv("PRODUCT_COPYRIGHT"),
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name:   req.Body.Name,
			Intros: req.Body.Intros,
			Actions: lo.Map(req.Body.Actions, func(a *model.TransactionalAction, i int) hermes.Action {
				return hermes.Action{
					Button: hermes.Button{
						Color: a.Color,
						Link:  a.Link,
						Text:  a.Text,
					},
				}
			}),
			Outros: req.Body.Outros,
		},
	}

	body, err := h.GenerateHTML(email)
	if err != nil {
		log.Errorf("Failed to generate email html: %s", err)
		return nil, errors.Internal()
	}

	msg := mail.NewMSG()
	msg.SetFrom(req.From)
	msg.AddTo(req.To)
	msg.SetSubject(req.Subject)
	msg.SetBody(mail.TextHTML, body)

	if err := msg.Send(smtp.Client); err != nil {
		log.Errorf("Failed to send email: %s", err)
		return nil, errors.Internal()
	}

	return &model.SendTransactionalResponse{
		Success: true,
	}, nil
}
