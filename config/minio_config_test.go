package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadMinioConfig(t *testing.T) {
	configPath := "C:/Users/1/Desktop/drawsee/config.yaml"
	config, err := LoadMinioConfig(configPath)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	assert.Equal(t, "http://42.193.107.127:9046", config.Endpoint)
	assert.Equal(t, "XOKQYoPekwnGasqY9cB3", config.AccessKey)
	assert.Equal(t, "Z65Sy7T2pegQFj5A7ZnRzCFeJdTBV2O4qj5GQkH3", config.SecretKey)
	assert.Equal(t, "drawsee", config.BucketName)
}

func TestNewMinioClient(t *testing.T) {
	configPath := "C:/Users/1/Desktop/drawsee/config.yaml"
	config, err := LoadMinioConfig(configPath)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	client, err := NewMinioClient(config)
	if err != nil {
		t.Fatalf("Failed to create MinIO client: %v", err)
	}

	assert.NotNil(t, client)
}
