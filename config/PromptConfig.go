package config

import (
	"fmt"
	"os"
	"reflect"
)

// PromptService 定义了PromptService接口
type PromptService interface {
	GetPrompt() string
}

// PromptServiceImpl 是PromptService的具体实现
type PromptServiceImpl struct {
	ResourceLoader ResourceLoader
}

// GetPrompt 实现了PromptService接口
func (p *PromptServiceImpl) GetPrompt() string {
	return "This is a prompt from the service."
}

// ResourceLoader 定义了ResourceLoader接口
type ResourceLoader interface {
	GetResource(path string) (string, error)
}

// FileResourceLoader 是ResourceLoader的具体实现
type FileResourceLoader struct{}

// GetResource 实现了ResourceLoader接口
func (f *FileResourceLoader) GetResource(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

// PromptServiceInvocationHandler 是一个代理处理器
type PromptServiceInvocationHandler struct {
	RealSubject PromptService
}

// NewPromptServiceInvocationHandler 创建一个新的代理处理器
func NewPromptServiceInvocationHandler(realSubject PromptService) *PromptServiceInvocationHandler {
	return &PromptServiceInvocationHandler{RealSubject: realSubject}
}

// Invoke 处理方法调用
func (h *PromptServiceInvocationHandler) Invoke(method reflect.Method, args []reflect.Value) (result []reflect.Value) {
	fmt.Println("Proxy: Logging the time of request.")
	realSubjectValue := reflect.ValueOf(h.RealSubject)
	return realSubjectValue.MethodByName(method.Name).Call(args)
}

// NewPromptServiceProxy 创建一个代理对象
func NewPromptServiceProxy(resourceLoader ResourceLoader) PromptService {
	realSubject := &PromptServiceImpl{ResourceLoader: resourceLoader}
	proxy := &PromptServiceInvocationHandler{RealSubject: realSubject}
	return proxy
}

// GetPromptProxy 实现了PromptService接口
func (p *PromptServiceInvocationHandler) GetPrompt() string {
	fmt.Println("Proxy: Logging the time of request.")
	return p.RealSubject.GetPrompt()
}
