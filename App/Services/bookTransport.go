package Services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-scene/App"
)

// CreateBookListRequest 创建图书列表请求
func CreateBookListRequest() App.EncodeRequest {
	return func(context *gin.Context) (any interface{}, err error) {
		req := &BookListRequest{}
		if err = context.BindQuery(req); err != nil {
			return nil, err
		}

		return req, nil
	}
}

// CreateBookListResponse 创建图书列表响应
func CreateBookListResponse() App.DecodeResponse {
	return func(context *gin.Context, i interface{}) error {
		res := i.(*BookListResponse)
		context.JSON(http.StatusOK, res)
		return nil
	}
}
