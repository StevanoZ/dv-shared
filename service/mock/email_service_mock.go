// Code generated by MockGen. DO NOT EDIT.
// Source: service/email_service.go

// Package shrd_mock_svc is a generated GoMock package.
package shrd_mock_svc

import (
	context "context"
	reflect "reflect"

	message "github.com/StevanoZ/dv-shared/message"
	gomock "github.com/golang/mock/gomock"
	rest "github.com/sendgrid/rest"
	mail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

// MockEmailClient is a mock of EmailClient interface.
type MockEmailClient struct {
	ctrl     *gomock.Controller
	recorder *MockEmailClientMockRecorder
}

// MockEmailClientMockRecorder is the mock recorder for MockEmailClient.
type MockEmailClientMockRecorder struct {
	mock *MockEmailClient
}

// NewMockEmailClient creates a new mock instance.
func NewMockEmailClient(ctrl *gomock.Controller) *MockEmailClient {
	mock := &MockEmailClient{ctrl: ctrl}
	mock.recorder = &MockEmailClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailClient) EXPECT() *MockEmailClientMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockEmailClient) Send(email *mail.SGMailV3) (*rest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", email)
	ret0, _ := ret[0].(*rest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockEmailClientMockRecorder) Send(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockEmailClient)(nil).Send), email)
}

// SendWithContext mocks base method.
func (m *MockEmailClient) SendWithContext(ctx context.Context, email *mail.SGMailV3) (*rest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendWithContext", ctx, email)
	ret0, _ := ret[0].(*rest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendWithContext indicates an expected call of SendWithContext.
func (mr *MockEmailClientMockRecorder) SendWithContext(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendWithContext", reflect.TypeOf((*MockEmailClient)(nil).SendWithContext), ctx, email)
}

// MockEmailSvc is a mock of EmailSvc interface.
type MockEmailSvc struct {
	ctrl     *gomock.Controller
	recorder *MockEmailSvcMockRecorder
}

// MockEmailSvcMockRecorder is the mock recorder for MockEmailSvc.
type MockEmailSvcMockRecorder struct {
	mock *MockEmailSvc
}

// NewMockEmailSvc creates a new mock instance.
func NewMockEmailSvc(ctrl *gomock.Controller) *MockEmailSvc {
	mock := &MockEmailSvc{ctrl: ctrl}
	mock.recorder = &MockEmailSvcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailSvc) EXPECT() *MockEmailSvcMockRecorder {
	return m.recorder
}

// SendVerifyOtp mocks base method.
func (m *MockEmailSvc) SendVerifyOtp(ctx context.Context, data message.OtpPayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendVerifyOtp", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendVerifyOtp indicates an expected call of SendVerifyOtp.
func (mr *MockEmailSvcMockRecorder) SendVerifyOtp(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendVerifyOtp", reflect.TypeOf((*MockEmailSvc)(nil).SendVerifyOtp), ctx, data)
}
