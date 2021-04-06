package App

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	// EndPoint 业务函数
	// request 业务参数
	// response 处理结果
	EndPoint func(ctx context.Context, request interface{}) (response interface{}, err error)

	// EncodeRequest 处理请求参数
	EncodeRequest func(*gin.Context) (interface{}, error)

	// DecodeResponse 返回结果处理函数
	DecodeResponse func(*gin.Context, interface{}) error
)

// RegisterHandler 注册行为
func RegisterHandler(point EndPoint, encode EncodeRequest, decode DecodeResponse) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		defer func() {
			if r := recover(); r != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": fmt.Sprintf("fatal error:%s", r),
				})
				return
			}
		}()

		// 1.获取参数
		req, err := encode(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":   http.StatusBadRequest,
				"status": "param error:" + err.Error(),
			})
			return
		}
		// 2.执行业务
		rsp, err := point(ctx, req)
		if err != nil {
			fmt.Fprintln(gin.DefaultWriter, "response error", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":   http.StatusInternalServerError,
				"status": "response error:" + err.Error(),
			})
			return
		}
		// 3.处理业务员结果
		err = decode(ctx, rsp)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":   http.StatusInternalServerError,
				"status": "server error:" + err.Error(),
			})
			return
		}
	}
}
