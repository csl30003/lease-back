package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lease/response"
	"path/filepath"
	"time"
)

func UploadProductImage(c *gin.Context) {
	file, err := c.FormFile("productImage")

	if err != nil {
		response.Failed(c, "上传失败")
		return
	}

	// 构建文件名
	fileExt := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt)

	fmt.Println(fileName)

	// 保存上传的文件
	err = c.SaveUploadedFile(file, "D:\\HBuilderProjects\\lease\\static\\images\\test\\"+fileName)
	if err != nil {
		response.Failed(c, "上传失败")
		return
	}

	// 构建文件访问路径
	path := fmt.Sprintf("/static/images/avatar/%s", fileName)

	response.Success(c, "上传成功", path)
}
