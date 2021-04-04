package models

type (
	Book struct {
		ID           int     `gorm:"column:id;auto_increment;primary_key" json:"id"`
		Name         string  `gorm:"column:name;type:varchar(50);" json:"name"`
		Introduction string  `gorm:"column:introduction;type:text;" json:"introduction"`
		Price1       float64 `gorm:"column:price1;type:decimal;" json:"price_1"`
		Price2       float64 `gorm:"column:price1;type:decimal;" json:"price_2"`
		Author       string  `gorm:"column:author;type:varchar(30);" json:"author"`
		Press        string  `gorm:"column:press;varchar(50);" json:"press"`
		Date         string  `gorm:"column:date;type:varchar(50);" json:"date"`
		Kind         int     `gorm:"column:kind;type:int;" json:"kind"`
		KindName     string  `gorm:"column:kind_name;type:varchar(20);" json:"kind_name"`
	}

	Books []*Book
)
