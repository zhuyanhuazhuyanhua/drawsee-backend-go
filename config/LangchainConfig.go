package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// AiModelConfig 用于映射单个模型的配置
type AiModelConfig struct {
	BaseURL   string `mapstructure:"base_url"`   // 对应 YAML 文件中的 base_url
	APIKey    string `mapstructure:"api_key"`    // 对应 YAML 文件中的 api_key
	ModelName string `mapstructure:"model_name"` // 对应 YAML 文件中的 model_name
}

// LangchainConfig 用于映射整个配置文件
type LangchainConfig struct {
	Doubao       AiModelConfig `mapstructure:"doubao"`       // 对应 YAML 文件中的 doubao
	DeepseekV3   AiModelConfig `mapstructure:"deepseekV3"`   // 对应 YAML 文件中的 deepseekV3
	DoubaoVision AiModelConfig `mapstructure:"doubaoVision"` // 对应 YAML 文件中的 doubaoVision
}

// LoadConfig 用于加载配置文件
func LoadConfig(path string) (*LangchainConfig, error) {
	fmt.Printf("开始加载配置文件，路径: %s\n", path)
	viper.SetConfigFile(path)   // 设置配置文件路径
	viper.SetConfigType("yaml") // 设置配置文件类型为 YAML

	if err := viper.ReadInConfig(); err != nil { // 读取配置文件
		fmt.Printf("读取配置文件失败，错误信息: %v\n", err)
		return nil, err
	}
	fmt.Println("配置文件读取成功")

	// 打印 viper 中的所有配置项
	fmt.Println("Viper 中的配置项:")
	for _, key := range viper.AllKeys() {
		fmt.Printf("%s: %v\n", key, viper.Get(key))
	}

	var config LangchainConfig
	// 使用 viper.Sub 方法获取 drawsee.models 子配置
	modelsConfig := viper.Sub("drawsee.models")
	if modelsConfig == nil {
		fmt.Println("未找到 drawsee.models 配置")
		return nil, fmt.Errorf("未找到 drawsee.models 配置")
	}

	if err := modelsConfig.Unmarshal(&config); err != nil { // 将配置文件内容映射到结构体
		fmt.Printf("配置内容映射到结构体失败，错误信息: %v\n", err)
		return nil, err
	}
	fmt.Println("配置内容成功映射到结构体")

	// 打印映射后的结构体内容
	fmt.Println("映射后的配置结构体内容:")
	fmt.Printf("Doubao: BaseURL=%s, APIKey=%s, ModelName=%s\n", config.Doubao.BaseURL, config.Doubao.APIKey, config.Doubao.ModelName)
	fmt.Printf("DeepseekV3: BaseURL=%s, APIKey=%s, ModelName=%s\n", config.DeepseekV3.BaseURL, config.DeepseekV3.APIKey, config.DeepseekV3.ModelName)
	fmt.Printf("DoubaoVision: BaseURL=%s, APIKey=%s, ModelName=%s\n", config.DoubaoVision.BaseURL, config.DoubaoVision.APIKey, config.DoubaoVision.ModelName)

	return &config, nil // 返回配置结构体
}
