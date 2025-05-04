package constant


// AI_TASK_PREFIX represents the prefix for AI task related Redis keys.
const AI_TASK_PREFIX = "ai-task:"

// CACHE_SPACE represents the cache space in Redis.
const CACHE_SPACE = "cache"

// CACHE_PREFIX represents the prefix for cache related Redis keys.
const CACHE_PREFIX = CACHE_SPACE + ":"

// INVITATION_CODE_PAGE_KEY represents the Redis key for invitation code page cache.
const INVITATION_CODE_PAGE_KEY = CACHE_PREFIX + "invitation-code-page"

// DASHBOARD_STATISTICS_KEY represents the Redis key for dashboard statistics cache.
const DASHBOARD_STATISTICS_KEY = CACHE_PREFIX + "dashboard-statistics"

// COUNT_SPACE represents the count space in Redis.
const COUNT_SPACE = "count"

// COUNT_PREFIX represents the prefix for count related Redis keys.
const COUNT_PREFIX = COUNT_SPACE + ":"

// USE_AI_COUNT_PREFIX represents the prefix for AI usage count related Redis keys.
const USE_AI_COUNT_PREFIX = COUNT_PREFIX + "use-ai:"

// CLEAN_AI_TASK_QUEUE_KEY represents the Redis key for cleaning AI task queue.
const CLEAN_AI_TASK_QUEUE_KEY = "clean-ai-task-queue"
