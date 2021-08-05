package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"path"
)

func createMyRender(p string) multitemplate.Renderer {
	base := path.Join(p, "/ui/html/base.layout.html")
	footer := path.Join(p, "/ui/html/footer.partial.html")
	r := multitemplate.NewRenderer()
	r.AddFromFiles("home", base, footer, path.Join(p , "/ui/html/home.page.html"))
	r.AddFromFiles("create", base, footer, path.Join(p , "/ui/html/create.page.html"))
	r.AddFromFiles("show", base, footer, path.Join(p , "/ui/html/show.page.html"))

	return r
}

func setupRouter(a *app) *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./ui/static")
	r.GET("/", a.home)
	return r
}

func main() {
	a := &app{}
	r := setupRouter(a)
	r.HTMLRender = createMyRender("./")
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
