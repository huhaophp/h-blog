package middleware

import (
	"bbs/app/model/categories"
	response "bbs/library"
	"github.com/gogf/gf/net/ghttp"
)

// 后台检查登录
func AdminAuthCheck(r *ghttp.Request) {
	auth := r.Session.Get("admin", "")
	if auth == "" || auth == nil {
		if r.IsAjaxRequest() || r.Header.Get("Accept") == "application/json" {
			response.JsonExit(r, 401, "登录已过期")
		} else {
			_ = r.Session.Set("error", "登录已过期")
			r.Response.RedirectTo("/admin/login")
		}
	} else {
		r.Middleware.Next()
	}
}

// 初始化视图全局变量
func InitializeLayoutGlobalViewVariables(r *ghttp.Request) {
	var req categories.ListReqEntity
	req.Status = 1
	data, _ := categories.List(&req)
	r.GetView().Assign("categories", data)
	r.Middleware.Next()
}
