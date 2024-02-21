package dto

type GetTwoCommentResp struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Content   string `json:"content"`
}

type GetRootCommentResp struct {
	ID                int              `json:"id"`
	CreatedAt         string           `json:"created_at"`
	UpdatedAt         string           `json:"updated_at"`
	Content           string           `json:"content"`
	UserID            int              `json:"user_id"`
	ProductID         int              `json:"product_id"`
	RootCommentID     int              `json:"root_comment_id"`
	ToCommentID       int              `json:"to_comment_id"`
	ChildCommentCount int              `json:"child_comment_count"`
	Name              string           `json:"name"`
	Avatar            string           `json:"avatar"`
	ThreeToComments   []ThreeToComment `json:"three_to_comments"`
}

type ThreeToComment struct {
	ID            int    `json:"id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	Content       string `json:"content"`
	UserID        int    `json:"user_id"`
	ProductID     int    `json:"product_id"`
	RootCommentID int    `json:"root_comment_id"`
	ToCommentID   int    `json:"to_comment_id"`
	Name          string `json:"name"`
	Avatar        string `json:"avatar"`
}
