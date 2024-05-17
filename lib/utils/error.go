package utils

import "food-order/lib/helper"

func errorType(err error) (int, error) {
	return helper.CommonError(err)
}
