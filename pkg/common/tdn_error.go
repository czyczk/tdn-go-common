package common

type TdnErrorCode interface {
	Code() string
	MessageTemplate() string
}

type TdnError interface {
	ErrorCode() TdnErrorCode
	Params() []string
	error
}

// Make `TdnErrorCode` instantiable in this package.
type defaultTdnErrorCode struct {
	code            string
	messageTemplate string
}

// Impl [TdnErrorCode] for [*defaultTdnErrorCode].
func (e *defaultTdnErrorCode) Code() string {
	return e.code
}

func (e *defaultTdnErrorCode) MessageTemplate() string {
	return e.messageTemplate
}

var (
	TdnErrorCode_DEFAULT = &defaultTdnErrorCode{
		code:            "DEFAULT",
		messageTemplate: "unknown error",
	}
)
