package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/bnb-chain/mind-marketplace-backend/models"
	"github.com/bnb-chain/mind-marketplace-backend/restapi/operations/account"
	"github.com/bnb-chain/mind-marketplace-backend/service"
)

func HandleGetAccount() func(params account.GetAccountParams) middleware.Responder {
	return func(params account.GetAccountParams) middleware.Responder {
		context := params.HTTPRequest.Context()
		response, err := service.AccountSvc.Get(context, params.Address)
		code, message := Error(err)
		payload := models.AccountResponse{
			Code:    code,
			Message: message}
		if err == nil {
			payload.Data = &models.AccountResponseData{Account: response}
		}
		return account.NewGetAccountOK().WithPayload(&payload)
	}
}

/*
func HandleUpdateAccount() func(request account.UpdateAccountParams) middleware.Responder {
	return func(params account.UpdateAccountParams) middleware.Responder {
		context := params.HTTPRequest.Context().Value(ctx.ContextKey).(ctx.CTX)
		response, err := service.AccountSvc.Update(context, params.Body)
		code, message := Error(err)
		payload := models.AccountResponse{
			Code:    code,
			Message: message}
		if err == nil {
			payload.Data = &models.AccountResponseData{Account: response}
		}
		return account.NewUpdateAccountOK().WithPayload(&payload)
	}
}
*/
