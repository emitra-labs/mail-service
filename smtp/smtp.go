package smtp

import (
	"net/url"
	"strconv"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

var Client *mail.SMTPClient

func Open(smtpURL string) {
	parsed, err := url.Parse(smtpURL)
	if err != nil {
		panic(err)
	}

	smtpClient := mail.NewSMTPClient()

	if parsed.Hostname() != "" {
		smtpClient.Host = parsed.Hostname()
	}

	if parsed.Port() != "" {
		smtpClient.Port, _ = strconv.Atoi(parsed.Port())
	}

	if parsed.User.Username() != "" {
		smtpClient.Username = parsed.User.Username()
	}

	if _, ok := parsed.User.Password(); ok {
		smtpClient.Password, _ = parsed.User.Password()
	}

	if parsed.Scheme == "smtps" {
		smtpClient.Encryption = mail.EncryptionSTARTTLS
	}

	smtpClient.KeepAlive = true

	Client, err = smtpClient.Connect()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			Client.Noop()
			time.Sleep(30 * time.Second)
		}
	}()
}

func Close() error {
	return Client.Close()
}
