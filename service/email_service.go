package shrd_service

import (
	"context"
	"fmt"
	"strconv"

	shrd_model "github.com/StevanoZ/dv-shared/model"
	shrd_utils "github.com/StevanoZ/dv-shared/utils"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailClient interface {
	Send(email *mail.SGMailV3) (*rest.Response, error)
	SendWithContext(ctx context.Context, email *mail.SGMailV3) (*rest.Response, error)
}
type EmailSvc interface {
	SendVerifyOtp(ctx context.Context, data shrd_model.OtpData) error
}

type OtpData struct {
	Email   string
	OtpCode int
}

type EmailSvcImpl struct {
	client EmailClient
	config *shrd_utils.BaseConfig
}

func NewEmailSvc(
	client EmailClient,
	config *shrd_utils.BaseConfig,
) EmailSvc {
	return &EmailSvcImpl{client: client, config: config}
}

func (s *EmailSvcImpl) SendVerifyOtp(ctx context.Context, data shrd_model.OtpData) error {
	otpCodeStr := strconv.Itoa(data.OtpCode)
	from := mail.NewEmail("", s.config.SenderEmail)
	subject := "Verify OTP"
	to := mail.NewEmail("", data.Email)
	htmlContent := "<strong>Use this OTP CODE to verify your email (expired between 5 minutes) : </strong>" + "<strong>" + otpCodeStr + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)

	response, err := s.client.SendWithContext(ctx, message)
	if err != nil {
		fmt.Println("Error Send Email", err)
		return err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Headers)
	}

	return nil
}
