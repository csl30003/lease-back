package model

type Comment struct {
	Model
	Content       string `gorm:"type:text;not null;comment:'评论内容'" json:"content"`
	UserID        int    `gorm:"not null;comment:'用户ID'" json:"user_id"`
	ProductID     int    `gorm:"not null;comment:'商品ID'" json:"product_id"`
	RootCommentID int    `gorm:"default:null;comment:'顶级评论ID'" json:"root_comment_id"`
	ToCommentID   int    `gorm:"default:null;comment:'回复目标评论ID'" json:"to_comment_id"`
}
