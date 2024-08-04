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
		Body: &mail.TransactionalBody{
			Name: "John Doe",
			Intros: []string{
				"We have received a request to reset your password. Click the button below to reset it.",
			},
			Actions: []*mail.TransactionalAction{
				{
					Link: "https://example.com",
					Text: "Reset password",
				},
			},
			Outros: []string{
				"If you did not request a password reset, please ignore this email.",
			},
		},
	})

	assert.NotEmpty(t, res)
	assert.Empty(t, err)

	if res != nil {
		assert.True(t, res.Success)
	}
}
