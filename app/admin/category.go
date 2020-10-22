package admin

import (
	"bbs/app/model/categories"
	response "bbs/library"
	"errors"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strconv"
)

const (
	CategoryListTpl   string = "admin/category/list.html"
	CategoryCreateTpl string = "admin/category/create.html"
	CategoryEditTpl   string = "admin/category/edit.html"
)

// Controller Base
type CategoryController struct{}

// Category List
func (c *CategoryController) List(r *ghttp.Request) {
	var req categories.ListReqEntity
	req.Status = -1
	list, err := categories.List(&req)
	if err != nil {
		response.ViewExit(r, layout, g.Map{"list": list, "mainTpl": errorTpl, "error": err.Error()})
	} else {
		response.ViewExit(r, layout, g.Map{"list": list, "mainTpl": CategoryListTpl})
	}
}

// Add add category
func (c *CategoryController) Add(r *ghttp.Request) {
	if r.Method == "GET" {
		response.ViewExit(r, layout, g.Map{"mainTpl": CategoryCreateTpl})
	}
	var data *categories.AddReqEntity
	if err := r.Parse(&data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	if categories.DoesTheNameExist(data.Name, 0) {
		response.RedirectBackWithError(r, gerror.New("名字已存在"))
	}
	insid, err := categories.Add(data)
	if err != nil || insid <= 0 {
		response.RedirectBackWithError(r, gerror.New("添加失败"))
	} else {
		response.RedirectToWithMessage(r, "/admin/categories", "添加成功")
	}
}

// Edit edit category
func (c *CategoryController) Edit(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterValue("id").(string))
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	if r.Method == "GET" {
		item, _ := categories.Find(id)
		response.ViewExit(r, layout, g.Map{"item": item, "mainTpl": CategoryEditTpl})
	}
	var data *categories.AddReqEntity
	if err := r.Parse(&data); err != nil {
		response.RedirectBackWithError(r, err)
	}
	if categories.DoesTheNameExist(data.Name, id) {
		response.RedirectBackWithError(r, gerror.New("名字已存在"))
	}
	insid, err := categories.Edit(id, data)
	if err != nil || insid <= 0 {
		response.RedirectBackWithError(r, gerror.New("编辑失败"))
	} else {
		response.RedirectToWithMessage(r, "/admin/categories", "编辑成功")
	}
}

// Delete del category
func (c *CategoryController) Delete(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterValue("id").(string))
	if err != nil {
		response.RedirectBackWithError(r, err)
	}
	rows, err := categories.Del(id)
	if err != nil || rows <= 0 {
		response.RedirectBackWithError(r, errors.New("删除失败"))
	} else {
		response.RedirectToWithMessage(r, "/admin/categories", "删除成功")
	}
}
