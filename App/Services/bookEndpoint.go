package Services

import (
	"context"
	"web-scene/App"
	"web-scene/models"
)

// BookListRequest 图书列表请求
type BookListRequest struct {
	Size int `form:"size"`
}

// BookListResponse 图书列表响应
type BookListResponse struct {
	Result *models.Books `json:"result"`
}

func BookListEndPoint(service *BookService) App.EndPoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*BookListRequest)
		return &BookListResponse{
			Result: service.GetList(req),
		}, err
	}
}
