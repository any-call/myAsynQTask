package myAsynQTask

import "github.com/hibiken/asynq"

// Consumer 负责从指定队列消费任务并分发给处理器。
type Consumer struct {
	server *asynq.Server
	mux    *asynq.ServeMux
}

// NewConsumer 创建一个任务消费服务。
func NewConsumer(redisOpt asynq.RedisConnOpt, config asynq.Config) *Consumer {
	return &Consumer{
		server: asynq.NewServer(redisOpt, config),
		mux:    asynq.NewServeMux(),
	}
}

// Handle 注册一个任务处理器。
func (c *Consumer) Handle(taskType string, handler asynq.Handler) {
	c.mux.Handle(taskType, handler)
}

// HandleFunc 注册一个任务处理函数。
func (c *Consumer) HandleFunc(taskType string, handler asynq.HandlerFunc) {
	c.mux.HandleFunc(taskType, handler)
}

// Use 注册消费处理中间件。
func (c *Consumer) Use(middlewares ...asynq.MiddlewareFunc) {
	c.mux.Use(middlewares...)
}

// Start 启动任务消费服务，并阻塞直到服务停止。
func (c *Consumer) Start() error {
	return c.server.Run(c.mux)
}

// Stop 请求消费服务停止接收新任务。
func (c *Consumer) Stop() {
	c.server.Stop()
}

// Shutdown 等待消费服务优雅退出。
func (c *Consumer) Shutdown() {
	c.server.Shutdown()
}
