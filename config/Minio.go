package config

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

// MinioConfig 用于映射 MinIO 配置
type MinioConfig struct {
	Endpoint   string `mapstructure:"endpoint"`
	AccessKey  string `mapstructure:"accessKey"`
	SecretKey  string `mapstructure:"secretKey"`
	BucketName string `mapstructure:"bucketName"`
}

// LoadMinioConfig 用于加载 MinIO 配置
func LoadMinioConfig(path string) (*MinioConfig, error) {
	viper.SetConfigFile(path)   // 设置配置文件路径
	viper.SetConfigType("yaml") // 设置配置文件类型为 YAML

	if err := viper.ReadInConfig(); err != nil { // 读取配置文件
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	// 解析嵌套的配置项
	minioConfig := viper.Sub("drawsee.minio")
	if minioConfig == nil {
		return nil, fmt.Errorf("failed to find 'drawsee.minio' in config")
	}

	var config MinioConfig
	if err := minioConfig.Unmarshal(&config); err != nil { // 将配置文件内容映射到结构体
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	fmt.Printf("Loaded config: %+v\n", config) // 添加调试信息

	return &config, nil // 返回配置结构体
}

// NewMinioClient 创建一个新的 MinIO 客户端
func NewMinioClient(config *MinioConfig) (*minio.Client, error) {
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: false, // 如果使用 HTTPS，将此值设置为 true
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %w", err)
	}
	return client, nil
}
