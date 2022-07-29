package utility

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

// 生成验证码
func GetRand() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}

// SendCode
// 发送验证码
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	mailUserName := "newbie0714@163.com" //邮箱账号
	mailPassword := "TWMJFOXUIIDNQXOK"   //邮箱授权码
	Subject := "Note验证码：10分钟有效"          //发送的主题
	e.From = "Get <newbie0714@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = Subject
	e.HTML = []byte("你的验证码为：<h1>" + code + "</h1>")
	return e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", mailUserName, mailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
}

// CodeSave
func CodeSave(code string) {

}
