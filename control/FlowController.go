package control

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FlowController struct
type FlowController struct {
	FlowService FlowService
}

// FlowService interface
type FlowService interface {
	GetConversations() ([]ConversationVO, error)
	DeleteConversation(convId int64) error
	GetNodes(convId int64) ([]NodeVO, error)
	UpdateNodes(updateNodesDTO UpdateNodesDTO) error
	UpdateNode(nodeId int64, updateNodeDTO UpdateNodeDTO) error
	DeleteNode(nodeId int64) error
	GetProcessingTasks(convId int64) ([]AiTaskVO, error)
	CreateTask(createAiTaskDTO CreateAiTaskDTO) (CreateAiTaskVO, error)
	GetCompletion(taskId int64) (*SseEmitter, error)
	GetResource(objectName string) (ResourceVO, error)
}

// SseEmitter struct
type SseEmitter struct {
	Timeout int64
}

// ConversationVO struct
type ConversationVO struct {
	// Define fields here
}

// NodeVO struct
type NodeVO struct {
	// Define fields here
}

// UpdateNodesDTO struct
type UpdateNodesDTO struct {
	// Define fields here
}

// UpdateNodeDTO struct
type UpdateNodeDTO struct {
	// Define fields here
}

// AiTaskVO struct
type AiTaskVO struct {
	// Define fields here
}

// CreateAiTaskDTO struct
type CreateAiTaskDTO struct {
	Name string `json:"name"` // 确保字段名和标签正确测试用
}

// CreateAiTaskVO struct
type CreateAiTaskVO struct {
	// Define fields here
}

// ResourceVO struct
type ResourceVO struct {
	// Define fields here
}

// RegisterRoutes registers the routes for FlowController
func (fc *FlowController) RegisterRoutes(router *gin.Engine) {
	flowGroup := router.Group("/flow")
	{
		flowGroup.GET("/conversations", fc.getConversations)
		flowGroup.DELETE("/conversations/:convId", fc.deleteConversation)
		flowGroup.GET("/nodes", fc.getNodes)
		flowGroup.POST("/nodes", fc.updateNodes)
		flowGroup.POST("/nodes/:nodeId", fc.updateNode)
		flowGroup.DELETE("/nodes/:nodeId", fc.deleteNode)
		flowGroup.GET("/tasks", fc.getProcessingTasks)
		flowGroup.POST("/tasks", fc.createTask)
		flowGroup.GET("/completion", fc.getCompletion)
		flowGroup.GET("/resources", fc.getResource)
	}
}

// getConversations handles GET /flow/conversations
func (fc *FlowController) getConversations(c *gin.Context) {
	conversations, err := fc.FlowService.GetConversations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, conversations)
}

// deleteConversation handles DELETE /flow/conversations/:convId
func (fc *FlowController) deleteConversation(c *gin.Context) {
	convId, err := strconv.ParseInt(c.Param("convId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid convId"})
		return
	}
	err = fc.FlowService.DeleteConversation(convId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// getNodes handles GET /flow/nodes
func (fc *FlowController) getNodes(c *gin.Context) {
	convId, err := strconv.ParseInt(c.Query("convId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid convId"})
		return
	}
	nodes, err := fc.FlowService.GetNodes(convId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nodes)
}

// updateNodes handles POST /flow/nodes
func (fc *FlowController) updateNodes(c *gin.Context) {
	var updateNodesDTO UpdateNodesDTO
	if err := c.ShouldBindJSON(&updateNodesDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := fc.FlowService.UpdateNodes(updateNodesDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// updateNode handles POST /flow/nodes/:nodeId
func (fc *FlowController) updateNode(c *gin.Context) {
	nodeId, err := strconv.ParseInt(c.Param("nodeId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nodeId"})
		return
	}
	var updateNodeDTO UpdateNodeDTO
	if err := c.ShouldBindJSON(&updateNodeDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = fc.FlowService.UpdateNode(nodeId, updateNodeDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// deleteNode handles DELETE /flow/nodes/:nodeId
func (fc *FlowController) deleteNode(c *gin.Context) {
	nodeId, err := strconv.ParseInt(c.Param("nodeId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nodeId"})
		return
	}
	err = fc.FlowService.DeleteNode(nodeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// getProcessingTasks handles GET /flow/tasks
func (fc *FlowController) getProcessingTasks(c *gin.Context) {
	convId, err := strconv.ParseInt(c.Query("convId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid convId"})
		return
	}
	tasks, err := fc.FlowService.GetProcessingTasks(convId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// createTask handles POST /flow/tasks
func (fc *FlowController) createTask(c *gin.Context) {
	var createAiTaskDTO CreateAiTaskDTO
	if err := c.ShouldBindJSON(&createAiTaskDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := fc.FlowService.CreateTask(createAiTaskDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// getCompletion handles GET /flow/completion
func (fc *FlowController) getCompletion(c *gin.Context) {
	taskId, err := strconv.ParseInt(c.Query("taskId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid taskId"})
		return
	}
	emitter, err := fc.FlowService.GetCompletion(taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Handle SseEmitter here
	c.JSON(http.StatusOK, emitter)
}

// getResource handles GET /flow/resources
func (fc *FlowController) getResource(c *gin.Context) {
	objectName := c.Query("objectName")
	if objectName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "objectName is required"})
		return
	}
	resource, err := fc.FlowService.GetResource(objectName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resource)
}
