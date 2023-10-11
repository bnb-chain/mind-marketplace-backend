package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/bnb-chain/greenfield-data-marketplace-backend/models"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/restapi/operations/item"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/service"
)

func HandleGetItem() func(params item.GetItemParams) middleware.Responder {
	return func(params item.GetItemParams) middleware.Responder {
		context := params.HTTPRequest.Context()
		response, err := service.ItemSvc.Get(context, params.ItemID)
		code, message := Error(err)
		payload := models.ItemResponse{
			Code:    code,
			Message: message}
		if err == nil {
			payload.Data = &models.ItemResponseData{
				Item: response,
			}
		}
		return item.NewGetItemOK().WithPayload(&payload)
	}
}

func HandleGetItemByGroup() func(params item.GetItemByGroupParams) middleware.Responder {
	return func(params item.GetItemByGroupParams) middleware.Responder {
		context := params.HTTPRequest.Context()
		response, err := service.ItemSvc.GetByGroup(context, params.GroupID)
		code, message := Error(err)
		payload := models.ItemResponse{
			Code:    code,
			Message: message}
		if err == nil {
			payload.Data = &models.ItemResponseData{
				Item: response,
			}
		}
		return item.NewGetItemOK().WithPayload(&payload)
	}
}

func HandleSearchItem() func(request item.SearchItemParams) middleware.Responder {
	return func(params item.SearchItemParams) middleware.Responder {
		context := params.HTTPRequest.Context()
		total, response, err := service.ItemSvc.Search(context, params.Body)
		code, message := Error(err)
		payload := models.PageItemResponse{
			Code:    code,
			Message: message}
		if err == nil {
			payload.Data = &models.PageItemResponseData{
				Items: response,
				Total: total,
			}
		}
		return item.NewSearchItemOK().WithPayload(&payload)
	}
}
