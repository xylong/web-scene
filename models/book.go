package models

type (
	Book struct {
		ID int `gorm:"column:id;auto_increment;primary_key"`
		Name string `gorm:"column:name;type:varchar(50);"`
		Introduction string `gorm:"column:introduction;type:text;"`
		Price1 float64 `gorm:"column:price1;type:decimal;"`
		Price2 float64 `gorm:"column:price1;type:decimal;"`
		Author float64 `gorm:"column:author;type:varchar(30);"`
		Press float64 `gorm:"column:press;varchar(50);"`
		Date float64 `gorm:"column:date;type:varchar(50);"`
		Kind float64 `gorm:"column:kind;type:int;"`
		Kindstr float64 `gorm:"column:kind_str;type:varchar(20);"`
	}

	Books []*Book
)