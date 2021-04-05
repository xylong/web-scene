package helper

import (
	"github.com/jinzhu/gorm"
	"web-scene/AppInit"
)

// Paging 分页
func Paging(page, size int) *gorm.DB {
	if page > 0 && size > 0 {
		return AppInit.GetDB().Limit(size).Offset((page - 1) * size)
	}
	return nil
}
