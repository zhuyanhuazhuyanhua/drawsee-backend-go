package config

import (
	"reflect"
	"testing"

	"drawsee/worker"
)

func TestConfigureRabbitListeners(t *testing.T) {
	// 初始化 AI 任务队列
	aiTaskQueues := []LinkedQueue{
		{Name: "ai_queue_1", Concurrency: 1},
		{Name: "ai_queue_2", Concurrency: 1},
	}

	// 初始化消息转换器
	jsonMessageConverter := &JsonMessageConverter{}

	// 初始化 AI 任务工作器
	aiTaskWorker := &worker.AITaskWorker{}

	// 初始化监听端点注册器
	registrar := &RabbitListenerEndpointRegistrar{}

	// 配置 RabbitMQ 监听器
	ConfigureRabbitListeners(aiTaskQueues, jsonMessageConverter, aiTaskWorker, registrar)

	// 验证端点是否正确注册
	expectedEndpoints := []RabbitListenerEndpoint{
		{
			Id:          "ai_queue_1.Endpoint",
			QueueNames:  []string{"ai_queue_1"},
			Concurrency: 1,
		},
		{
			Id:          "ai_queue_2.Endpoint",
			QueueNames:  []string{"ai_queue_2"},
			Concurrency: 1,
		},
	}

	if len(registrar.Endpoints) != len(expectedEndpoints) {
		t.Errorf("Expected %d endpoints, got %d", len(expectedEndpoints), len(registrar.Endpoints))
	}

	for i, endpoint := range registrar.Endpoints {
		if endpoint.Id != expectedEndpoints[i].Id ||
			!reflect.DeepEqual(endpoint.QueueNames, expectedEndpoints[i].QueueNames) ||
			endpoint.Concurrency != expectedEndpoints[i].Concurrency {
			t.Errorf("Endpoint mismatch: expected %+v, got %+v", expectedEndpoints[i], endpoint)
		}
	}
}
