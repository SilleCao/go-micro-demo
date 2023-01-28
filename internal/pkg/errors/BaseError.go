package errors

import "net/http"

type BaseError struct {
	ErrMsg     string `json:"errMsg"`
	ErrCode    int    `json:"errCode"`
	StatusCode int    `json:"-"`
}

func (err *BaseError) Error() string {
	return err.ErrMsg
}

func newError(errMsg string, errCode int, statusCode int) error {
	return &BaseError{errMsg, errCode, statusCode}
}

//400
func NewBadRequestErr(errMsg string, errCode int) error {
	return newError(errMsg, errCode, http.StatusBadRequest)
}

//401
func NewUnauthorizedErr(errMsg string, errCode int) error {
	return newError(errMsg, errCode, http.StatusUnauthorized)
}

//403
func NewForbiddenErr(errMsg string, errCode int) error {
	return newError(errMsg, errCode, http.StatusForbidden)
}

//404
func NewNotFoundErr(errMsg string, errCode int) error {
	return newError(errMsg, errCode, http.StatusNotFound)
}

//500
func NewInteralServerErr(errMsg string, errCode int) error {
	return newError(errMsg, errCode, http.StatusInternalServerError)
}

//503
func NewServiceUnavailableErr(errMsg string, errCode int) error {
	return newError(errMsg, errCode, http.StatusServiceUnavailable)
}

func HandleErr(err error) BaseError {
	switch err := err.(type) {
	case *BaseError:
		return *err
	default:
		return BaseError{err.Error(), 1, http.StatusBadRequest}
	}
}
