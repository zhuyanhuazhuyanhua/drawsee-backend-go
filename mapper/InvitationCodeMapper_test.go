package mapper_test

import (
	"drawsee/mapper"
	"drawsee/pojo/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockInvitationCodeMapper(t *testing.T) {
	mockMapper := mapper.NewMockInvitationCodeMapper()

	// Prepare test data
	id := int64(1)
	code := "TEST123"
	usedBy := int64(100)

	invitationCode := &entity.InvitationCode{
		ID:     &id,
		Code:   code,
		UsedBy: &usedBy,
	}

	// Test Insert
	err := mockMapper.Insert(invitationCode)
	assert.NoError(t, err)

	// Test GetById
	foundCode, err := mockMapper.GetById(id)
	assert.NoError(t, err)
	assert.Equal(t, invitationCode, foundCode)

	// Test GetByCode
	foundByCode, err := mockMapper.GetByCode(code)
	assert.NoError(t, err)
	assert.Equal(t, invitationCode, foundByCode)

	// Test GetByUsedBy
	foundByUsedBy, err := mockMapper.GetByUsedBy(usedBy)
	assert.NoError(t, err)
	assert.Equal(t, invitationCode, foundByUsedBy)
}
