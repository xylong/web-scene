package Services

import (
	"context"
	"web-scene/App"
	"web-scene/models"
)

// BookListRequest 图书列表请求
type BookListRequest struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

// BookListResponse 图书列表响应
type BookListResponse struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data *models.Books `json:"data,omitempty"`
}

// BookDetailRequest 图书详情请求
type BookDetailRequest struct {
	ID int `uri:"id" binding:"required,gt=0,max=100000"`
}

func BookListEndPoint(service *BookService) App.EndPoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*BookListRequest)
		return &BookListResponse{
			Data: service.GetList(req),
		}, err
	}
}
