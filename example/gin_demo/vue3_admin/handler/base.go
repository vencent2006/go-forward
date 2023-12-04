package handler

const (
	CODE_SUCC = 200
	BOOL_SUCC = true
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}
