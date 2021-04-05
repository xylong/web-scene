package Services

import (
	"web-scene/models"
	"web-scene/pkg/helper"
)

type BookService struct {
}

func NewBookService() *BookService {
	return &BookService{}
}

// List 获取图书列表
func (service *BookService) GetList(request *BookListRequest) (books *models.Books) {
	books = &models.Books{}

	if pagination := helper.Paging(request.Page, request.Size); pagination != nil {
		pagination.Order("id desc").Find(books)
	}

	return
}
