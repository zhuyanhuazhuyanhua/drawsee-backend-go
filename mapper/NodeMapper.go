package mapper

import (
	"drawsee/pojo/entity"
)

// NodeMapper 定义节点相关的数据操作接口
type NodeMapper interface {
	GetById(id int64) (*entity.Node, error)
	GetByConvId(convId int64) ([]*entity.Node, error)
	Insert(node *entity.Node) error
	Update(node *entity.Node) error
	UpdateDataAndIsDeletedBatch(nodes []*entity.Node) error
	UpdatePositionAndHeightBatch(nodes []*entity.Node) error
}

// MockNodeMapper 模拟 NodeMapper 实现，用于测试
type MockNodeMapper struct {
	nodes map[int64]*entity.Node
}

// NewMockNodeMapper 创建一个新的 MockNodeMapper 实例
func NewMockNodeMapper() *MockNodeMapper {
	return &MockNodeMapper{
		nodes: make(map[int64]*entity.Node),
	}
}

// GetById 模拟根据 ID 获取节点
func (m *MockNodeMapper) GetById(id int64) (*entity.Node, error) {
	node, exists := m.nodes[id]
	if !exists {
		return nil, nil
	}
	return node, nil
}

// GetByConvId 模拟根据会话 ID 获取节点列表
func (m *MockNodeMapper) GetByConvId(convId int64) ([]*entity.Node, error) {
	var result []*entity.Node
	for _, node := range m.nodes {
		if node.ConvID != nil && *node.ConvID == convId {
			result = append(result, node)
		}
	}
	return result, nil
}

// Insert 模拟插入节点
func (m *MockNodeMapper) Insert(node *entity.Node) error {
	if node.ID != nil {
		m.nodes[*node.ID] = node
	}
	return nil
}

// Update 模拟更新节点
func (m *MockNodeMapper) Update(node *entity.Node) error {
	if node.ID != nil {
		m.nodes[*node.ID] = node
	}
	return nil
}

// UpdateDataAndIsDeletedBatch 模拟批量更新节点的数据和删除状态
func (m *MockNodeMapper) UpdateDataAndIsDeletedBatch(nodes []*entity.Node) error {
	for _, node := range nodes {
		if node.ID != nil {
			if existingNode, exists := m.nodes[*node.ID]; exists {
				existingNode.Data = node.Data
				existingNode.IsDeleted = node.IsDeleted
			}
		}
	}
	return nil
}

// UpdatePositionAndHeightBatch 模拟批量更新节点的位置和高度
func (m *MockNodeMapper) UpdatePositionAndHeightBatch(nodes []*entity.Node) error {
	for _, node := range nodes {
		if node.ID != nil {
			if existingNode, exists := m.nodes[*node.ID]; exists {
				existingNode.Position = node.Position
				existingNode.Height = node.Height
			}
		}
	}
	return nil
}
