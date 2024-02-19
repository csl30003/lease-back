package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lease/model"
	"lease/response"
	"path/filepath"
	"strconv"
	"time"
)

func UploadProductMainImage(c *gin.Context) {
	file, err := c.FormFile("productMainImage")
	if err != nil {
		response.Failed(c, "上传失败")
		return
	}

	// 构建文件名
	fileExt := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt)

	// 保存上传的文件
	err = c.SaveUploadedFile(file, "D:\\HBuilderProjects\\lease\\static\\images\\product\\"+fileName)
	if err != nil {
		response.Failed(c, "上传失败")
		return
	}

	// 构建文件访问路径
	path := fmt.Sprintf("/static/images/product/%s", fileName)

	// 修改数据库中的主图
	productId, _ := strconv.Atoi(c.Param("id"))
	model.UpdateProductMainImage(productId, path)

	response.Success(c, "上传成功", path)
}

func UploadProductImage(c *gin.Context) {
	file, err := c.FormFile("productImage")
	if err != nil {
		response.Failed(c, "上传失败")
		return
	}

	// 构建文件名
	fileExt := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt)

	// 保存上传的文件
	err = c.SaveUploadedFile(file, "D:\\HBuilderProjects\\lease\\static\\images\\product\\"+fileName)
	if err != nil {
		response.Failed(c, "上传失败")
		return
	}

	// 构建文件访问路径
	path := fmt.Sprintf("/static/images/product/%s", fileName)

	// 把副图加入数据库
	productId, _ := strconv.Atoi(c.Param("id"))
	model.AddProductImage(productId, path)

	response.Success(c, "上传成功", path)
}
