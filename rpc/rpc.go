package rpc

import (
	"context"

	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/mail-service/controller"
	"github.com/emitra-labs/mail-service/model"
	"github.com/emitra-labs/pb/mail"
)

type Server struct {
	mail.UnimplementedMailServer
}

func (s *Server) SendTransactional(ctx context.Context, req *mail.SendTransactionalRequest) (*mail.SendTransactionalResponse, error) {
	res, err := controller.SendTransactional(ctx, &model.SendTransactionalRequest{
		From:    req.From,
		To:      req.To,
		Subject: req.Subject,
	})
	if err != nil {
		return nil, errors.GRPCStatus(err)
	}

	return &mail.SendTransactionalResponse{
		Success: res.Success,
	}, nil
}
