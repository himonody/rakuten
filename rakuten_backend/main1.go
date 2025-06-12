package main

//
//import (
//	"encoding/json"
//	"log"
//	"net/http"
//	"strconv"
//	"strings"
//	"sync"
//	"time"
//
//	"github.com/gin-gonic/gin"
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/sqlite" // SQLite驱动
//	"gopkg.in/olahol/melody.v1"                // WebSocket库
//)
//
//// -----------------------------
//// 常量定义
//// -----------------------------
//const (
//	CategoryGeneral = "general" // 通用消息
//	CategorySystem  = "system"  // 系统消息
//	CategoryUser    = "user"    // 用户消息
//	CategoryAlert   = "alert"   // 警告消息
//
//	pushQueueSize   = 1000 // 推送任务队列缓冲大小
//	pushWorkerCount = 5    // 推送工作协程数量
//)
//
//// -----------------------------
//// 数据库模型定义
//// -----------------------------
//
//// Message 消息表，存储消息基本信息
//type Message struct {
//	ID        uint      `gorm:"primary_key" json:"id"`
//	Title     string    `json:"title"`     // 消息标题
//	Content   string    `json:"content"`   // 消息正文
//	SenderID  uint      `json:"sender_id"` // 发送者用户ID
//	Popup     bool      `json:"popup"`     // 是否弹窗通知
//	Category  string    `json:"category"`  // 消息分类
//	CreatedAt time.Time `json:"created_at"`
//}
//
//// Recipient 收件人表，关联消息和接收者，记录消息状态
//type Recipient struct {
//	ID         uint `gorm:"primary_key"`
//	MessageID  uint `gorm:"index"` // 关联消息ID
//	ReceiverID uint `gorm:"index"` // 接收者用户ID
//	Read       bool // 是否已读
//	Delivered  bool // 是否已推送送达
//	PopupShown bool // 是否弹窗已显示
//}
//
//// -----------------------------
//// 推送任务结构体
//// -----------------------------
//
//// PushTask 代表一个用户的一条推送消息任务
//type PushTask struct {
//	UserID  uint    // 接收推送的用户ID
//	Message Message // 需要推送的消息内容
//}
//
//// -----------------------------
//// 全局变量声明
//// -----------------------------
//
//var (
//	db        *gorm.DB       // 数据库连接实例
//	m         *melody.Melody // Melody WebSocket 管理实例
//	userConn  sync.Map       // 并发安全的用户ID->连接Session映射
//	pushQueue chan PushTask  // 推送任务缓冲队列（无阻塞设计）
//)
//
//// -----------------------------
//// 工具函数
//// -----------------------------
//
//// initDB 初始化数据库连接并自动迁移表结构
//func initDB() {
//	var err error
//	// 使用SQLite作为示范数据库
//	db, err = gorm.Open("sqlite3", "inbox.db")
//	if err != nil {
//		log.Fatalf("数据库连接失败: %v", err)
//	}
//	// 自动迁移消息表和收件人表结构
//	if err := db.AutoMigrate(&Message{}, &Recipient{}).Error; err != nil {
//		log.Fatalf("数据库迁移失败: %v", err)
//	}
//}
//
//// isValidCategory 校验消息分类是否合法
//func isValidCategory(cat string) bool {
//	switch cat {
//	case CategoryGeneral, CategorySystem, CategoryUser, CategoryAlert:
//		return true
//	}
//	return false
//}
//
//// enqueuePushTask 非阻塞地将推送任务写入推送队列
//// 如果队列已满，任务将被丢弃并记录日志
//func enqueuePushTask(task PushTask) {
//	select {
//	case pushQueue <- task:
//		// 成功入队
//	default:
//		// 队列满，丢弃任务，防止阻塞
//		log.Printf("推送队列已满，丢弃任务 user:%d message:%d", task.UserID, task.Message.ID)
//	}
//}
//
//// -----------------------------
//// WebSocket 推送工作协程
//// -----------------------------
//
//// pushWorker 持续监听推送队列，异步推送消息给在线用户
//func pushWorker(id int) {
//	for task := range pushQueue {
//		val, ok := userConn.Load(task.UserID)
//		if !ok {
//			// 用户未在线，跳过推送
//			continue
//		}
//		session := val.(*melody.Session)
//
//		// 组装推送消息体（JSON格式）
//		jsonMsg := map[string]interface{}{
//			"type":       "message",
//			"title":      task.Message.Title,
//			"popup":      task.Message.Popup,
//			"content":    task.Message.Content,
//			"category":   task.Message.Category,
//			"created_at": task.Message.CreatedAt.Format(time.RFC3339),
//		}
//
//		payload, err := json.Marshal(jsonMsg)
//		if err != nil {
//			log.Printf("worker %d: 消息序列化失败: %v", id, err)
//			continue
//		}
//
//		// 通过WebSocket推送消息
//		if err := session.Write(payload); err != nil {
//			log.Printf("worker %d: 推送用户 %d 消息失败: %v", id, task.UserID, err)
//			continue
//		}
//
//		// 更新数据库，标记该用户消息已推送送达，弹窗状态
//		if err := db.Model(&Recipient{}).
//			Where("message_id = ? AND receiver_id = ?", task.Message.ID, task.UserID).
//			Updates(map[string]interface{}{
//				"delivered":   true,
//				"popup_shown": task.Message.Popup,
//			}).Error; err != nil {
//			log.Printf("worker %d: 更新推送状态失败: %v", id, err)
//		}
//	}
//}
//
//// -----------------------------
//// HTTP接口处理
//// -----------------------------
//
//// createMessageHandler 接收客户端消息发布请求
//// 支持消息分类，消息写库后异步推送给指定用户
//func createMessageHandler(c *gin.Context) {
//	type Request struct {
//		Title       string `json:"title" binding:"required"`        // 消息标题，必填
//		Content     string `json:"content" binding:"required"`      // 消息正文，必填
//		SenderID    uint   `json:"sender_id" binding:"required"`    // 发送者用户ID，必填
//		ReceiverIDs []uint `json:"receiver_ids" binding:"required"` // 接收用户列表，必填
//		Popup       bool   `json:"popup"`                           // 是否弹窗通知，选填，默认false
//		Category    string `json:"category"`                        // 消息分类，选填，默认general
//	}
//
//	var req Request
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误：" + err.Error()})
//		return
//	}
//
//	// 默认分类为general，且转小写，去空格
//	req.Category = strings.ToLower(strings.TrimSpace(req.Category))
//	if req.Category == "" {
//		req.Category = CategoryGeneral
//	} else if !isValidCategory(req.Category) {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的消息分类"})
//		return
//	}
//
//	if len(req.ReceiverIDs) == 0 {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "接收者不能为空"})
//		return
//	}
//
//	// 构造消息对象
//	msg := Message{
//		Title:     req.Title,
//		Content:   req.Content,
//		SenderID:  req.SenderID,
//		Popup:     req.Popup,
//		Category:  req.Category,
//		CreatedAt: time.Now(),
//	}
//
//	// 使用事务保证消息和收件人同时写入
//	err := db.Transaction(func(tx *gorm.DB) error {
//		// 写入消息表
//		if err := tx.Create(&msg).Error; err != nil {
//			return err
//		}
//
//		// 批量写入收件人表
//		recs := make([]Recipient, 0, len(req.ReceiverIDs))
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
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库写入失败：" + err.Error()})
//		return
//	}
//
//	// 异步推送：遍历接收者，非阻塞写入推送队列
//	for _, rid := range req.ReceiverIDs {
//		enqueuePushTask(PushTask{
//			UserID:  rid,
//			Message: msg,
//		})
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "消息发布成功", "id": msg.ID})
//}
//
//// listMessagesHandler 查询用户消息列表，支持分页和分类过滤
//func listMessagesHandler(c *gin.Context) {
//	// 用户ID必须
//	userIDStr := c.Query("user_id")
//	if userIDStr == "" {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少 user_id"})
//		return
//	}
//	userID, err := strconv.Atoi(userIDStr)
//	if err != nil || userID <= 0 {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id 格式错误"})
//		return
//	}
//
//	// 支持按分类过滤，默认查询所有
//	category := strings.ToLower(c.Query("category"))
//	if category != "" && !isValidCategory(category) {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的消息分类"})
//		return
//	}
//
//	// 分页参数
//	pageStr := c.DefaultQuery("page", "1")
//	pageSizeStr := c.DefaultQuery("page_size", "20")
//	page, err1 := strconv.Atoi(pageStr)
//	pageSize, err2 := strconv.Atoi(pageSizeStr)
//	if err1 != nil || err2 != nil || page <= 0 || pageSize <= 0 || pageSize > 100 {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "分页参数错误"})
//		return
//	}
//
//	offset := (page - 1) * pageSize
//
//	// 先查收件人表，根据用户ID和分类过滤消息ID
//	var recipients []Recipient
//	query := db.Where("receiver_id = ?", userID)
//	if category != "" {
//		// 联表过滤消息分类
//		query = query.Joins("JOIN messages ON recipients.message_id = messages.id").
//			Where("messages.category = ?", category)
//	}
//	if err := query.Offset(offset).Limit(pageSize).Find(&recipients).Error; err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败：" + err.Error()})
//		return
//	}
//
//	// 收集消息ID
//	msgIDs := make([]uint, 0, len(recipients))
//	for _, r := range recipients {
//		msgIDs = append(msgIDs, r.MessageID)
//	}
//
//	if len(msgIDs) == 0 {
//		// 无消息
//		c.JSON(http.StatusOK, gin.H{"messages": []Message{}})
//		return
//	}
//
//	// 查询消息详情
//	var messages []Message
//	if err := db.Where("id IN (?)", msgIDs).Order("created_at DESC").Find(&messages).Error; err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询消息详情失败：" + err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"messages": messages})
//}
//
//// markReadHandler 标记某条消息为已读
//func markReadHandler(c *gin.Context) {
//	type Request struct {
//		UserID    uint `json:"user_id" binding:"required"`
//		MessageID uint `json:"message_id" binding:"required"`
//	}
//	var req Request
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误：" + err.Error()})
//		return
//	}
//
//	// 更新收件人表，标记已读
//	if err := db.Model(&Recipient{}).
//		Where("receiver_id = ? AND message_id = ?", req.UserID, req.MessageID).
//		Update("read", true).Error; err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库更新失败：" + err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "已标记为已读"})
//}
//
//// -----------------------------
//// WebSocket 连接处理
//// -----------------------------
//
//func wsHandler(c *gin.Context) {
//	uidStr := c.Query("user_id")
//	if uidStr == "" {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id 参数必填"})
//		return
//	}
//	uid, err := strconv.Atoi(uidStr)
//	if err != nil || uid <= 0 {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id 格式错误"})
//		return
//	}
//	// 将user_id放入melody Session Keys，方便管理连接
//	err := m.HandleRequestWithKeys(c.Writer, c.Request, map[string]interface{}{"user_id": uid})
//	if err != nil {
//		return
//	}
//}
//
//func main() {
//	// 初始化数据库连接
//	initDB()
//	defer db.Close()
//
//	// 初始化 Melody
//	m = melody.New()
//
//	// 初始化推送任务队列（带缓冲）
//	pushQueue = make(chan PushTask, pushQueueSize)
//
//	// 启动推送工作协程池
//	for i := 0; i < pushWorkerCount; i++ {
//		go pushWorker(i + 1)
//	}
//
//	// Melody连接建立时回调，保存用户连接到userConn
//	m.HandleConnect(func(s *melody.Session) {
//		uidVal, ok := s.Keys["user_id"]
//		if !ok {
//			// 未传user_id，直接断开连接
//			err := s.Close()
//			if err != nil {
//				return
//			}
//			return
//		}
//		uid := uidVal.(int)
//		userConn.Store(uint(uid), s) // 存储连接，方便推送
//		log.Printf("用户 %d 已连接 WebSocket", uid)
//	})
//
//	// 连接关闭时删除映射
//	m.HandleDisconnect(func(s *melody.Session) {
//		uidVal, ok := s.Keys["user_id"]
//		if ok {
//			uid := uidVal.(int)
//			userConn.Delete(uint(uid))
//			log.Printf("用户 %d 断开 WebSocket", uid)
//		}
//	})
//
//	// WebSocket接收消息示范（可以根据需求扩展）
//	m.HandleMessage(func(s *melody.Session, msg []byte) {
//		// 简单打印收到的消息
//		log.Printf("收到用户消息: %s", string(msg))
//	})
//
//	// Gin 路由
//	r := gin.Default()
//
//	// HTTP接口
//	r.POST("/messages", createMessageHandler) // 发布消息接口
//	r.GET("/messages", listMessagesHandler)   // 查询消息接口
//	r.POST("/messages/read", markReadHandler) // 标记消息已读接口
//	r.GET("/ws", wsHandler)                   // WebSocket连接接口
//
//	// 启动HTTP服务器
//	if err := r.Run(":8080"); err != nil {
//		log.Fatalf("服务器启动失败: %v", err)
//	}
//}
