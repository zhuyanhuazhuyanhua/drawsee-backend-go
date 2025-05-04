// package main

// import (
// 	"drawsee/config"
// 	"fmt"
// )

// func main() {
// 	cfg, err := config.LoadConfig("config.yaml")
// 	if err != nil {
// 		fmt.Printf("Failed to load config: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("Doubao Base URL: %s\n", cfg.Doubao.BaseURL)
// 	fmt.Printf("Doubao API Key: %s\n", cfg.Doubao.APIKey)
// 	fmt.Printf("Doubao Model Name: %s\n", cfg.Doubao.ModelName)

//		// 其他模型的配置也可以类似地使用
//	} langchainConfig.go测试

// package main

// import (
// 	"fmt"

// 	"drawsee/config"
// )

// func main() {
// 	// 创建ResourceLoader实例
// 	resourceLoader := &config.FileResourceLoader{}

// 	// 创建代理实例
// 	promptService := config.NewPromptServiceProxy(resourceLoader)

// 	// 调用代理方法
// 	fmt.Println(promptService.GetPrompt())

// }PromptConfig.go测试
// package main

// import (
// 	"drawsee/config"
// 	"fmt"
// )

// func main() {
// 	// 初始化连接工厂
// 	connectionFactory := &config.ConnectionFactory{}

// 	// 初始化RabbitConfig
// 	rabbitConfig := config.RabbitConfig{
// 		AIConfig: config.MqConfig{
// 			QueueName:        "ai_queue",
// 			RoutingKey:       "ai_routing_key",
// 			QueueCount:       3,
// 			QueueConcurrency: 1,
// 			ExchangeName:     "ai_exchange",
// 		},
// 		AnimationConfig: config.MqConfig{
// 			QueueName:        "animation_queue",
// 			RoutingKey:       "animation_routing_key",
// 			QueueCount:       2,
// 			QueueConcurrency: 1,
// 			ExchangeName:     "animation_exchange",
// 		},
// 	}

// 	// 创建RabbitAdmin
// 	admin := rabbitConfig.NewRabbitAdmin(connectionFactory)

// 	// 获取AI任务队列
// 	aiTaskQueues := rabbitConfig.GetLinkedQueues(rabbitConfig.AIConfig)
// 	// 声明AI任务队列
// 	rabbitConfig.DeclareExchangeAndQueues(aiTaskQueues, admin, rabbitConfig.AIConfig)

// 	// 获取动画任务队列
// 	animationTaskQueues := rabbitConfig.GetLinkedQueues(rabbitConfig.AnimationConfig)
// 	// 声明动画任务队列
// 	rabbitConfig.DeclareExchangeAndQueues(animationTaskQueues, admin, rabbitConfig.AnimationConfig)

//		fmt.Println("RabbitMQ configuration completed.")
//	}RabiitMQConfig.go测试
// package main

// import (
// 	"drawsee/config"
// 	"drawsee/worker"
// 	"fmt"
// 	"log"
// )

// func main() {
// 	// 初始化 AI 任务队列
// 	aiTaskQueues := []config.LinkedQueue{
// 		{Name: "ai_queue_1", Concurrency: 1},
// 		{Name: "ai_queue_2", Concurrency: 1},
// 	}

// 	// 初始化消息转换器
// 	jsonMessageConverter := &config.JsonMessageConverter{}

// 	// 初始化 AI 任务工作器
// 	aiTaskWorker := &worker.AITaskWorker{}

// 	// 初始化监听端点注册器
// 	registrar := &config.RabbitListenerEndpointRegistrar{}

// 	// 配置 RabbitMQ 监听器
// 	config.ConfigureRabbitListeners(aiTaskQueues, jsonMessageConverter, aiTaskWorker, registrar)

// 	// 模拟消息处理
// 	for _, endpoint := range registrar.Endpoints {
// 		log.Printf("Endpoint registered: %s", endpoint.Id)
// 		// 模拟消息接收
// 		message := []byte(`{"taskId":"task123"}`)
// 		endpoint.MessageListener(message)
// 	}

//		fmt.Println("RabbitMQ listener configuration completed.")
//	}RabbitMQListener.go测试
// package main

// import (
// 	"drawsee/config"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	mux := config.SetupRoutes()

//		log.Println("Starting server at :8080")
//		if err := http.ListenAndServe(":8080", mux); err != nil {
//			log.Fatalf("Error starting server: %v", err)
//		}
//	}RestClient.go测试
package main

import (
	"encoding/json"
	"fmt"

	"drawsee/pojo/dto"
	"drawsee/pojo/mongo"
)

func main() {
	// 创建一个 KnowledgeResource 实例
	res := mongo.KnowledgeResource{
		Type:  "text",
		Value: "This is a sample knowledge resource",
	}

	// 创建一个 AddKnowledgeDTO 实例
	addKnowledge := dto.AddKnowledgeDTO{
		Name:      "Sample Knowledge",
		Aliases:   []string{"alias1", "alias2"},
		Resources: []mongo.KnowledgeResource{res},
		ParentId:  "12345",
	}

	// 将结构体序列化为 JSON
	jsonData, err := json.Marshal(addKnowledge)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// 打印 JSON 数据
	fmt.Println("JSON data:", string(jsonData))

	// 模拟出错
	if err := simulateError(); err != nil {
		fmt.Println("An error occurred:", err)
	}
}

// simulateError 模拟一个错误
func simulateError() error {
	return fmt.Errorf("something went wrong")
}

//APIConfig.go测试
