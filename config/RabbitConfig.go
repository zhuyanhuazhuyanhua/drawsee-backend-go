package config

import (
	"fmt"
)

// RabbitConfig 定义了RabbitMQ配置
type RabbitConfig struct {
	AIConfig        MqConfig
	AnimationConfig MqConfig
}

// GetLinkedQueues 获取队列列表
func (c *RabbitConfig) GetLinkedQueues(mqConfig MqConfig) []LinkedQueue {
	var queues []LinkedQueue
	for i := 1; i <= mqConfig.QueueCount; i++ {
		queue := LinkedQueue{
			Name:         fmt.Sprintf("%s_%d", mqConfig.QueueName, i),
			RoutingKey:   fmt.Sprintf("%s_%d", mqConfig.RoutingKey, i),
			Concurrency:  mqConfig.QueueConcurrency,
			ExchangeName: mqConfig.ExchangeName,
		}
		queues = append(queues, queue)
	}
	return queues
}

// DeclareExchangeAndQueues 声明交换机和队列
func (c *RabbitConfig) DeclareExchangeAndQueues(taskQueues []LinkedQueue, admin *RabbitAdmin, mqConfig MqConfig) {
	admin.DeclareExchange(mqConfig.ExchangeName)
	for _, configQueue := range taskQueues {
		admin.DeclareQueue(configQueue.Name)
		admin.DeclareBinding(configQueue.Name, mqConfig.ExchangeName, configQueue.RoutingKey)
	}
}

// NewRabbitAdmin 创建一个新的RabbitAdmin
func (c *RabbitConfig) NewRabbitAdmin(connectionFactory *ConnectionFactory) *RabbitAdmin {
	return NewRabbitAdmin(connectionFactory)
}
