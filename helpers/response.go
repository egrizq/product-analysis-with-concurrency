package helpers

import "product-store-management/model"

func Response(payload string, status int, message string) model.Response {
	return model.Response{
		Payload:    payload,
		StatusCode: status,
		Message:    message,
	}
}
