package model

import (
	"lease/database"
	"time"
)

type Message struct {
	Model
	Content string `gorm:"column:content;type:text;not null;comment:消息内容" json:"content"`
	FromID  int    `gorm:"column:from_id;type:int;not null;comment:发送者ID" json:"from_id"`
	ToID    int    `gorm:"column:to_id;type:int;not null;comment:接收者ID" json:"to_id"`
	Status  int    `gorm:"column:status;type:tinyint(1);default:0;comment:消息状态" json:"status"`
}

type MessageUser struct {
	ID       int       `json:"id"`
	FriendID int       `json:"to_id"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"created_at"`
	Name     string    `json:"name"`
	Avatar   string    `json:"avatar"`
}

func GetChatUser(fromID int) (messageUserList []MessageUser) {
	database.DB.Raw("SELECT uni_table.id, friend_id, uni_table.content, uni_table.created_at, u.name, u.avatar FROM ( SELECT id, to_id as friend_id,content,created_at FROM messages WHERE (from_id = ?) AND (to_id <> ?) AND deleted_at IS NULL UNION SELECT id, from_id as friend_id,content,created_at FROM messages WHERE (from_id <> ?) AND (to_id = ?) AND deleted_at IS NULL ORDER BY created_at desc ) AS uni_table INNER JOIN users u on friend_id = u.id AND u.deleted_at IS NULL GROUP BY friend_id ORDER BY uni_table.created_at DESC;", fromID, fromID, fromID, fromID).Scan(&messageUserList)
	return
}

func GetMessage(fromID, toID int) (messageList []Message) {
	database.DB.Where("from_id = ? AND to_id = ?", fromID, toID).Or("from_id = ? AND to_id = ?", toID, fromID).Order("created_at").Find(&messageList)
	return
}

func AddMessage(message Message) {
	database.DB.Create(&message)
}
