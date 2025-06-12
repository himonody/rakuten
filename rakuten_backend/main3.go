package main

//
//import (
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"strconv"
//	"sync"
//	"time"
//
//	"github.com/gin-gonic/gin"
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/sqlite"
//	"gopkg.in/olahol/melody.v1"
//)
//
//// Message 表示一条系统消息结构体
//// 包含发送者、标题、内容、是否弹窗提示、创建时间和多个接收人
//// 每条消息可以推送给多个接收者，通过 Recipients 关联
//// GORM 会自动根据字段标签生成对应表结构
//
//type Message struct {
//	ID         uint        `gorm:"primary_key"`                // 消息唯一 ID
//	Title      string      `gorm:"type:varchar(255);not null"` // 消息标题
//	Content    string      `gorm:"type:text;not null"`         // 消息内容
//	SenderID   uint        `gorm:"not null"`                   // 发送者用户 ID
//	Popup      bool        `gorm:"default:false"`              // 是否显示为弹窗
//	CreatedAt  time.Time   // 创建时间
//	Recipients []Recipient `gorm:"foreignKey:MessageID"` // 接收人列表
//}
//
//// Recipient 表示消息接收人的结构体
//// 每个接收人记录消息是否已读、是否已送达、是否弹窗展示等状态信息
//
//type Recipient struct {
//	ID         uint       `gorm:"primary_key"` // 唯一 ID
//	MessageID  uint       // 所属消息 ID
//	ReceiverID uint       // 接收者用户 ID
//	Read       bool       // 是否已读
//	ReadAt     *time.Time // 读取时间（如果已读）
//	Delivered  bool       // 是否已推送送达
//	PopupShown bool       // 弹窗是否已显示
//}
//
//var (
//	db       *gorm.DB       // GORM 数据库实例
//	m        *melody.Melody // Melody 实例，用于 WebSocket 管理
//	userConn sync.Map       // 存储在线用户连接 userID -> *melody.Session
//)
//
//// initDB 初始化数据库连接和自动建表
//// 使用 SQLite 存储数据并自动迁移表结构
//func initDB() {
//	var err error
//	db, err = gorm.Open("sqlite3", "inbox.db")
//	if err != nil {
//		panic("failed to connect database")
//	}
//	db.AutoMigrate(&Message{}, &Recipient{})
//}
//
//// createMessageHandler 是创建消息的接口处理器
//// 从请求中读取消息内容、发送者、接收者列表、是否弹窗
//// 插入消息和接收人记录，并尝试向在线用户推送
//func createMessageHandler(c *gin.Context) {
//	type Request struct {
//		Title       string `json:"title"`        // 消息标题
//		Content     string `json:"content"`      // 消息内容
//		SenderID    uint   `json:"sender_id"`    // 发送者 ID
//		ReceiverIDs []uint `json:"receiver_ids"` // 接收者 ID 列表
//		Popup       bool   `json:"popup"`        // 是否弹窗
//	}
//	var req Request
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	msg := Message{
//		Title:     req.Title,
//		Content:   req.Content,
//		SenderID:  req.SenderID,
//		Popup:     req.Popup,
//		CreatedAt: time.Now(),
//	}
//
//	// 使用事务创建消息和接收人记录，确保一致性
//	err := db.Transaction(func(tx *gorm.DB) error {
//		if err := tx.Create(&msg).Error; err != nil {
//			return err
//		}
//		recs := make([]Recipient, 0)
//		for _, rid := range req.ReceiverIDs {
//			recs = append(recs, Recipient{
//				MessageID:  msg.ID,
//				ReceiverID: rid,
//				Read:       false,
//				Delivered:  false,
//				PopupShown: false,
//			})
//		}
//		return tx.Create(&recs).Error
//	})
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	// 推送消息给每个接收人（如果在线）
//	for _, rid := range req.ReceiverIDs {
//		go deliverMessageToUser(rid, msg)
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "created", "id": msg.ID})
//}
//
//// deliverMessageToUser 向指定用户推送消息（如果在线）
//// 并更新消息为已送达和弹窗状态
//func deliverMessageToUser(userID uint, msg Message) {
//	if val, ok := userConn.Load(userID); ok {
//		session := val.(*melody.Session)
//		jsonMsg := map[string]interface{}{
//			"type":       "message",
//			"title":      msg.Title,
//			"popup":      msg.Popup,
//			"content":    msg.Content,
//			"created_at": msg.CreatedAt,
//		}
//		if payload, err := json.Marshal(jsonMsg); err == nil {
//			session.Write(payload)
//			db.Model(&Recipient{}).Where("message_id = ? AND receiver_id = ?", msg.ID, userID).Updates(map[string]interface{}{
//				"delivered":   true,
//				"popup_shown": msg.Popup,
//			})
//		}
//	}
//}
//
//// getUnreadMessages 获取指定用户所有未读的消息
//// 使用 Preload 加载关联的消息内容
//func getUnreadMessages(c *gin.Context) {
//	uidStr := c.Query("user_id")
//	uid, _ := strconv.Atoi(uidStr)
//	var recs []Recipient
//	db.Preload("Message").Where("receiver_id = ? AND read = false", uid).Find(&recs)
//
//	var list []gin.H
//	for _, rec := range recs {
//		list = append(list, gin.H{
//			"id":         rec.Message.ID,
//			"title":      rec.Message.Title,
//			"content":    rec.Message.Content,
//			"popup":      rec.Message.Popup,
//			"delivered":  rec.Delivered,
//			"read":       rec.Read,
//			"created_at": rec.Message.CreatedAt,
//		})
//	}
//
//	c.JSON(http.StatusOK, gin.H{"messages": list})
//}
//
//// markAsRead 标记某条消息为已读
//// 需要传入消息 ID 和用户 ID，记录读取时间
//func markAsRead(c *gin.Context) {
//	msgID := c.Param("id")
//	uidStr := c.Query("user_id")
//	uid, _ := strconv.Atoi(uidStr)
//	readAt := time.Now()
//	db.Model(&Recipient{}).Where("message_id = ? AND receiver_id = ?", msgID, uid).
//		Updates(map[string]interface{}{"read": true, "read_at": readAt})
//	c.JSON(http.StatusOK, gin.H{"message": "marked as read"})
//}
//
//// wsHandler 处理 WebSocket 建立连接请求
//// 保存用户 ID 信息到 Session 中，供后续推送使用
//func wsHandler(c *gin.Context) {
//	uidStr := c.Query("user_id")
//	uid, _ := strconv.Atoi(uidStr)
//	err := m.HandleRequestWithKeys(c.Writer, c.Request, map[string]interface{}{"user_id": uid})
//	if err != nil {
//		return
//	}
//}
//
//// main 函数是程序入口，初始化数据库、设置路由和 WebSocket 事件
//func main() {
//	initDB()
//	defer db.Close()
//
//	r := gin.Default()
//	m = melody.New()
//
//	r.POST("/message", createMessageHandler)    // 创建消息
//	r.GET("/message/unread", getUnreadMessages) // 获取未读消息
//	r.POST("/message/:id/read", markAsRead)     // 标记消息为已读
//	r.GET("/ws", wsHandler)                     // 建立 WebSocket
//
//	// 建立连接时记录用户连接
//	m.HandleConnect(func(s *melody.Session) {
//		uid := s.Keys["user_id"].(int)
//		userConn.Store(uint(uid), s)
//	})
//
//	// 用户断开连接时清理
//	m.HandleDisconnect(func(s *melody.Session) {
//		uid := s.Keys["user_id"].(int)
//		userConn.Delete(uint(uid))
//	})
//
//	r.Run(":8080")
//}
