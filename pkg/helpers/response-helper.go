package helpers

type ResponseHelper struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
	Trace   any    `json:"trace,omitempty"`
}

func NewResponseHelper(code int, message string, data any, meta any, trace any) *ResponseHelper {
	return &ResponseHelper{
		Code:    code,
		Message: message,
		Data:    data,
		Meta:    meta,
		Trace:   trace,
	}
}
