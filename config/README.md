# Configuration Module Documentation

## Overview
The `config` directory contains all the configuration-related code files and configuration parameter files for the project. These files are mainly used to manage and initialize various configurations, such as cache configurations, database configurations, and message queue configurations, ensuring the project can run stably and flexibly.

## File Description

### Configuration Files
| File Name | Size | Function Description |
| --- | --- | --- |
| `config.yaml` | - | Project configuration parameter file, written in YAML format, used to store various configuration parameters such as database connection information, AI model configuration, and MinIO storage configuration. |

### Core Code Files
| File Name | Size | Function Description |
| --- | --- | --- |
| `CacheConfig.go` | - | Defines the `CacheManager` structure and related functions, used to manage Redis cache operations, including getting and setting cache data. |
| `CorsConfig.go` | - | Handles CORS (Cross-Origin Resource Sharing) configuration to ensure that the application can handle cross-origin requests correctly. |
| `LangchainConfig.go` | - | Loads and parses the AI model configuration from the `config.yaml` file, including information such as the base URL, API key, and model name of each AI model. |
| `Minio.go` | - | Loads the MinIO storage configuration from the `config.yaml` file and provides functions to create a MinIO client. |
| `PromptConfig.go` | - | Defines the `PromptService` interface and related structures, used to load and manage prompt resources. |
| `rabbit.go` | - | Implements RabbitMQ-related logic, including message conversion, listener configuration, and endpoint registration. |
| `RabbitConfig.go` | - | Manages RabbitMQ configuration, including declaring exchanges and queues. |
| `RestClientConfig.go` | - | Defines the `RestClientConfig` structure and provides functions for sending HTTP POST requests. |
| `SaTokenConfig.go` | - | Handles Sa-Token authentication configuration to ensure the security of the application. |

### Test Files
| File Name | Size | Function Description |
| --- | --- | --- |
| `minio_config_test.go` | - | Tests the functionality of loading MinIO configuration, ensuring that the configuration can be correctly loaded and parsed. |
| `promptconfig_test.go` | - | Tests the `PromptConfig.go` file to ensure that the prompt resource loading and management functions work as expected. |
| `rabbit_config_test.go` | - | Tests the RabbitMQ configuration to ensure that exchanges and queues can be correctly declared. |
| `rabbit_test.go` | - | Tests the RabbitMQ listener configuration to ensure that messages can be correctly consumed and processed. |
| `RestClientConfig_test.go` | - | Tests the `RestClientConfig.go` file to ensure that HTTP POST requests can be correctly sent. |
| `SaTokenConfig_test.go` | - | Tests the Sa-Token authentication configuration to ensure the security of the application. |

## Usage Examples

### Loading AI Model Configuration
```go
package main

import (
    "fmt"
    "drawsee/config"
)

func main() {
    cfg, err := config.LoadConfig("config.yaml")
    if err != nil {
        fmt.Printf("Failed to load config: %v\n", err)
        return
    }

    fmt.Printf("Doubao Base URL: %s\n", cfg.Doubao.BaseURL)
    fmt.Printf("Doubao API Key: %s\n", cfg.Doubao.APIKey)
    fmt.Printf("Doubao Model Name: %s\n", cfg.Doubao.ModelName)
}