package interfaces

// AITaskWorker 定义了AI任务工作接口
type AITaskWorker interface {
	ProcessTask(message AiTaskMessage)
}

// AiTaskMessage 定义了AI任务消息
type AiTaskMessage struct {
	TaskId string `json:"taskId"`
}
