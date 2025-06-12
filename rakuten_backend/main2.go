package main

//
//import (
//	"context"
//	"encoding/json"
//	"errors"
//	"log"
//	"net/http"
//	"strconv"
//	"strings"
//	"sync"
//	"time"
//
//	"github.com/gin-gonic/gin"
//	"gopkg.in/olahol/melody.v1"
//	"gorm.io/driver/sqlite"
//	"gorm.io/gorm"
//	"gorm.io/gorm/logger"
//)
//
//const (
//	CategoryGeneral = "general"
//	CategorySystem  = "system"
//	CategoryUser    = "user"
//	CategoryAlert   = "alert"
//
//	pushQueueSize   = 1000
//	pushWorkerCount = 5
//)
//
//type Message struct {
//	ID        uint      `gorm:"primaryKey" json:"id"`
//	Title     string    `json:"title"`
//	Content   string    `json:"content"`
//	SenderID  uint      `json:"sender_id"`
//	Popup     bool      `json:"popup"`
//	Category  string    `json:"category"`
//	CreatedAt time.Time `json:"created_at"`
//}
//
//type Recipient struct {
//	ID         uint `gorm:"primaryKey"`
//	MessageID  uint `gorm:"index;not null"`
//	ReceiverID uint `gorm:"index;not null"`
//	Read       bool
//	Delivered  bool
//	PopupShown bool
//}
//
//type PushTask struct {
//	UserID  uint
//	Message Message
//}
//
//type Server struct {
//	db        *gorm.DB
//	melody    *melody.Melody
//	userConns sync.Map // map[uint]*melody.Session
//
//	pushQueue chan PushTask
//}
//
//func main() {
//	server, err := NewServer()
//	if err != nil {
//		log.Fatalf("启动失败: %v", err)
//	}
//	server.Run(":8080")
//}
//
//func NewServer() (*Server, error) {
//	// 使用带日志的 GORM
//	db, err := gorm.Open(sqlite.Open("inbox.db"), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Info),
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	if err := db.AutoMigrate(&Message{}, &Recipient{}); err != nil {
//		return nil, err
//	}
//
//	m := melody.New()
//
//	s := &Server{
//		db:        db,
//		melody:    m,
//		pushQueue: make(chan PushTask, pushQueueSize),
//	}
//
//	// 启动推送工作池
//	for i := 0; i < pushWorkerCount; i++ {
//		go s.pushWorker(i + 1)
//	}
//
//	// Melody 事件注册
//	m.HandleConnect(s.onConnect)
//	m.HandleDisconnect(s.onDisconnect)
//
//	return s, nil
//}
//
//func (s *Server) Run(addr string) {
//	r := gin.Default()
//
//	r.POST("/messages", s.createMessageHandler)
//	r.GET("/messages", s.listMessagesHandler)
//	r.POST("/messages/read", s.markReadHandler)
//	r.GET("/ws", s.wsHandler)
//
//	log.Printf("服务启动，监听 %s", addr)
//	if err := r.Run(addr); err != nil {
//		log.Fatalf("Gin 运行失败: %v", err)
//	}
//}
//
//// ----------- WebSocket 连接管理 ------------
//
//func (s *Server) onConnect(session *melody.Session) {
//	userIDRaw, ok := session.Get("user_id")
//	if !ok {
//		log.Printf("新连接未携带 user_id，拒绝")
//		session.Close()
//		return
//	}
//
//	userID, ok := userIDRaw.(int)
//	if !ok || userID <= 0 {
//		log.Printf("user_id 无效: %v", userIDRaw)
//		session.Close()
//		return
//	}
//
//	s.userConns.Store(uint(userID), session)
//	log.Printf("用户 %d 连接建立", userID)
//}
//
//func (s *Server) onDisconnect(session *melody.Session) {
//	userIDRaw, ok := session.Get("user_id")
//	if !ok {
//		return
//	}
//	userID, ok := userIDRaw.(int)
//	if !ok {
//		return
//	}
//	s.userConns.Delete(uint(userID))
//	log.Printf("用户 %d 断开连接", userID)
//}
//
//func (s *Server) wsHandler(c *gin.Context) {
//	userIDStr := c.Query("user_id")
//	if userIDStr == "" {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id 参数必填"})
//		return
//	}
//	userID, err := strconv.Atoi(userIDStr)
//	if err != nil || userID <= 0 {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id 格式错误"})
//		return
//	}
//
//	err = s.melody.HandleRequestWithKeys(c.Writer, c.Request, map[string]interface{}{"user_id": userID})
//	if err != nil {
//		log.Printf("WebSocket 连接处理失败: %v", err)
//	}
//}
//
//// ----------- 消息处理 ------------
//
//// 校验消息分类
//func isValidCategory(cat string) bool {
//	switch cat {
//	case CategoryGeneral, CategorySystem, CategoryUser, CategoryAlert:
//		return true
//	}
//	return false
//}
//
//// 创建消息接口
//func (s *Server) createMessageHandler(c *gin.Context) {
//	var req struct {
//		Title       string `json:"title" binding:"required"`
//		Content     string `json:"content" binding:"required"`
//		SenderID    uint   `json:"sender_id" binding:"required"`
//		ReceiverIDs []uint `json:"receiver_ids" binding:"required,min=1"`
//		Popup       bool   `json:"popup"`
//		Category    string `json:"category"`
//	}
//
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误: " + err.Error()})
//		return
//	}
//
//	req.Category = strings.ToLower(strings.TrimSpace(req.Category))
//	if req.Category == "" {
//		req.Category = CategoryGeneral
//	} else if !isValidCategory(req.Category) {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的消息分类"})
//		return
//	}
//
//	msg := Message{
//		Title:     req.Title,
//		Content:   req.Content,
//		SenderID:  req.SenderID,
//		Popup:     req.Popup,
//		Category:  req.Category,
//		CreatedAt: time.Now(),
//	}
//
//	ctx := c.Request.Context()
//
//	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
//		if err := tx.Create(&msg).Error; err != nil {
//			return err
//		}
//
//		recipients := make([]Recipient, len(req.ReceiverIDs))
//		for i, rid := range req.ReceiverIDs {
//			recipients[i] = Recipient{
//				MessageID:  msg.ID,
//				ReceiverID: rid,
//			}
//		}
//
//		if err := tx.Create(&recipients).Error; err != nil {
//			return err
//		}
//		return nil
//	})
//
//	if err != nil {
//		log.Printf("创建消息失败: %v", err)
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库写入失败: " + err.Error()})
//		return
//	}
//
//	// 异步推送
//	go func() {
//		for _, rid := range req.ReceiverIDs {
//			task := PushTask{UserID: rid, Message: msg}
//			s.enqueuePushTask(task)
//		}
//	}()
//
//	c.JSON(http.StatusOK, gin.H{"message": "消息发布成功", "id": msg.ID})
//}
//
//func (s *Server) enqueuePushTask(task PushTask) {
//	select {
//	case s.pushQueue <- task:
//		// 入队成功
//	default:
//		log.Printf("推送队列已满，丢弃任务 user:%d message:%d", task.UserID, task.Message.ID)
//	}
//}
//
//func (s *Server) pushWorker(workerID int) {
//	for task := range s.pushQueue {
//		val, ok := s.userConns.Load(task.UserID)
//		if !ok {
//			// 用户不在线
//			continue
//		}
//		session, ok := val.(*melody.Session)
//		if !ok {
//			continue
//		}
//
//		msgPayload := map[string]interface{}{
//			"type":       "message",
//			"title":      task.Message.Title,
//			"popup":      task.Message.Popup,
//			"content":    task.Message.Content,
//			"category":   task.Message.Category,
//			"created_at": task.Message.CreatedAt.Format(time.RFC3339),
//		}
//
//		data, err := json.Marshal(msgPayload)
//		if err != nil {
//			log.Printf("[worker %d] 消息序列化失败: %v", workerID, err)
//			continue
//		}
//
//		if err := session.Write(data); err != nil {
//			log.Printf("[worker %d] 推送用户 %d 消息失败: %v", workerID, task.UserID, err)
//			continue
//		}
//
//		// 更新数据库状态
//		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//		err = s.db.WithContext(ctx).Model(&Recipient{}).
//			Where("message_id = ? AND receiver_id = ?", task.Message.ID, task.UserID).
//			Update("delivered", true).Error
//		cancel()
//
//		if err != nil {
//			log.Printf("[worker %d] 更新推送状态失败 user:%d message:%d err:%v", workerID, task.UserID, task.Message.ID, err)
//		}
//	}
//}
//
//// 查询消息接口，支持分页
//func (s *Server) listMessagesHandler(c *gin.Context) {
//	userIDStr := c.Query("user_id")
//	pageStr := c.DefaultQuery("page", "1")
//	pageSizeStr := c.DefaultQuery("page_size", "10")
//
//	userID, err := strconv.Atoi(userIDStr)
//	if err != nil || userID <= 0 {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id 参数错误"})
//		return
//	}
//
//	page, _ := strconv.Atoi(pageStr)
//	if page <= 0 {
//		page = 1
//	}
//	pageSize, _ := strconv.Atoi(pageSizeStr)
//	if pageSize <= 0 || pageSize > 50 {
//		pageSize = 10
//	}
//
//	var messages []Message
//	offset := (page - 1) * pageSize
//
//	// 联表查询只获取用户相关消息
//	err = s.db.WithContext(c.Request.Context()).
//		Model(&Message{}).
//		Joins("JOIN recipients ON recipients.message_id = messages.id").
//		Where("recipients.receiver_id = ?", userID).
//		Offset(offset).
//		Limit(pageSize).
//		Order("messages.created_at DESC").
//		Find(&messages).Error
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询消息失败: " + err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"messages": messages})
//}
//
//// 标记消息已读接口
//func (s *Server) markReadHandler(c *gin.Context) {
//	var req struct {
//		MessageID uint `json:"message_id" binding:"required"`
//		UserID    uint `json:"user_id" binding:"required"`
//	}
//
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误: " + err.Error()})
//		return
//	}
//
//	if err := s.markMessageRead(c.Request.Context(), req.MessageID, req.UserID); err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败: " + err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "标记已读成功"})
//}
//
//func (s *Server) markMessageRead(ctx context.Context, messageID, userID uint) error {
//	result := s.db.WithContext(ctx).
//		Model(&Recipient{}).
//		Where("message_id = ? AND receiver_id = ?", messageID, userID).
//		Update("read", true)
//
//	if result.Error != nil {
//		return result.Error
//	}
//
//	if result.RowsAffected == 0 {
//		return errors.New("未找到对应的消息记录")
//	}
//
//	return nil
//}
