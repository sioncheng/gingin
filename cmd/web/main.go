package main

import (
	"log"
	"os"
	"path"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sioncheng/gingin/cmd/repository"
	"github.com/sioncheng/gingin/cmd/service"
)

func createMyRender(p string) multitemplate.Renderer {
	base := path.Join(p, "/ui/html/base.layout.html")
	footer := path.Join(p, "/ui/html/footer.partial.html")
	r := multitemplate.NewRenderer()
	r.AddFromFiles("home", base, footer, path.Join(p, "/ui/html/home.page.html"))
	r.AddFromFiles("create", base, footer, path.Join(p, "/ui/html/create.page.html"))
	r.AddFromFiles("show", base, footer, path.Join(p, "/ui/html/show.page.html"))

	return r
}

func setupRouter(a *app) *gin.Engine {
	r := gin.Default()
	r.GET("/", a.Home)
	r.GET("/snippet/:id", a.ShowSnippet)
	return r
}

func main() {
	db, err := gorm.Open("mysql", "web:Web@12345@(127.0.0.1:3306)/snippetbox?parseTime=true")
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
	}
	defer func() {
		db.Close()
	}()

	sr := repository.SnippetRepositoryImpl{DB: db}
	sc := service.SnippetService{Repository: &sr}

	a := &app{sc}
	r := setupRouter(a)
	r.HTMLRender = createMyRender("./")
	r.Static("/static", path.Join("./", "/ui/static"))
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
