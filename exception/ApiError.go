package exception

import (
	"fmt"
	"strings"
)

// ApiError 定义了API错误类型
type ApiError struct {
	Code    int
	Message string
}

// 定义错误常量
const (
	// admin
	NotAdmin        = "not_admin"
	AdminHadExisted = "admin_had_existed"

	// invitation_code
	InvitationCodeNotExist    = "invitation_code_not_existed"
	InvitationCodeSendFail    = "invitation_code_send_failed"
	InvitationCodeAlreadyUsed = "invitation_code_already_used"

	// file
	FileTypeNotSupported = "file_type_not_supported"
	FileUploadFail       = "file_upload_failed"
	ImageRecognizeFail   = "image_recognize_failed"

	// node
	NodeNotExist = "node_not_existed"

	// ai_task
	AITaskNotExist    = "ai_task_not_existed"
	AITaskIsWaiting   = "ai_task_is_waiting"
	AITaskIsFinished  = "ai_task_is_finished"
	AITaskExceedLimit = "ai_task_exceed_limit"

	// conversation
	ConversationNotExist = "conversation_not_existed"

	// knowledge
	KnowledgeHadExisted     = "knowledge_had_existed"
	KnowledgeNotExist       = "knowledge_not_existed"
	KnowledgeParentNotExist = "knowledge_parent_not_existed"

	// 参数
	ParamError = "param_error"

	// 系统
	SystemError = "system_error"

	// 权限
	NotLogin = "not_login"

	// 登录
	UserNotExist  = "user_not_exist"
	PasswordError = "password_error"

	// 注册
	UserHadExisted = "user_had_existed"
)

// GetApiError 根据错误名称获取对应的ApiError
func GetApiError(name string) (ApiError, error) {
	switch strings.ToLower(name) {
	case NotAdmin:
		return ApiError{Code: 401, Message: "不是管理员"}, nil
	case AdminHadExisted:
		return ApiError{Code: 409, Message: "管理员已存在"}, nil
	case InvitationCodeNotExist:
		return ApiError{Code: 404, Message: "邀请码不存在"}, nil
	case InvitationCodeSendFail:
		return ApiError{Code: 500, Message: "邀请码发送失败"}, nil
	case InvitationCodeAlreadyUsed:
		return ApiError{Code: 409, Message: "邀请码已使用"}, nil
	case FileTypeNotSupported:
		return ApiError{Code: 400, Message: "文件类型不支持"}, nil
	case FileUploadFail:
		return ApiError{Code: 500, Message: "文件上传失败"}, nil
	case ImageRecognizeFail:
		return ApiError{Code: 500, Message: "图片识别失败"}, nil
	case NodeNotExist:
		return ApiError{Code: 404, Message: "节点不存在"}, nil
	case AITaskNotExist:
		return ApiError{Code: 404, Message: "not_exist"}, nil
	case AITaskIsWaiting:
		return ApiError{Code: 409, Message: "waiting"}, nil
	case AITaskIsFinished:
		return ApiError{Code: 409, Message: "finished"}, nil
	case AITaskExceedLimit:
		return ApiError{Code: 409, Message: "24小时内您的对话次数已达上限，请明天再试"}, nil
	case ConversationNotExist:
		return ApiError{Code: 404, Message: "会话不存在"}, nil
	case KnowledgeHadExisted:
		return ApiError{Code: 409, Message: "同名称知识点已存在"}, nil
	case KnowledgeNotExist:
		return ApiError{Code: 404, Message: "知识点不存在"}, nil
	case KnowledgeParentNotExist:
		return ApiError{Code: 404, Message: "父知识点不存在"}, nil
	case ParamError:
		return ApiError{Code: 400, Message: "参数错误"}, nil
	case SystemError:
		return ApiError{Code: 500, Message: "服务器内部错误"}, nil
	case NotLogin:
		return ApiError{Code: 401, Message: "未登录"}, nil
	case UserNotExist:
		return ApiError{Code: 404, Message: "用户名不存在"}, nil
	case PasswordError:
		return ApiError{Code: 401, Message: "密码错误"}, nil
	case UserHadExisted:
		return ApiError{Code: 409, Message: "用户名已存在"}, nil
	default:
		return ApiError{}, fmt.Errorf("unknown error name: %s", name)
	}
}

// Error 实现error接口
func (e ApiError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

// NotLoginError 未登录错误类型
type NotLoginError struct{}

func (e NotLoginError) Error() string {
	apiError, _ := GetApiError(NotLogin)
	return apiError.Error()
}
