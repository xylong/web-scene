package Services

import (
	"context"
	"web-scene/App"
)

// BookListRequest 图书列表请求
type BookListRequest struct {
	Page int `form:"page" binding:"required,gt=0,max=50"`
	Size int `form:"size" binding:"required,gt=0,max=100"`
}

// BookDetailRequest 图书详情请求
type BookDetailRequest struct {
	ID int `uri:"id" binding:"required,gt=0,max=100000"`
}

// BookResponse 图书详情响应
type BookResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// BookListEndPoint 图书列表业务
func BookListEndPoint(service *BookService) App.EndPoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*BookListRequest)
		res, err := service.GetList(req)
		return &BookResponse{
			Data: res,
		}, err
	}
}

// BookDetailEndPoint 图书详情业务
func BookDetailEndPoint(service *BookService) App.EndPoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*BookDetailRequest)
		res, err := service.GetDetail(req)
		return &BookResponse{
			Data: res,
		}, err
	}
}
