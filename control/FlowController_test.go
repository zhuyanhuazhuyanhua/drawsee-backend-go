package control

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockFlowService is a mock implementation of FlowService
type MockFlowService struct {
}

func (m *MockFlowService) GetConversations() ([]ConversationVO, error) {
	return []ConversationVO{}, nil
}

func (m *MockFlowService) DeleteConversation(convId int64) error {
	return nil
}

func (m *MockFlowService) GetNodes(convId int64) ([]NodeVO, error) {
	return []NodeVO{}, nil
}

func (m *MockFlowService) UpdateNodes(updateNodesDTO UpdateNodesDTO) error {
	return nil
}

func (m *MockFlowService) UpdateNode(nodeId int64, updateNodeDTO UpdateNodeDTO) error {
	return nil
}

func (m *MockFlowService) DeleteNode(nodeId int64) error {
	return nil
}

func (m *MockFlowService) GetProcessingTasks(convId int64) ([]AiTaskVO, error) {
	return []AiTaskVO{}, nil
}

func (m *MockFlowService) CreateTask(createAiTaskDTO CreateAiTaskDTO) (CreateAiTaskVO, error) {
	return CreateAiTaskVO{}, nil
}

func (m *MockFlowService) GetCompletion(taskId int64) (*SseEmitter, error) {
	return &SseEmitter{}, nil
}

func (m *MockFlowService) GetResource(objectName string) (ResourceVO, error) {
	return ResourceVO{}, nil
}

func TestFlowController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := &MockFlowService{}
	fc := &FlowController{FlowService: mockService}
	fc.RegisterRoutes(router)

	// Test GET /flow/conversations
	req, _ := http.NewRequest("GET", "/flow/conversations", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	// Test DELETE /flow/conversations/:convId
	req, _ = http.NewRequest("DELETE", "/flow/conversations/1", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	// Test GET /flow/nodes
	req, _ = http.NewRequest("GET", "/flow/nodes?convId=1", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	// Test POST /flow/nodes
	updateNodesDTO := UpdateNodesDTO{}
	jsonData, _ := json.Marshal(updateNodesDTO)
	req, _ = http.NewRequest("POST", "/flow/nodes", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	// Test POST /flow/nodes/:nodeId
	updateNodeDTO := UpdateNodeDTO{}
	jsonData, _ = json.Marshal(updateNodeDTO)
	req, _ = http.NewRequest("POST", "/flow/nodes/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	// Test DELETE /flow/nodes/:nodeId
	req, _ = http.NewRequest("DELETE", "/flow/nodes/1", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	// Test GET /flow/tasks
	req, _ = http.NewRequest("GET", "/flow/tasks?convId=1", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	// Test POST /flow/tasks
	createAiTaskDTO := CreateAiTaskDTO{
		// 填充一些示例数据
		Name: "Test Task",
	}
	jsonData, _ = json.Marshal(createAiTaskDTO)
	req, _ = http.NewRequest("POST", "/flow/tasks", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	// Test GET /flow/completion
	req, _ = http.NewRequest("GET", "/flow/completion?taskId=1", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	// Test GET /flow/resources
	req, _ = http.NewRequest("GET", "/flow/resources?objectName=test", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
