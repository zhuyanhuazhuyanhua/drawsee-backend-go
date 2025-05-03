# Drawsee 项目介绍

## 项目概述
Drawsee 是一个基于 Spring Boot 3 的 Web 项目，主要用于处理 AI 任务、动画渲染、知识管理等功能。项目采用 Java 17 开发，集成了多种中间件和技术栈，提供了丰富的接口模块。

## 技术栈
- **Spring Boot 3**: 项目的基础框架
- **Spring MVC**: Web 请求处理
- **RabbitMQ**: 消息队列，用于任务分发
- **MongoDB**: 非关系型数据库，存储部分业务数据
- **Redis**: 缓存和分布式锁，使用 Redisson 对接
- **MySQL**: 关系型数据库，使用 MyBatis 对接
- **Minio**: 对象存储，用于文件管理
- **Lombok**: 简化代码
- **Sa-Token**: 鉴权框架
- **Langchain4j**: AI 对话框架

## 项目结构
项目主包为 `cn.yifan.drawsee`，以下是各文件夹的详细说明：

### `constant/`
存放项目中的常量定义，如 AI 模型、任务状态、Redis 键等。这些常量用于统一管理项目中使用的固定值，避免硬编码。

### `config/`
Spring Boot 配置类，包括缓存配置、CORS 配置、RabbitMQ 配置等。每个配置类都对应一个特定的功能模块，如 `CacheConfig` 用于配置 Redis 缓存，`RabbitConfig` 用于配置 RabbitMQ 等。

### `controller/`
处理 HTTP 请求的控制器类。每个控制器类对应一个业务模块，如 `AIController` 处理 AI 任务相关的请求，`AnimationController` 处理动画渲染相关的请求。

### `service/`
业务逻辑层，处理核心业务逻辑。每个服务类对应一个业务模块，如 `AIService` 处理 AI 任务的业务逻辑，`AnimationService` 处理动画渲染的业务逻辑。

### `mapper/`
MyBatis 的 Mapper 接口，用于数据库操作。每个 Mapper 接口对应一个数据库表，如 `TaskMapper` 用于操作任务表。

### `repository/`
MongoDB 的 Repository 接口，用于 MongoDB 操作。每个 Repository 接口对应一个 MongoDB 集合，如 `KnowledgeRepository` 用于操作知识集合。

### `pojo/`
实体类，包括数据库实体、消息实体等。这些实体类用于在项目中传递数据，如 `Task` 表示任务实体，`AiTaskMessage` 表示 AI 任务消息实体。

### `exception/`
自定义异常类。这些异常类用于处理项目中的特定异常情况，如 `TaskNotFoundException` 表示任务未找到的异常。

### `annotation/`
自定义注解，如 `@PromptParam`、`@PromptResource` 等。这些注解用于简化代码，如 `@PromptParam` 用于从请求中提取参数并注入到方法参数中。

### `util/`
工具类，提供各种实用方法。这些工具类用于处理项目中常用的操作，如 `StringUtil` 提供字符串处理的方法，`DateUtil` 提供日期处理的方法。

### `worker/`
后台任务处理类，如 AI 任务处理。这些类用于处理后台任务，如 `AITaskWorker` 处理 AI 任务的后台逻辑。`worker` 模块采用了**模板方法设计模式**，通过 `WorkFlow` 基类定义任务处理的整体流程，具体的任务处理逻辑由子类实现。

#### 模板方法设计模式
`WorkFlow` 类定义了任务处理的整体流程，包括以下步骤：
1. **参数校验并初始化**：`validateAndInit` 方法用于校验任务参数并初始化上下文。
2. **更新任务状态**：`updateTaskToProcessing` 方法将任务状态更新为"处理中"。
3. **创建初始节点**：`createInitNodes` 方法用于创建任务处理的初始节点。
4. **流式输出文本**：`streamChat` 方法用于处理流式输出文本的逻辑。
5. **创建其他节点或更新节点数据**：`createOtherNodesOrUpdateNodeData` 方法用于在任务处理完成后创建其他节点或更新节点数据。
6. **更新任务状态为成功**：`updateTaskToSuccess` 方法将任务状态更新为"成功"。

具体的任务处理逻辑由 `WorkFlow` 的子类实现，如 `AnimationWorkFlow` 处理动画生成任务，`KnowledgeWorkFlow` 处理知识查询任务等。每个子类可以根据需要重写 `WorkFlow` 中的方法，以实现特定的任务处理逻辑。

#### 主要类介绍
- **`WorkFlow`**：任务处理的基类，定义了任务处理的整体流程。
- **`AITaskWorker`**：AI 任务处理类，根据任务类型调用相应的 `WorkFlow` 子类处理任务。
- **`AnimationWorkFlow`**：处理动画生成任务，包括动画分镜生成、动画代码生成和动画渲染。
- **`KnowledgeWorkFlow`**：处理知识查询任务，包括知识点的查询和相关知识点的展示。
- **`KnowledgeDetailWorkFlow`**：处理知识详情查询任务，包括知识点的详细信息和相关资源的展示。
- **`PlannerWorkFlow`**：处理任务规划任务，包括任务的分割和规划。
- **`SolverFirstWorkFlow`**：处理问题求解任务，包括问题的初步求解。
- **`SolverContinueWorkFlow`**：处理问题继续求解任务，包括问题的进一步求解。
- **`SolverSummaryWorkFlow`**：处理问题总结任务，包括问题的总结和归纳。
- **`HtmlMakerWorkFlow`**：处理 HTML 生成任务，包括 HTML 代码的生成和展示。

#### 示例代码
```java:src/main/java/cn/yifan/drawsee/worker/WorkFlow.java
public final void execute(WorkContext workContext) {
    // 参数校验并初始化
    Boolean isValid = validateAndInit(workContext);
    if (!isValid) return;
    // 更新任务状态
    updateTaskToProcessing(workContext);
    // 更新conversation
    Conversation conversation = workContext.getConversation();
    // 更新现在的时间戳
    conversation.setUpdatedAt(new Timestamp(System.currentTimeMillis()));
    conversationMapper.update(conversation);
    try {
        // 创建初始节点
        createInitNodes(workContext);
    } catch (JsonProcessingException e) {
        throw new RuntimeException(e); // TODO 异常处理
    }
    // title
    updateConvTitle(workContext);

    // 流式输出文本
    try {
        streamChat(workContext, new StreamingResponseHandler<AiMessage>() {
            @Override
            public void onNext(String token) {
                Node streamNode = workContext.getStreamNode();
                RStream<String, Object> redisStream = workContext.getRedisStream();
                Map<String, Object> textData = new ConcurrentHashMap<>();
                textData.put("nodeId", streamNode.getId());
                textData.put("content", token);
                redisStream.add(StreamAddArgs.entries(
                "type", AiTaskMessageType.TEXT, "data", textData
                ));
            }

            @Override
            public void onComplete(Response<AiMessage> response) {
                workContext.setStreamResponse(response);
                try {
                    createOtherNodesOrUpdateNodeData(workContext);
                } catch (JsonProcessingException e) {
                    throw new RuntimeException(e); // TODO 异常处理
                }
                // 发送结束消息
                if (workContext.getIsSendDone()) {
                    RStream<String, Object> redisStream = workContext.getRedisStream();
                    redisStream.add(StreamAddArgs.entries(
                    "type", AiTaskMessageType.DONE, "data", ""
                    ));
                }
                // update task and nodes
                updateTaskToSuccess(workContext);
                AiTask aiTask = workContext.getAiTask();
                aiTaskMapper.update(aiTask);
                try {
                    updateStreamNode(workContext);
                } catch (JsonProcessingException e) {
                    throw new RuntimeException(e); // TODO 异常处理
                }
                List<Node> nodesToUpdate = workContext.getNodesToUpdate();
                nodeMapper.updateDataAndIsDeletedBatch(nodesToUpdate);
                // 定时删除
                redissonClient.getQueue(RedisKey.CLEAN_AI_TASK_QUEUE_KEY).add(aiTask.getId());
            }

            @Override
            public void onError(Throwable error) {
                log.error("流式输出文本失败, taskMessage: {}", workContext.getAiTaskMessage(), error);
                RStream<String, Object> redisStream = workContext.getRedisStream();
                redisStream.add(StreamAddArgs.entries(
                "type", AiTaskMessageType.ERROR, "data", error.getMessage()
                ));
                // update task to failed
                AiTask aiTask = workContext.getAiTask();
                aiTask.setStatus(AiTaskStatus.FAILED);
                aiTask.setResult(error.getMessage());
                aiTaskMapper.update(aiTask);
            }
        });
    } catch (JsonProcessingException e) {
        throw new RuntimeException(e); // TODO 异常处理
    }
}
```

通过模板方法设计模式，`worker` 模块实现了任务处理流程的统一管理，同时允许子类根据具体任务类型实现特定的处理逻辑，提高了代码的复用性和可维护性。

### `schedule/`
定时任务类。这些类用于处理定时任务，如 `CleanTaskScheduler` 定时清理过期任务。

## 接口模块
### AI 任务模块
- **AI 任务处理**: 通过 RabbitMQ 分发和处理 AI 任务，支持多种任务类型（如知识查询、动画生成等）。任务处理流程包括任务创建、任务分发、任务执行和任务结果返回。
- **AI 模型管理**: 集成多种 AI 模型（如 Doubao、DeepseekV3），提供统一的接口调用。模型管理包括模型配置、模型调用和模型结果处理。

### 动画渲染模块
- **动画生成**: 调用 Python 服务生成动画，支持多种动画类型。动画生成流程包括动画请求、动画生成和动画结果返回。

### 知识管理模块
- **知识查询**: 提供知识点的查询和详情展示，支持多种知识来源（如 Bilibili、动画等）。知识查询流程包括知识搜索、知识详情获取和知识展示。

### 文件管理模块
- **文件存储**: 使用 Minio 进行文件存储和管理，支持图片、视频等文件类型。文件管理包括文件上传、文件下载和文件删除。

### 鉴权模块
- **用户鉴权**: 使用 Sa-Token 进行用户鉴权和权限管理。鉴权流程包括用户登录、权限验证和权限管理。

## 版本信息
- **Spring Boot**: 3.4.3
- **Java**: 17
- **Lombok**: 1.18.20
- **Sa-Token**: 1.39.0
- **Langchain4j**: 1.0.0-alpha1
- **Redisson**: 3.41.0
- **MyBatis**: 3.0.4
- **Minio**: 8.5.17

## 其他
- **缓存管理**: 使用 Redis 进行缓存管理，支持多种缓存策略。
- **消息队列**: 使用 RabbitMQ 进行任务分发，支持高并发处理。
- **定时任务**: 使用 Spring 的 `@Scheduled` 注解实现定时任务。

## 开发指南
1. **环境准备**: 确保安装 JDK 17 和 Maven。
2. **项目启动**: 使用 `mvn spring-boot:run` 启动项目。
3. **接口测试**: 使用 ApiFox 进行接口测试。
4. **部署**: 使用 Docker 进行项目部署。

## 注意事项
- 项目依赖的中间件（如 Redis、RabbitMQ、Minio）需要提前配置好。
- 配置文件中的敏感信息（如 API Key）需要妥善保管。
