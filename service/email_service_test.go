package shrd_service

import (
	"context"
	"errors"
	"testing"

	"github.com/StevanoZ/dv-shared/message"
	mock_svc "github.com/StevanoZ/dv-shared/service/mock"
	shrd_utils "github.com/StevanoZ/dv-shared/utils"
	"github.com/golang/mock/gomock"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/stretchr/testify/assert"
)

func initEmailSvc(ctrl *gomock.Controller) (EmailSvc, *mock_svc.MockEmailClient) {
	config := shrd_utils.LoadBaseConfig("../app", "test")
	emailClient := mock_svc.NewMockEmailClient(ctrl)

	return NewEmailSvc(emailClient, config), emailClient
}

func TestNewSgClient(t *testing.T) {
	config := shrd_utils.LoadBaseConfig("../app", "test")
	assert.NotNil(t, NewSgClient(config))
}
func TestEmailSvc(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	emailSvc, client := initEmailSvc(ctrl)
	otpData := message.OtpPayload{
		Email:   "test@test.com",
		OtpCode: 227799,
	}
	t.Run("Success sent email", func(t *testing.T) {
		client.EXPECT().SendWithContext(ctx, gomock.AssignableToTypeOf(&mail.SGMailV3{})).
			Return(&rest.Response{}, nil).Times(1)

		err := emailSvc.SendVerifyOtp(ctx, otpData)
		assert.NoError(t, err)
	})

	t.Run("Failed sent email", func(t *testing.T) {
		client.EXPECT().SendWithContext(ctx, gomock.AssignableToTypeOf(&mail.SGMailV3{})).
			Return(nil, errors.New("failed when sending email")).Times(1)

		err := emailSvc.SendVerifyOtp(ctx, otpData)
		assert.Error(t, err)
		assert.Equal(t, "failed when sending email", err.Error())
	})
}
