package logger

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/org-lib/bus/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Log         *zap.Logger
	logFilePath = config.Config.V.GetString("log.path")
)

func Exists(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func init() {
	logFile, _ := filepath.Split(logFilePath)
	os.MkdirAll(logFile, 0777)

	Log = NewLogger()
}

func newLoggerHook() *lumberjack.Logger {
	maxSize := config.Config.V.GetInt("log.max_size")
	maxAge := config.Config.V.GetInt("log.max_age")
	maxBackup := config.Config.V.GetInt("log.max_backup")

	if logFilePath == "" {
		logFilePath = "./log/agent.log"
	}
	if maxAge == 0 {
		maxAge = 30
	}
	if maxSize == 0 {
		maxSize = 100
	}
	if maxBackup == 0 {
		maxBackup = 30
	}

	return &lumberjack.Logger{
		Filename:   logFilePath, // 日志文件路径
		MaxSize:    maxSize,     // 每个日志文件的最大大小,单位：MB
		MaxBackups: maxBackup,   // 日志文件最多保留多少个
		MaxAge:     maxAge,      // 文件最多保留多少天
		Compress:   true,        // 是否压缩
		LocalTime:  true,
	}
}

func newLogLevel() zapcore.Level {
	logLevel := config.Config.V.GetString("log.level")
	switch logLevel {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel // 若未配置等级或配置错误将默认设置日志等级为INFO
	}
}

func NewLogger() *zap.Logger {
	hook := newLoggerHook()

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05:000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志格式为JSON,日志信息写入"操作系统标准输出"与"指定路径的日志文件"
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		// zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(hook)),
		newLogLevel(), // 设置日志记录等级
	)

	// 增加堆栈跟踪信息
	caller := zap.AddCaller()

	// 开启文件及行号
	devel := zap.Development()

	// 创建一个zap日志
	logger := zap.New(core, caller, devel)

	return logger
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		Log.Info(
			path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					Log.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					Log.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					Log.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
