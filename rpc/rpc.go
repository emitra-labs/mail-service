package rpc

import (
	"context"

	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/mail-service/controller"
	"github.com/emitra-labs/mail-service/model"
	"github.com/emitra-labs/pb/mail"
	"github.com/samber/lo"
)

type Server struct {
	mail.UnimplementedMailServer
}

func (s *Server) SendTransactional(ctx context.Context, req *mail.SendTransactionalRequest) (*mail.SendTransactionalResponse, error) {
	res, err := controller.SendTransactional(ctx, &model.SendTransactionalRequest{
		From:    req.From,
		To:      req.To,
		Subject: req.Subject,
		Body: &model.TransactionalBody{
			Name:   req.Body.Name,
			Intros: req.Body.Intros,
			Actions: lo.Map(req.Body.Actions, func(a *mail.TransactionalAction, i int) *model.TransactionalAction {
				return &model.TransactionalAction{
					Color: a.Color,
					Link:  a.Link,
					Text:  a.Text,
				}
			}),
			Outros: req.Body.Outros,
		},
	})
	if err != nil {
		return nil, errors.ToGRPCStatus(err)
	}

	return &mail.SendTransactionalResponse{
		Success: res.Success,
	}, nil
}
