# GoWeb--轻量级golang语言web框架
## 目前实现的功能
1.简单的路由协议

2.针对单个路由的链式Middleware

## 使用方法：

```
import (
	"github.com/anakin/goweb"
	"net/http"
)

type AController struct {
	goweb.Controller
}

func (t *AController) GetHandler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]int)
	m["test"] = 1
	m["name"] = 2
	t.Data = m
	t.ToJson(w, r)
}


//middleware
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("agent:", r.UserAgent())
		next.ServeHTTP(w, r)
	}
}

func main() {
	a := &AController{}
	router := goweb.NewRouter()
	router.GET("/", a.GetHandler,Auth)
	goweb := goweb.New()
	goweb.Run(router)
}
```

## TODO
1.复杂路由支持（正则，路由组）

2.添加ORM

3.Session支持(准备把beego的session集成进来)

4.Middleware完善

5.html模板功能

6.安全机制

