package dbModel

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"column:username;type:varchar(20);" json:"username"`
	Password     string `gorm:"column:password;type:varchar(200);" json:"password"`
	Email        string `gorm:"column:email;type:varchar(36);" json:"email"`
	Github       string `gorm:"column:github;type:varchar(36);" json:"github"`
	Age          int    `gorm:"column:age;type:int;" json:"age"`
	Sex          int    `gorm:"column:sex;type:int;" json:"sex"`
	Tel          string `gorm:"column:tel;type:varchar(20);" json:"tel"`
	Professional string `gorm:"column:professional;type:varchar(20);" json:"professional"`
	Article      []Article
}
type Article struct {
	gorm.Model
	UserID      uint
	Title       string `gorm:"column:title;type:varchar(36);" json:"title"`
	ImgUrl      string `gorm:"column:imgUrl;type:varchar(500);" json:"imgUrl"`
	Abstract    string `gorm:"column:abstract;type:varchar(36);" json:"abstract"`
	Status      int    `gorm:"column:status;type:int;" json:"status"`
	ContentHtml string `gorm:"column:contentHtml;type:MEDIUMTEXT;" json:"contentHtml"`
	ArticleLink ArticleLink
}

type ArticleLink struct {
	gorm.Model
	ArticleID uint
	StepOn    int    `gorm:"column:stepOn;type:int;" json:"stepOn"`
	GiveLike  int    `gorm:"column:giveLike;type:int;" json:"giveLike"`
	Discuss   string `gorm:"column:discuss;type:varchar(36);" json:"discuss"`
}
type Discuss struct {
	gorm.Model
	ArticleID uint
	Comment   string `gorm:"column:comment;type:varchar(36);" json:"comment"`
	Status    int    `gorm:"column:status;type:int;" json:"status"`
}
