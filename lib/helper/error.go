package helper

import (
	"fmt"
	"food-order/lib/constant"
	"net/http"
	"strings"
)

var commonErrorMap = map[error]int{
	constant.ErrNotFound:   http.StatusNotFound,
	constant.ErrConflict:   http.StatusConflict,
	constant.ErrBadRequest: http.StatusBadRequest,
}

// CommonError is
func CommonError(err error) (int, error) {

	if strings.Contains(err.Error(), "required") {
		return http.StatusBadRequest, fmt.Errorf(err.Error())
	}

	switch err {
	case constant.ErrNotFound:
		return commonErrorMap[constant.ErrNotFound], constant.ErrNotFound
	case constant.ErrConflict:
		return commonErrorMap[constant.ErrConflict], constant.ErrConflict
	case constant.ErrBadRequest:
		return commonErrorMap[constant.ErrBadRequest], constant.ErrBadRequest
	}
	return http.StatusInternalServerError, fmt.Errorf(err.Error())
}
