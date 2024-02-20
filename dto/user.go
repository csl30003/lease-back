package dto

type LoginReq struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResp struct {
	Token  string `json:"token"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Gender int    `json:"gender"`
}

type RegisterReq struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
