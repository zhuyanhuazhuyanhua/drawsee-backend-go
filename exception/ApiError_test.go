package exception

import (
	"testing"
)

func TestGetApiError(t *testing.T) {
	tests := []struct {
		name     string
		expected ApiError
	}{
		{NotAdmin, ApiError{Code: 401, Message: "不是管理员"}},
		{AdminHadExisted, ApiError{Code: 409, Message: "管理员已存在"}},
		{InvitationCodeNotExist, ApiError{Code: 404, Message: "邀请码不存在"}},
		{InvitationCodeSendFail, ApiError{Code: 500, Message: "邀请码发送失败"}},
		{InvitationCodeAlreadyUsed, ApiError{Code: 409, Message: "邀请码已使用"}},
		{FileTypeNotSupported, ApiError{Code: 400, Message: "文件类型不支持"}},
		{FileUploadFail, ApiError{Code: 500, Message: "文件上传失败"}},
		{ImageRecognizeFail, ApiError{Code: 500, Message: "图片识别失败"}},
		{NodeNotExist, ApiError{Code: 404, Message: "节点不存在"}},
		{AITaskNotExist, ApiError{Code: 404, Message: "not_exist"}},
		{AITaskIsWaiting, ApiError{Code: 409, Message: "waiting"}},
		{AITaskIsFinished, ApiError{Code: 409, Message: "finished"}},
		{AITaskExceedLimit, ApiError{Code: 409, Message: "24小时内您的对话次数已达上限，请明天再试"}},
		{ConversationNotExist, ApiError{Code: 404, Message: "会话不存在"}},
		{KnowledgeHadExisted, ApiError{Code: 409, Message: "同名称知识点已存在"}},
		{KnowledgeNotExist, ApiError{Code: 404, Message: "知识点不存在"}},
		{KnowledgeParentNotExist, ApiError{Code: 404, Message: "父知识点不存在"}},
		{ParamError, ApiError{Code: 400, Message: "参数错误"}},
		{SystemError, ApiError{Code: 500, Message: "服务器内部错误"}},
		{NotLogin, ApiError{Code: 401, Message: "未登录"}},
		{UserNotExist, ApiError{Code: 404, Message: "用户名不存在"}},
		{PasswordError, ApiError{Code: 401, Message: "密码错误"}},
		{UserHadExisted, ApiError{Code: 409, Message: "用户名已存在"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := GetApiError(tt.name)
			if err != nil {
				t.Fatalf("GetApiError(%s) failed: %v", tt.name, err)
			}
			if actual.Code != tt.expected.Code || actual.Message != tt.expected.Message {
				t.Errorf("GetApiError(%s) = %v, want %v", tt.name, actual, tt.expected)
			}
		})
	}

	// 测试未知错误名称
	_, err := GetApiError("unknown_error")
	if err == nil {
		t.Errorf("GetApiError(unknown_error) should return an error")
	}
}
