package model

import "time"

// Message 消息表，存储消息基本信息
type Message struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`         // 消息标题
	Content   string    `gorm:"column:content" json:"content"`     // 消息正文
	SenderID  int       `gorm:"column:sender_id" json:"sender_id"` // 发送者用户ID
	Popup     int       `gorm:"column:popup" json:"popup"`         // 是否弹窗通知 0 否 1 是
	Category  int       `gorm:"column:category" json:"category"`   // 消息分类
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// Recipient 收件人表，关联消息和接收者，记录消息状态
type Recipient struct {
	ID         int       `gorm:"column:id;primary_key" json:"id"`
	MessageID  int       `gorm:"column:message_id" json:"message_id"`   // 关联消息ID
	ReceiverID int       `gorm:"column:receiver_id" json:"receiver_id"` // 接收者用户ID
	Read       int       `gorm:"column:read" json:"read"`               // 是否已读
	Delivered  int       `gorm:"column:delivered" json:"delivered"`     // 是否已推送送达
	PopupShown int       `gorm:"column:popup_shown" json:"popup_shown"` // 是否弹窗已显示
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}
