package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Articles     []Article
	Attentions   []Attention
	Username     string `gorm:"column:username;type:varchar(20);" json:"username"`
	Password     string `gorm:"column:password;type:varchar(200);" json:"password"`
	Email        string `gorm:"column:email;type:varchar(36);" json:"email"`
	Github       string `gorm:"column:github;type:varchar(36);" json:"github"`
	Age          int    `gorm:"column:age;type:int;" json:"age"`
	Sex          int    `gorm:"column:sex;type:int;" json:"sex"`
	Tel          string `gorm:"column:tel;type:varchar(20);" json:"tel"`
	Professional string `gorm:"column:professional;type:varchar(20);" json:"professional"`
}
type Article struct {
	gorm.Model
	UserID        uint
	ArticleInfoID uint
	ContentHtml   string `gorm:"column:contentHtml;type:MEDIUMTEXT;" json:"contentHtml"`
}
type ArticleInfo struct {
	gorm.Model
	ArticleID uint
	Title     string `gorm:"column:title;type:varchar(50);" json:"title"`
	Cover     string `gorm:"column:cover;type:varchar(50);" json:"cover"`
	Like      int    `gorm:"column:like;type:int;" json:"like"`
	Unlike    int    `gorm:"column:unlike;type:int;" json:"unlike"`
}
type Attention struct {
	gorm.Model
	UserID uint
}

type CodeLite struct {
	gorm.Model
	Email string `gorm:"column:email;type:varchar(36);" json:"email"`
	Code  string `gorm:"column:code;type:varchar(36);" json:"code"`
}
