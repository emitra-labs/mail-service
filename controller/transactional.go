package controller

import (
	"context"

	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/mail-service/model"
)

func SendTransactional(ctx context.Context, req *model.SendTransactionalRequest) (*model.SendTransactionalResponse, error) {
	return nil, errors.Internal("Not implemented")
}
