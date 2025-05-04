package dto_test

import (
	"encoding/json"
	"testing"

	"drawsee/pojo/dto"
	"drawsee/pojo/mongo"
)

func TestAddKnowledgeDTO(t *testing.T) {
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
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	expected := `{"name":"Sample Knowledge","aliases":["alias1","alias2"],"resources":[{"type":"text","value":"This is a sample knowledge resource"}],"parentId":"12345"}`
	if string(jsonData) != expected {
		t.Errorf("Expected JSON: %s, got: %s", expected, string(jsonData))
	}
}
