# goProject
# 基于Gin、Gorm、Vue 实现的博客网站系统

> 
> 后台语言：Golang、框架：Gin、GORM


## 参考链接
GOLANG下载网址： https://golang.google.cn/dl/

参考文档 Module：https://www.kancloud.cn/aceld/golang/1958311

GORM中文官网：https://gorm.io/zh_CN/docs/

GIN中文官网：https://gin-gonic.com/zh-cn/docs/

## 整合 Swagger
参考文档： https://github.com/swaggo/gin-swagger
接口访问地址：http://localhost:8080/swagger/index.html
```text
// GetProblemList
// @Tags 公共方法
// @Summary 公共方法
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /problem-list [get]
```

## 安装jwt 设置token时使用
```shell
go get github.com/dgrijalva/jwt-go
```
参考文档 :https://github.com/dgrijalva/jwt-go
## 使用邮箱发送验证码
```shell
go get go get github.com/jordan-wright/email
```
参考文档 :https://github.com/jordan-wright/email
参考文档 :https://blog.51cto.com/u_6192297/5268325
## 配置 
+ 将 MailPassword 配置到环境变量中
