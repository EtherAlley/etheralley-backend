package common

type ContextKey string

const contextKeyPrefix = "etheralley context key "

func (c ContextKey) String() string {
	return contextKeyPrefix + string(c)
}

var (
	ContextKeyRequestId        = ContextKey("request id")
	ContextKeyAddress          = ContextKey("address")
	ContextKeyContract         = ContextKey("contract")
	ContextKeyTransaction      = ContextKey("transaction")
	ContextKeyRequestStartTime = ContextKey("request start time")
)
