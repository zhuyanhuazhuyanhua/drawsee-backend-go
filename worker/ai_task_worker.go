package worker

import (
	"drawsee/interfaces"
	"log"
	"time"
)

// AITaskWorker AI任务工作器
type AITaskWorker struct{}

// ProcessTask 处理AI任务
func (w *AITaskWorker) ProcessTask(message interfaces.AiTaskMessage) {
	log.Printf("Processing task: %s", message.TaskId)
	time.Sleep(1 * time.Second) // 模拟任务处理时间
	log.Printf("Task %s processed", message.TaskId)
}
