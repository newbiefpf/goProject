package user

import (
	"github.com/gin-gonic/gin"
)

//var SqlDb *sql.DB

func FindUserInfo(c *gin.Context) {
	//db := db.DB

	//err := SqlDb.QueryRow("select * from user ")
	//if err != nil {
	//	fmt.Println(err)
	//	c.JSON(http.StatusOK, err)
	//	return
	//}
	//Data := err
	//c.JSON(http.StatusOK, Data)
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
