package Services

import (
	"fmt"
	"web-scene/AppInit"
	"web-scene/models"
	"web-scene/pkg/helper"
)

type BookService struct {
}

func NewBookService() *BookService {
	return &BookService{}
}

// List 获取图书列表
func (service *BookService) GetList(request *BookListRequest) (books *models.Books, err error) {
	books = &models.Books{}

	if pagination := helper.Paging(request.Page, request.Size); pagination != nil {
		err = pagination.Order("id desc").Find(books).Error
	}

	return
}

// GetDetail 获取图书详情
func (service *BookService) GetDetail(request *BookDetailRequest) (book *models.Book, err error) {
	book = &models.Book{}

	if AppInit.GetDB().First(book, request.ID).RowsAffected != 1 {
		err = fmt.Errorf("book:%d not exist", request.ID)
	}

	return
}
