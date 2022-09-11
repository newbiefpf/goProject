package tableModel

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username     string
	Password     string
	Email        string
	Github       string
	Age          int
	Sex          int
	Tel          string
	Professional string
}
