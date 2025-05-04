# Annotation 模块文档

## 概述
`annotation` 目录包含了项目中与注解相关的代码文件，这些文件主要用于定义和处理注解逻辑，辅助实现项目的一些特定功能，如获取提示资源、提取提示参数等。

## 文件说明

### 核心代码文件
| 文件名              | 大小     | 功能描述                                                                                                           |
| ------------------- | -------- | ------------------------------------------------------------------------------------------------------------------ |
| `PromptParam.go`    | 723 字节 | 定义 `PromptParam` 结构体和获取提示参数值的函数，用于从结构体中提取带有 `promptparam` 标签的字段值。               |
| `PromptResource.go` | 641 字节 | 定义 `PromptResource` 结构体和获取提示资源的函数，通过反射检查方法参数，判断是否存在 `PromptResource` 类型的参数。 |
| `ValuSet.go`        | 495 字节 | 定义 `ValueSet` 结构体和创建 `ValueSet` 实例的函数，用于存储消息、分组、负载和值等信息。                           |

### 测试文件
| 文件名                  | 大小      | 功能描述                                                                       |
| ----------------------- | --------- | ------------------------------------------------------------------------------ |
| `PromptParam_test.go`   | 1110 字节 | 包含 `PromptParam.go` 中函数的测试用例，验证提示参数值的提取逻辑。             |
| `PromptResorce_test.go` | 1103 字节 | 对 `PromptResource.go` 中获取提示资源函数进行测试，确保函数按预期工作。        |
| `ValueSet_test.go`      | 1082 字节 | 对 `ValuSet.go` 中相关功能进行测试，保证 `ValueSet` 实例的创建和使用逻辑正确。 |

## 使用示例

### 获取提示参数值
```go
package main

import (
	"fmt"
	"drawsee/annotation"
)

type TestStruct struct {
	Param string `promptparam:"test_value"`
}

func main() {
	test := &TestStruct{}
	value, err := annotation.GetPromptParamValue(test)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("PromptParam value:", value)
}