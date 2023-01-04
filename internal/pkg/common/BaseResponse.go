package common

type BaseResponse struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
	ErrDetails error  `json:"errDetails"`
}

func NewSuccessResponse() BaseResponse {
	var response BaseResponse
	response.Code = 0
	response.Message = "success"
	return response
}

func NewErrResponse(code int, message string, err error) BaseResponse {
	var response BaseResponse
	response.Code = code
	response.Message = message
	if response.Message == "<empty>" {
		response.Message = "fail"
	}
	response.ErrDetails = err
	return response
}
