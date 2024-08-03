package tests_test

import (
	"context"
	"testing"

	"github.com/emitra-labs/pb/mail"
	"github.com/stretchr/testify/assert"
)

func TestSendTransactional_Success(t *testing.T) {
	res, err := mailClient.SendTransactional(context.Background(), &mail.SendTransactionalRequest{
		From:    "Example <no-reply@example.com>",
		To:      "John Doe <john@example.com>",
		Subject: "This is a test",
	})

	assert.NotEmpty(t, res)
	assert.Empty(t, err)

	if res != nil {
		assert.True(t, res.Success)
	}
}
