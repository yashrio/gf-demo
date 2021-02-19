# gf-demo

## 部署

安装goFrame框架

创建数据库gf-demo， 示例数据库脚本在document文件夹内

## 运行

```go
gf run main.go --swagger
```

## 接口示例

登录接口使用postman访问，登录后得到token

| 接口名称 | 地址                             | 请求数据                                                 |
| -------- | -------------------------------- | -------------------------------------------------------- |
| 登录     | http://localhost:8199/auth/login | {"stype": "mobile","mobile": "13718878727"} |

其他接口可以使用swagger访问 http://localhost:8199/swagger

接口请求header

```html
Content-Type:application/json
Authorization:Bearer token
```

