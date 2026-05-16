package myAsynQTask

import "github.com/hibiken/asynq"

// Scheduler 负责按时间规则定时生产任务。
type Scheduler struct {
	scheduler *asynq.Scheduler
}

// NewScheduler 创建一个定时任务生产服务。
func NewScheduler(redisOpt asynq.RedisConnOpt, opts *asynq.SchedulerOpts) *Scheduler {
	return &Scheduler{
		scheduler: asynq.NewScheduler(redisOpt, opts),
	}
}

// Register 注册一个按定时表达式生产的任务。
func (s *Scheduler) Register(spec string, task *asynq.Task, opts ...asynq.Option) (string, error) {
	return s.scheduler.Register(spec, task, opts...)
}

// Start 启动定时任务生产服务，并阻塞直到调度器停止。
func (s *Scheduler) Start() error {
	return s.scheduler.Run()
}

// Shutdown 等待定时任务生产服务优雅退出。
func (s *Scheduler) Shutdown() {
	s.scheduler.Shutdown()
}
