package service

import (
	"github.com/gin-gonic/gin"
	"lease/dto"
	"lease/model"
	"lease/response"
	"reflect"
	"strconv"
)

func GetTwoComment(c *gin.Context) {
	parentId := c.Param("productId")
	parentIdInt, err := strconv.Atoi(parentId)
	if err != nil {
		response.Failed(c, "参数错误")
		return
	}

	twoComment := model.GetTwoComment(parentIdInt)

	var getTwoCommentResp []dto.GetTwoCommentResp
	// 把时间格式化为字符串
	for i := 0; i < len(twoComment); i++ {
		getTwoCommentResp = append(getTwoCommentResp, dto.GetTwoCommentResp{
			ID:        twoComment[i].ID,
			CreatedAt: twoComment[i].CreatedAt.Format("2006-01-02 15:04:05"),
			Name:      twoComment[i].Name,
			Avatar:    twoComment[i].Avatar,
			Content:   twoComment[i].Content,
		})
	}
	response.Success(c, "获取成功", getTwoCommentResp)
}

func GetRootComment(c *gin.Context) {
	parentId := c.Param("productId")
	parentIdInt, err := strconv.Atoi(parentId)
	if err != nil {
		response.Failed(c, "参数错误")
		return
	}

	rootCommentAndToCommentCount := model.GetRootCommentAndToCommentCount(parentIdInt)

	// 将rootCommentAndToCommentCount里的每个id组合成一个字符串，并且用逗号隔开
	var rootCommentIDStr string
	for i := 0; i < len(rootCommentAndToCommentCount); i++ {
		rootCommentIDStr += strconv.Itoa(rootCommentAndToCommentCount[i].ID) + ","
	}
	// 如果rootCommentIDStr长度不为0就去掉最后一个逗号
	if len(rootCommentIDStr) != 0 {
		rootCommentIDStr = rootCommentIDStr[:len(rootCommentIDStr)-1]
	}

	threeToComment := model.GetThreeToComment(rootCommentIDStr)

	// 把rootCommentAndToCommentCount和threeToComment组合成一个数组
	var getRootCommentResp []dto.GetRootCommentResp
	for i := 0; i < len(rootCommentAndToCommentCount); i++ {
		var threeToComments []dto.ThreeToComment
		for j := 0; j < len(threeToComment); j++ {
			if rootCommentAndToCommentCount[i].ID == threeToComment[j].RootCommentID {
				threeToComments = append(threeToComments, dto.ThreeToComment{
					ID:            threeToComment[j].ID,
					CreatedAt:     threeToComment[j].CreatedAt.Format("2006-01-02 15:04:05"),
					UpdatedAt:     threeToComment[j].UpdatedAt.Format("2006-01-02 15:04:05"),
					Content:       threeToComment[j].Content,
					UserID:        threeToComment[j].UserID,
					ProductID:     threeToComment[j].ProductID,
					RootCommentID: threeToComment[j].RootCommentID,
					ToCommentID:   threeToComment[j].ToCommentID,
					Name:          threeToComment[j].Name,
					Avatar:        threeToComment[j].Avatar,
				})
			}
		}

		getRootCommentResp = append(getRootCommentResp, dto.GetRootCommentResp{
			ID:                rootCommentAndToCommentCount[i].ID,
			CreatedAt:         rootCommentAndToCommentCount[i].CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:         rootCommentAndToCommentCount[i].UpdatedAt.Format("2006-01-02 15:04:05"),
			Content:           rootCommentAndToCommentCount[i].Content,
			UserID:            rootCommentAndToCommentCount[i].UserID,
			ProductID:         rootCommentAndToCommentCount[i].ProductID,
			RootCommentID:     rootCommentAndToCommentCount[i].RootCommentID,
			ToCommentID:       rootCommentAndToCommentCount[i].ToCommentID,
			ChildCommentCount: rootCommentAndToCommentCount[i].ChildCommentCount,
			Name:              rootCommentAndToCommentCount[i].Name,
			Avatar:            rootCommentAndToCommentCount[i].Avatar,
			ThreeToComments:   threeToComments,
		})
	}
	response.Success(c, "获取成功", getRootCommentResp)
}

func GetCommentList(c *gin.Context) {
	commentId := c.Param("commentId")
	commentIdInt, err := strconv.Atoi(commentId)
	if err != nil {
		response.Failed(c, "参数错误")
		return
	}
	if commentIdInt == 0 {
		response.Failed(c, "参数错误")
		return
	}

	rootComment, err := model.GetCommentByID(commentIdInt)
	if err != nil {
		response.Failed(c, "获取失败")
		return
	}

	toCommentList, err := model.GetCommentListByRootCommentID(commentIdInt)
	if err != nil {
		response.Failed(c, "获取失败")
		return
	}

	// 将rootComment和toCommentList组合成一个数组
	getCommentListResp := []dto.ThreeToComment{
		{
			ID:            rootComment.ID,
			CreatedAt:     rootComment.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:     rootComment.UpdatedAt.Format("2006-01-02 15:04:05"),
			Content:       rootComment.Content,
			UserID:        rootComment.UserID,
			ProductID:     rootComment.ProductID,
			RootCommentID: rootComment.RootCommentID,
			ToCommentID:   rootComment.ToCommentID,
			Name:          rootComment.Name,
			Avatar:        rootComment.Avatar,
		},
	}
	for i := 0; i < len(toCommentList); i++ {
		getCommentListResp = append(getCommentListResp, dto.ThreeToComment{
			ID:            toCommentList[i].ID,
			CreatedAt:     toCommentList[i].CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:     toCommentList[i].UpdatedAt.Format("2006-01-02 15:04:05"),
			Content:       toCommentList[i].Content,
			UserID:        toCommentList[i].UserID,
			ProductID:     toCommentList[i].ProductID,
			RootCommentID: toCommentList[i].RootCommentID,
			ToCommentID:   toCommentList[i].ToCommentID,
			Name:          toCommentList[i].Name,
			Avatar:        toCommentList[i].Avatar,
		})
	}

	response.Success(c, "获取成功", getCommentListResp)
}

func AddComment(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	userId := int(claimsValueElem.FieldByName("ID").Int())

	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		response.Failed(c, "参数错误")
		return
	}

	comment.UserID = userId

	model.AddComment(comment)
	response.Success(c, "添加成功", comment)
}
