package main

import (
	"fmt"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("home", "./ui/html/base.layout.tmpl", "./ui/html/footer.partial.tmpl", "./ui/html/home.page.tmpl")
	return r
}

func setupRouter(a *app) *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./ui/static")
	r.HTMLRender = createMyRender()
	fmt.Println(r.HTMLRender.Instance("home", nil))
	r.GET("/", a.home)
	return r
}

func main() {
	a := &app{}
	r := setupRouter(a)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
