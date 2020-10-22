package admin

import (
	response "bbs/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	HomeTpl string = "admin/home.html"
)

type HomeController struct{}

// GET|后台首页
func (c *HomeController) Home(r *ghttp.Request) {
	response.ViewExit(r, layout, g.Map{"mainTpl": HomeTpl})
}
