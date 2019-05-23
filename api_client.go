package eventcar

// ClientApi Api Event is define full message
type ClientApi interface {

	// On add event by name for message
	On(eventName string, fn func(ReqMsgContext)) error

	// FireByQueue
	FireByQueue(eventName string, msgBody []byte, callback ...Callback) error

	// FireByPublish
	FireByPublish(eventName string, msgBody []byte, callback ...Callback) error
}
