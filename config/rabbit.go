package config

import (
	"encoding/json"
	"fmt"
	"log"

	"drawsee/interfaces"
)

// MqConfig 定义了MQ配置
type MqConfig struct {
	QueueName        string
	RoutingKey       string
	QueueCount       int
	QueueConcurrency int
	ExchangeName     string
}

// LinkedQueue 定义了队列配置
type LinkedQueue struct {
	Name         string
	RoutingKey   string
	Concurrency  int
	ExchangeName string
}

// RabbitAdmin 管理RabbitMQ的声明
type RabbitAdmin struct {
	connectionFactory *ConnectionFactory
}

// ConnectionFactory 连接工厂
type ConnectionFactory struct {
	// 这里可以添加连接工厂的配置
}

// DeclareExchange 声明交换机
func (a *RabbitAdmin) DeclareExchange(exchangeName string) {
	fmt.Printf("Declaring exchange: %s\n", exchangeName)
	// 实际实现中，这里会调用RabbitMQ的API来声明交换机
}

// DeclareQueue 声明队列
func (a *RabbitAdmin) DeclareQueue(queueName string) {
	fmt.Printf("Declaring queue: %s\n", queueName)
	// 实际实现中，这里会调用RabbitMQ的API来声明队列
}

// DeclareBinding 声明绑定
func (a *RabbitAdmin) DeclareBinding(queueName, exchangeName, routingKey string) {
	fmt.Printf("Declaring binding: queue=%s, exchange=%s, routingKey=%s\n", queueName, exchangeName, routingKey)
	// 实际实现中，这里会调用RabbitMQ的API来声明绑定
}

// NewRabbitAdmin 创建一个新的RabbitAdmin
func NewRabbitAdmin(connectionFactory *ConnectionFactory) *RabbitAdmin {
	return &RabbitAdmin{connectionFactory: connectionFactory}
}

// MessageConverter 消息转换器接口
type MessageConverter interface {
	FromMessage(message []byte) (interface{}, error)
}

// JsonMessageConverter JSON消息转换器
type JsonMessageConverter struct{}

// FromMessage 实现了MessageConverter接口
func (c *JsonMessageConverter) FromMessage(message []byte) (interface{}, error) {
	var msg interfaces.AiTaskMessage
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

// RabbitListenerEndpoint RabbitMQ监听端点
type RabbitListenerEndpoint struct {
	Id              string
	QueueNames      []string
	Concurrency     int
	MessageListener func([]byte)
}

// RabbitListenerEndpointRegistrar RabbitMQ监听端点注册器
type RabbitListenerEndpointRegistrar struct {
	Endpoints []RabbitListenerEndpoint
}

// RegisterEndpoint 注册端点
func (r *RabbitListenerEndpointRegistrar) RegisterEndpoint(endpoint RabbitListenerEndpoint) {
	r.Endpoints = append(r.Endpoints, endpoint)
}

// ConfigureRabbitListeners 配置RabbitMQ监听器
func ConfigureRabbitListeners(
	aiTaskQueues []LinkedQueue,
	jsonMessageConverter MessageConverter,
	aiTaskWorker interfaces.AITaskWorker,
	registrar *RabbitListenerEndpointRegistrar,
) {
	for _, queue := range aiTaskQueues {
		endpoint := RabbitListenerEndpoint{
			Id:          queue.Name + ".Endpoint",
			QueueNames:  []string{queue.Name},
			Concurrency: queue.Concurrency,
			MessageListener: func(message []byte) {
				data, err := jsonMessageConverter.FromMessage(message)
				if err != nil {
					log.Printf("Error converting message: %v", err)
					return
				}
				if msg, ok := data.(interfaces.AiTaskMessage); ok {
					log.Printf("Consuming task: %s, thread: %s, starting processing", msg.TaskId, "main")
					aiTaskWorker.ProcessTask(msg)
				}
			},
		}
		registrar.RegisterEndpoint(endpoint)
	}
}
