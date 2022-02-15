package common

type ContextKey string

func (c ContextKey) String() string {
	return "etheralley context key " + string(c)
}

var (
	ContextKeyRequestId   = ContextKey("request id")
	ContextKeyAddress     = ContextKey("address")
	ContextKeyContract    = ContextKey("contract")
	ContextKeyTransaction = ContextKey("transaction")
)
