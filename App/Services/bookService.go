package Services

import (
	"web-scene/AppInit"
	"web-scene/models"
)

type BookService struct {
}

func NewBookService() *BookService {
	return &BookService{}
}

// List 获取图书列表
func (service *BookService) GetList(request *BookListRequest) *models.Books {
	books := &models.Books{}
	AppInit.GetDB().Limit(request.Size).Order("id desc").Find(books)
	return books
}
