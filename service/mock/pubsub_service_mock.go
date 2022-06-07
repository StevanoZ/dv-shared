// Code generated by MockGen. DO NOT EDIT.
// Source: service/pubsub_service.go

// Package shrd_mock_svc is a generated GoMock package.
package shrd_mock_svc

import (
	context "context"
	reflect "reflect"

	pubsub "cloud.google.com/go/pubsub"
	gomock "github.com/golang/mock/gomock"
)

// MockPubSubClient is a mock of PubSubClient interface.
type MockPubSubClient struct {
	ctrl     *gomock.Controller
	recorder *MockPubSubClientMockRecorder
}

// MockPubSubClientMockRecorder is the mock recorder for MockPubSubClient.
type MockPubSubClientMockRecorder struct {
	mock *MockPubSubClient
}

// NewMockPubSubClient creates a new mock instance.
func NewMockPubSubClient(ctrl *gomock.Controller) *MockPubSubClient {
	mock := &MockPubSubClient{ctrl: ctrl}
	mock.recorder = &MockPubSubClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPubSubClient) EXPECT() *MockPubSubClientMockRecorder {
	return m.recorder
}

// CheckTopicAndPublish mocks base method.
func (m *MockPubSubClient) CheckTopicAndPublish(ctx context.Context, topicsName []string, orderingKey string, data any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CheckTopicAndPublish", ctx, topicsName, orderingKey, data)
}

// CheckTopicAndPublish indicates an expected call of CheckTopicAndPublish.
func (mr *MockPubSubClientMockRecorder) CheckTopicAndPublish(ctx, topicsName, orderingKey, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTopicAndPublish", reflect.TypeOf((*MockPubSubClient)(nil).CheckTopicAndPublish), ctx, topicsName, orderingKey, data)
}

// Close mocks base method.
func (m *MockPubSubClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockPubSubClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPubSubClient)(nil).Close))
}

// CreateSubscriptionIfNotExists mocks base method.
func (m *MockPubSubClient) CreateSubscriptionIfNotExists(ctx context.Context, id string, topic *pubsub.Topic) (*pubsub.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscriptionIfNotExists", ctx, id, topic)
	ret0, _ := ret[0].(*pubsub.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubscriptionIfNotExists indicates an expected call of CreateSubscriptionIfNotExists.
func (mr *MockPubSubClientMockRecorder) CreateSubscriptionIfNotExists(ctx, id, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscriptionIfNotExists", reflect.TypeOf((*MockPubSubClient)(nil).CreateSubscriptionIfNotExists), ctx, id, topic)
}

// CreateTopicIfNotExists mocks base method.
func (m *MockPubSubClient) CreateTopicIfNotExists(ctx context.Context, topicName string) (*pubsub.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTopicIfNotExists", ctx, topicName)
	ret0, _ := ret[0].(*pubsub.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTopicIfNotExists indicates an expected call of CreateTopicIfNotExists.
func (mr *MockPubSubClientMockRecorder) CreateTopicIfNotExists(ctx, topicName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTopicIfNotExists", reflect.TypeOf((*MockPubSubClient)(nil).CreateTopicIfNotExists), ctx, topicName)
}

// PublishTopics mocks base method.
func (m *MockPubSubClient) PublishTopics(ctx context.Context, topics []*pubsub.Topic, data any, orderingKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishTopics", ctx, topics, data, orderingKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishTopics indicates an expected call of PublishTopics.
func (mr *MockPubSubClientMockRecorder) PublishTopics(ctx, topics, data, orderingKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishTopics", reflect.TypeOf((*MockPubSubClient)(nil).PublishTopics), ctx, topics, data, orderingKey)
}

// PullMessages mocks base method.
func (m *MockPubSubClient) PullMessages(ctx context.Context, id string, topic *pubsub.Topic, callback func(context.Context, *pubsub.Message)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullMessages", ctx, id, topic, callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// PullMessages indicates an expected call of PullMessages.
func (mr *MockPubSubClientMockRecorder) PullMessages(ctx, id, topic, callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullMessages", reflect.TypeOf((*MockPubSubClient)(nil).PullMessages), ctx, id, topic, callback)
}
