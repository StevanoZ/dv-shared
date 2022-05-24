// Code generated by MockGen. DO NOT EDIT.
// Source: pubsub/pubsub.go

// Package pubsub_client is a generated GoMock package.
package pubsub_client

import (
	context "context"
	reflect "reflect"

	pubsub "cloud.google.com/go/pubsub"
	gomock "github.com/golang/mock/gomock"
)

// MockGooglePubSub is a mock of GooglePubSub interface.
type MockGooglePubSub struct {
	ctrl     *gomock.Controller
	recorder *MockGooglePubSubMockRecorder
}

// MockGooglePubSubMockRecorder is the mock recorder for MockGooglePubSub.
type MockGooglePubSubMockRecorder struct {
	mock *MockGooglePubSub
}

// NewMockGooglePubSub creates a new mock instance.
func NewMockGooglePubSub(ctrl *gomock.Controller) *MockGooglePubSub {
	mock := &MockGooglePubSub{ctrl: ctrl}
	mock.recorder = &MockGooglePubSubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGooglePubSub) EXPECT() *MockGooglePubSubMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockGooglePubSub) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockGooglePubSubMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockGooglePubSub)(nil).Close))
}

// CreateSubscription mocks base method.
func (m *MockGooglePubSub) CreateSubscription(ctx context.Context, id string, cfg pubsub.SubscriptionConfig) (*pubsub.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscription", ctx, id, cfg)
	ret0, _ := ret[0].(*pubsub.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubscription indicates an expected call of CreateSubscription.
func (mr *MockGooglePubSubMockRecorder) CreateSubscription(ctx, id, cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscription", reflect.TypeOf((*MockGooglePubSub)(nil).CreateSubscription), ctx, id, cfg)
}

// CreateTopic mocks base method.
func (m *MockGooglePubSub) CreateTopic(ctx context.Context, topicID string) (*pubsub.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTopic", ctx, topicID)
	ret0, _ := ret[0].(*pubsub.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTopic indicates an expected call of CreateTopic.
func (mr *MockGooglePubSubMockRecorder) CreateTopic(ctx, topicID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTopic", reflect.TypeOf((*MockGooglePubSub)(nil).CreateTopic), ctx, topicID)
}

// Subscription mocks base method.
func (m *MockGooglePubSub) Subscription(id string) *pubsub.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscription", id)
	ret0, _ := ret[0].(*pubsub.Subscription)
	return ret0
}

// Subscription indicates an expected call of Subscription.
func (mr *MockGooglePubSubMockRecorder) Subscription(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscription", reflect.TypeOf((*MockGooglePubSub)(nil).Subscription), id)
}

// Topic mocks base method.
func (m *MockGooglePubSub) Topic(id string) *pubsub.Topic {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Topic", id)
	ret0, _ := ret[0].(*pubsub.Topic)
	return ret0
}

// Topic indicates an expected call of Topic.
func (mr *MockGooglePubSubMockRecorder) Topic(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Topic", reflect.TypeOf((*MockGooglePubSub)(nil).Topic), id)
}

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

// CreateSubscriptionIfNotExists mocks base method.
func (m *MockPubSubClient) CreateSubscriptionIfNotExists(ctx context.Context, client GooglePubSub, id string, topic *pubsub.Topic) (*pubsub.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscriptionIfNotExists", ctx, client, id, topic)
	ret0, _ := ret[0].(*pubsub.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubscriptionIfNotExists indicates an expected call of CreateSubscriptionIfNotExists.
func (mr *MockPubSubClientMockRecorder) CreateSubscriptionIfNotExists(ctx, client, id, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscriptionIfNotExists", reflect.TypeOf((*MockPubSubClient)(nil).CreateSubscriptionIfNotExists), ctx, client, id, topic)
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
func (m *MockPubSubClient) PullMessages(ctx context.Context, id string, topic *pubsub.Topic, callback func(context.Context, *pubsub.Message)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PullMessages", ctx, id, topic, callback)
}

// PullMessages indicates an expected call of PullMessages.
func (mr *MockPubSubClientMockRecorder) PullMessages(ctx, id, topic, callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullMessages", reflect.TypeOf((*MockPubSubClient)(nil).PullMessages), ctx, id, topic, callback)
}
