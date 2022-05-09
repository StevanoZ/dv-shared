package shrd_service

type MessageBrokerClient interface {
	SendEvents(topic []string, key string, message interface{}) error
	ListenEvent(topic string, cb func(payload any, errMsg error, close func())) error
}
