package service

import (
	"github.com/gin-gonic/gin"
	"lease/model"
	"lease/response"
	"strconv"
)

func GetCategory(c *gin.Context) {
	parentId, _ := strconv.Atoi(c.Param("parentId"))

	categoryList := model.GetCategoryByParentID(parentId)

	response.Success(c, "获取成功", categoryList)
}
