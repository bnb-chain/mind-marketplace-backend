package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/bnb-chain/greenfield-data-marketplace-backend/models"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/restapi/operations/purchase"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/service"
)

func HandleGetPurchase() func(params purchase.GetPurchaseParams) middleware.Responder {
	return func(params purchase.GetPurchaseParams) middleware.Responder {
		context := params.HTTPRequest.Context()
		response, err := service.PurchaseSvc.Get(context, params.PurchaseID)
		code, message := Error(err)
		payload := models.PurchaseResponse{
			Code:    code,
			Message: message}
		if err == nil {
			payload.Data = &models.PurchaseResponseData{
				Purchase: response,
			}
		}
		return purchase.NewGetPurchaseOK().WithPayload(&payload)
	}
}

func HandleSearchPurchase() func(request purchase.SearchPurchaseParams) middleware.Responder {
	return func(params purchase.SearchPurchaseParams) middleware.Responder {
		context := params.HTTPRequest.Context()
		total, response, err := service.PurchaseSvc.Search(context, params.Body)
		code, message := Error(err)
		payload := models.PagePurchaseResponse{
			Code:    code,
			Message: message}
		if err == nil {
			payload.Data = &models.PagePurchaseResponseData{
				Purchases: response,
				Total:     total,
			}
		}
		return purchase.NewSearchPurchaseOK().WithPayload(&payload)
	}
}
