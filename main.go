package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	// Example 适合用在测试代码中
	// Development 在开发环境中使用
	// Production 用在生产环境
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// zap 也提供了便捷的方法 SugarLogger, 可以使用 printf 格式符的方式.
	// 调用 logger.Sugar() 即可创建 SugaredLogger.
	// SugaredLogger 的使用比 Logger 简单, 只是性能比 Logger 低 50% 左右, 可以用在非热点函数中.
	sugar := logger.Sugar()
	sugar.Infow("请求处理完成",
		"url", "http://example.com/api",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("处理 URL: %s", "http://example.com")
	logger.Info("用户登录",
		zap.String("username", "alice"),
		zap.Int("user_id", 12345),
		zap.Duration("latency", 10*time.Millisecond),
	)

	devLogger, _ := zap.NewDevelopment()
	defer devLogger.Sync()
	devLogger.Info("开发环境日志", zap.String("env", "local"))
}
