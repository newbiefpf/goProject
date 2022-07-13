package user

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      string `json:"age"`
	Sex      int    `json:"sex"`
	Role     int    `json:"role"`
}
