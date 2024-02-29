package model

import (
	"lease/database"
	"time"
)

type Comment struct {
	Model
	Content       string `gorm:"type:text;not null;comment:'评论内容'" json:"content"`
	UserID        int    `gorm:"not null;comment:'用户ID'" json:"user_id"`
	ProductID     int    `gorm:"not null;comment:'商品ID'" json:"product_id"`
	RootCommentID int    `gorm:"default:null;comment:'顶级评论ID'" json:"root_comment_id"`
	ToCommentID   int    `gorm:"default:null;comment:'回复目标评论ID'" json:"to_comment_id"`
}

type TwoComment struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Content   string    `json:"content"`
}

// GetTwoComment 获取两条评论
func GetTwoComment(productId int) (twoComment []TwoComment) {
	database.DB.Raw("SELECT c.id, u.name, u.avatar, c.content, c.created_at FROM comments AS c JOIN users AS u ON c.user_id = u.id AND u.deleted_at IS NULL WHERE c.product_id = ? AND c.root_comment_id IS NULL AND c.deleted_at IS  NULL ORDER BY c.created_at DESC LIMIT 2;", productId).Scan(&twoComment)
	return
}

type RootCommentAndToCommentCount struct {
	Model
	Content           string `json:"content"`
	UserID            int    `json:"user_id"`
	ProductID         int    `json:"product_id"`
	RootCommentID     int    `json:"root_comment_id"`
	ToCommentID       int    `json:"to_comment_id"`
	ChildCommentCount int    `json:"child_comment_count"`
	Name              string `json:"name"`
	Avatar            string `json:"avatar"`
}

// GetRootCommentAndToCommentCount 获取顶级评论和其回复评论数量
func GetRootCommentAndToCommentCount(productID int) (rootCommentAndToCommentCount []RootCommentAndToCommentCount) {
	database.DB.Raw("SELECT c1.*, COUNT(c2.id) AS child_comment_count, u.name, u.avatar FROM comments c1 LEFT JOIN comments c2 ON c2.root_comment_id = c1.id AND c2.deleted_at IS NULL INNER JOIN users u ON c1.user_id = u.id AND u.deleted_at IS NULL WHERE c1.product_id = ? AND c1.root_comment_id IS NULL AND c1.deleted_at IS NULL GROUP BY c1.id, c1.created_at ORDER BY c1.created_at DESC;", productID).Scan(&rootCommentAndToCommentCount)
	return
}

type ThreeToComment struct {
	Model
	Content       string `json:"content"`
	UserID        int    `json:"user_id"`
	ProductID     int    `json:"product_id"`
	RootCommentID int    `json:"root_comment_id"`
	ToCommentID   int    `json:"to_comment_id"`
	Name          string `json:"name"`
	Avatar        string `json:"avatar"`
}

// GetThreeToComment 获取顶级评论的三条回复评论
func GetThreeToComment(rootCommentIDStr string) (threeToComment []ThreeToComment) {
	rawStr := "SELECT c1.*, u.name, u.avatar " +
		"FROM comments c1 " +
		"INNER JOIN users u ON c1.user_id = u.id AND u.deleted_at IS NULL " +
		"WHERE c1.root_comment_id IN (" + rootCommentIDStr + ") AND c1.deleted_at IS NULL " +
		"AND (SELECT COUNT(*) " +
		"FROM comments c2 " +
		"WHERE c2.root_comment_id = c1.root_comment_id AND c2.id <= c1.id AND c2.deleted_at IS NULL ) <= 3 " +
		"ORDER BY c1.created_at ASC;"
	database.DB.Raw(rawStr).Scan(&threeToComment)
	return
}

func GetCommentByID(id int) (comment ThreeToComment, err error) {
	// 连User表查
	err = database.DB.Table("comments").
		Select("comments.*, users.name, users.avatar").
		Joins("INNER JOIN users ON comments.user_id = users.id").
		Where("comments.id = ?", id).
		First(&comment).
		Error
	return
}

func GetCommentListByRootCommentID(rootCommentID int) (commentList []ThreeToComment, err error) {
	// 连User表查
	err = database.DB.Table("comments").
		Select("comments.*, users.name, users.avatar").
		Joins("INNER JOIN users ON comments.user_id = users.id").
		Where("comments.root_comment_id = ?", rootCommentID).
		Find(&commentList).
		Error
	return
}

func AddComment(comment Comment) {
	database.DB.Create(&comment)
	return
}
