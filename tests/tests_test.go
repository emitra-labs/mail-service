package tests_test

import (
	"os"
	"testing"

	"github.com/emitra-labs/common/grpc/testkit"
	"github.com/emitra-labs/mail-service/rpc"
	"github.com/emitra-labs/mail-service/smtp"
	pb "github.com/emitra-labs/pb/mail"
	"google.golang.org/grpc"
)

var mailClientConn *grpc.ClientConn
var mailClient pb.MailClient
var closeMailServer func()

func TestMain(m *testing.M) {
	setup()
	defer teardown()
	os.Exit(m.Run())
}

func setup() {
	// Open smtp connection
	smtp.Open(os.Getenv("SMTP_URL"))

	// Setup mail gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterMailServer(grpcServer, &rpc.Server{})
	mailClientConn, closeMailServer = testkit.NewClientConn(grpcServer)
	mailClient = pb.NewMailClient(mailClientConn)
}

func teardown() {
	closeMailServer()
	smtp.Close()
}
