package utils

import "fmt"

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (cu *CustomError) Error() string {
	return fmt.Sprintf("error: %s code: %d", cu.Message, cu.Code)
}
