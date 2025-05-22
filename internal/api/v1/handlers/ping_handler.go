package handlers

import (
	"github.com/KasperSaaby/calculatron-service/generated/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func GetPingHandler() operations.GetPingHandler {
	return operations.GetPingHandlerFunc(func(params operations.GetPingParams) middleware.Responder {
		return operations.NewGetPingOK().WithPayload("pong")
	})
}
