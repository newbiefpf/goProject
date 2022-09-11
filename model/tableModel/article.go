package tableModel

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	UserID      int
	Title       string
	ImgUrl      string
	Abstract    string
	Status      int
	ContentHtml string
}