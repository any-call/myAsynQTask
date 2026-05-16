package myAsynQTask

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
)

// Producer 负责生产任务并写入 Redis 队列。
type Producer struct {
	client *asynq.Client
}

// NewProducer 创建一个任务生产者。
func NewProducer(redisOpt asynq.RedisConnOpt) *Producer {
	return &Producer{
		client: asynq.NewClient(redisOpt),
	}
}

// Enqueue 生产一个立即执行的任务。
func (p *Producer) Enqueue(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return p.client.EnqueueContext(ctx, task, opts...)
}

// EnqueueIn 生产一个延迟执行的任务。
func (p *Producer) EnqueueIn(ctx context.Context, task *asynq.Task, delay time.Duration, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	opts = append(opts, asynq.ProcessIn(delay))
	return p.client.EnqueueContext(ctx, task, opts...)
}

// EnqueueAt 生产一个在指定时间执行的任务。
func (p *Producer) EnqueueAt(ctx context.Context, task *asynq.Task, at time.Time, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	opts = append(opts, asynq.ProcessAt(at))
	return p.client.EnqueueContext(ctx, task, opts...)
}

// Close 关闭任务生产者使用的 Redis 连接。
func (p *Producer) Close() error {
	return p.client.Close()
}
