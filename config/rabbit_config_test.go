package config

import (
	"testing"
)

func TestGetLinkedQueues(t *testing.T) {
	mqConfig := MqConfig{
		QueueName:        "test_queue",
		RoutingKey:       "test_routing_key",
		QueueCount:       2,
		QueueConcurrency: 1,
		ExchangeName:     "test_exchange",
	}

	rabbitConfig := RabbitConfig{}
	queues := rabbitConfig.GetLinkedQueues(mqConfig)

	expectedQueues := []LinkedQueue{
		{Name: "test_queue_1", RoutingKey: "test_routing_key_1", Concurrency: 1, ExchangeName: "test_exchange"},
		{Name: "test_queue_2", RoutingKey: "test_routing_key_2", Concurrency: 1, ExchangeName: "test_exchange"},
	}

	if len(queues) != len(expectedQueues) {
		t.Errorf("Expected %d queues, got %d", len(expectedQueues), len(queues))
	}

	for i, queue := range queues {
		if queue.Name != expectedQueues[i].Name ||
			queue.RoutingKey != expectedQueues[i].RoutingKey ||
			queue.Concurrency != expectedQueues[i].Concurrency ||
			queue.ExchangeName != expectedQueues[i].ExchangeName {
			t.Errorf("Queue mismatch: expected %+v, got %+v", expectedQueues[i], queue)
		}
	}
}

func TestDeclareExchangeAndQueues(t *testing.T) {
	mqConfig := MqConfig{
		QueueName:        "test_queue",
		RoutingKey:       "test_routing_key",
		QueueCount:       2,
		QueueConcurrency: 1,
		ExchangeName:     "test_exchange",
	}

	rabbitConfig := RabbitConfig{}
	queues := rabbitConfig.GetLinkedQueues(mqConfig)

	admin := &RabbitAdmin{connectionFactory: &ConnectionFactory{}}

	rabbitConfig.DeclareExchangeAndQueues(queues, admin, mqConfig)
}
