package utils

import (
	"strings"
	"model"
	"log"
)

const (
	validHeader = "application/json;v=1"
)

func ValidateHeaders(headers map[string]string) model.genericError {
	if (!validateAccept(headers["Accept"])) {
		return model.genericError {
			StatusCode: 406
			Body: {
				ErrorCode: "200111"
				ErrorMsg: "Method Not Acceptable"
				DeveloperCode: "400111"
				DeveloperMsg: "Method Not Supported for " + headers["Accept"]
			}
		}
	}
	return nil
}

func validateAccept(accept string) bool {
	return strings.Compare(accept, validHeader) == 0;
}