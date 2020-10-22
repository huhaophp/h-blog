package admin

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	uploadDirPath string = "./public/uploadfile/"
)

// Controller Base
type FileController struct{}

// Upload uploads files to /tmp .
func (c *FileController) Store(r *ghttp.Request) {
	file := r.GetUploadFile("file")
	name, err := file.Save(uploadDirPath, true)
	if err != nil {
		_ = r.Response.WriteJsonExit(g.Map{
			"msg":    err.Error(),
			"status": "fail",
		})
	} else {
		_ = r.Response.WriteJsonExit(g.Map{
			"status": "success",
			"name":   "/uploadfile/" + name,
		})
	}
}

// Markdown File Store Upload uploads files to /tmp .
func (c *FileController) MarkdownFileStore(r *ghttp.Request) {
	file := r.GetUploadFile("editormd-image-file")
	name, err := file.Save(uploadDirPath, true)
	if err != nil {
		_ = r.Response.WriteJsonExit(g.Map{
			"success": 0,
			"path":    "",
			"message": err.Error(),
			"url":     "",
		})
	} else {
		_ = r.Response.WriteJsonExit(g.Map{
			"success": 1,
			"path":    "/uploadfile/" + name,
			"message": "上传成功",
			"url":     "/uploadfile/" + name,
		})
	}

}
