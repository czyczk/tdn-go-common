package common

type TdnControllerView[T any] struct {
	Data  T             `json:"data,omitempty"`
	Error *TdnErrorView `json:"error,omitempty"`
	IsOk  bool          `json:"isOk"`
}

type TdnErrorView struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewTdnControllerViewOfOk[T any](data T) TdnControllerView[T] {
	view := TdnControllerView[T]{
		Data: data,
		IsOk: true,
	}
	return view
}

func NewTdnControllerViewOfError(err error) TdnControllerView[interface{}] {
	var errorView *TdnErrorView

	if tdnError, ok := err.(TdnError); ok {
		errorView = &TdnErrorView{
			Code:    tdnError.ErrorCode().Code(),
			Message: tdnError.Error(),
		}
	} else {
		errorView = &TdnErrorView{
			Code:    TdnErrorCode_DEFAULT.Code(),
			Message: err.Error(),
		}
	}

	return TdnControllerView[interface{}]{
		Error: errorView,
		IsOk:  false,
	}
}
