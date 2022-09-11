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
	ContentHtml string `gorm:"column:contentHtml;type:MEDIUMTEXT;" json:"contentHtml"`
}
