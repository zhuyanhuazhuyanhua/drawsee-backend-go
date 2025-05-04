package mapper_test

import (
	"drawsee/mapper"
	"drawsee/pojo/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockNodeMapper(t *testing.T) {
	mockMapper := mapper.NewMockNodeMapper()

	// 准备测试数据
	id := int64(1)
	convId := int64(100)
	nodeType := "test_type"
	data := "test_data"
	position := "test_position"
	height := int64(10)
	parentId := int64(2)
	userId := int64(3)
	isDeleted := false

	node := &entity.Node{
		ID:        &id,
		Type:      nodeType,
		Data:      data,
		Position:  position,
		Height:    &height,
		ParentID:  &parentId,
		ConvID:    &convId,
		UserID:    &userId,
		IsDeleted: &isDeleted,
	}

	// 测试 Insert
	err := mockMapper.Insert(node)
	assert.NoError(t, err)

	// 测试 GetById
	foundNode, err := mockMapper.GetById(id)
	assert.NoError(t, err)
	assert.Equal(t, node, foundNode)

	// 测试 GetByConvId
	nodes, err := mockMapper.GetByConvId(convId)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(nodes))

	// 测试 Update
	newData := "updated_data"
	node.Data = newData
	err = mockMapper.Update(node)
	assert.NoError(t, err)
	updatedNode, err := mockMapper.GetById(id)
	assert.NoError(t, err)
	assert.Equal(t, newData, updatedNode.Data)

	// 测试 UpdateDataAndIsDeletedBatch
	newIsDeleted := true
	node.IsDeleted = &newIsDeleted
	err = mockMapper.UpdateDataAndIsDeletedBatch([]*entity.Node{node})
	assert.NoError(t, err)
	updatedBatchNode, err := mockMapper.GetById(id)
	assert.NoError(t, err)
	assert.Equal(t, newIsDeleted, *updatedBatchNode.IsDeleted)

	// 测试 UpdatePositionAndHeightBatch
	newPosition := "updated_position"
	newHeight := int64(20)
	node.Position = newPosition
	node.Height = &newHeight
	err = mockMapper.UpdatePositionAndHeightBatch([]*entity.Node{node})
	assert.NoError(t, err)
	updatedHeightNode, err := mockMapper.GetById(id)
	assert.NoError(t, err)
	assert.Equal(t, newPosition, updatedHeightNode.Position)
	assert.Equal(t, newHeight, *updatedHeightNode.Height)
}
