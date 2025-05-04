package mapper

import (
	"drawsee/pojo/entity"
)

// InvitationCodeMapper defines the data operation interface related to invitation codes
type InvitationCodeMapper interface {
	GetById(id int64) (*entity.InvitationCode, error)
	GetByCode(code string) (*entity.InvitationCode, error)
	GetByUsedBy(usedBy int64) (*entity.InvitationCode, error)
	Insert(invitationCode *entity.InvitationCode) error
}

// MockInvitationCodeMapper simulates the implementation of InvitationCodeMapper for testing
type MockInvitationCodeMapper struct {
	codes     map[int64]*entity.InvitationCode
	codeMap   map[string]*entity.InvitationCode
	usedByMap map[int64]*entity.InvitationCode
}

// NewMockInvitationCodeMapper creates a new MockInvitationCodeMapper instance
func NewMockInvitationCodeMapper() *MockInvitationCodeMapper {
	return &MockInvitationCodeMapper{
		codes:     make(map[int64]*entity.InvitationCode),
		codeMap:   make(map[string]*entity.InvitationCode),
		usedByMap: make(map[int64]*entity.InvitationCode),
	}
}

// GetById simulates getting an invitation code by ID
func (m *MockInvitationCodeMapper) GetById(id int64) (*entity.InvitationCode, error) {
	code, exists := m.codes[id]
	if !exists {
		return nil, nil
	}
	return code, nil
}

// GetByCode simulates getting an invitation code by code string
func (m *MockInvitationCodeMapper) GetByCode(code string) (*entity.InvitationCode, error) {
	invitationCode, exists := m.codeMap[code]
	if !exists {
		return nil, nil
	}
	return invitationCode, nil
}

// GetByUsedBy simulates getting an invitation code by the user who used it
func (m *MockInvitationCodeMapper) GetByUsedBy(usedBy int64) (*entity.InvitationCode, error) {
	code, exists := m.usedByMap[usedBy]
	if !exists {
		return nil, nil
	}
	return code, nil
}

// Insert simulates inserting an invitation code
func (m *MockInvitationCodeMapper) Insert(invitationCode *entity.InvitationCode) error {
	if invitationCode.ID != nil {
		m.codes[*invitationCode.ID] = invitationCode
	}
	if invitationCode.Code != "" {
		m.codeMap[invitationCode.Code] = invitationCode
	}
	if invitationCode.UsedBy != nil {
		m.usedByMap[*invitationCode.UsedBy] = invitationCode
	}
	return nil
}
