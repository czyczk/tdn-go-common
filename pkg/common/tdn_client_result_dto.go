package common

type TdnClientResultDto struct {
	Data  map[string]interface{} `json:"data,omitempty"`
	Error *TdnErrorDto           `json:"error,omitempty"`
	IsOk  bool                   `json:"isOk"`
}

type TdnErrorDto struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
