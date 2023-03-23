package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger           *zap.Logger
	ignoreUserAgents = []string{
		"elb",
	}
)

func Init() *zap.Logger {
	logger, _ = zap.NewProduction()
	return logger
}

func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			now := time.Now()
			err := next(c)
			if err != nil {
				c.Error(err)
			}
			fields := getLogField(c)
			// fieldが空の場合はskip
			if len(fields) == 0 {
				return nil
			}
			// access logの時はレイテンシーも出力
			fields = append(fields, zap.Int64("latency", time.Since(now).Nanoseconds()))
			logger.Info("", fields...)
			return nil
		}
	}
}

func Info(msg string, c echo.Context) {
	if c == nil {
		SugerInfo(msg)
		return
	}
	logger.WithOptions(zap.AddCallerSkip(1)).Info(msg, getLogField(c)...)
}

func Error(msg string, c echo.Context) {
	if c == nil {
		SugerError(msg)
		return
	}
	logger.WithOptions(zap.AddCallerSkip(1)).Error(msg, getLogField(c)...)
}

func Fatal(msg string) {
	logger.WithOptions(zap.AddCallerSkip(1)).Fatal(msg)
}

func SugerInfo(msg string) {
	logger.WithOptions(zap.AddCallerSkip(1)).Sugar().Infow(
		msg,
		"time", time.Now().String(),
	)
}

func SugerError(msg string) {
	logger.WithOptions(zap.AddCallerSkip(1)).Sugar().Errorw(
		msg,
		"time", time.Now().String(),
	)
}

func getLogField(c echo.Context) []zapcore.Field {
	req := c.Request()
	res := c.Response()
	// NOTE helthCheck等はアクセスログを吐かないように
	if isIgnoreUserAgent(req.UserAgent()) {
		return nil
	}
	var userID interface{}
	if user := c.Get("userID"); user != nil {
		userID = user
	}

	return []zapcore.Field{
		zap.String("time", time.Now().String()),
		zap.String("remote_ip", c.RealIP()),
		zap.String("real_ip", req.Header.Get("X-Real-IP")),
		zap.String("x_forwarded_for", req.Header.Get("X-Forwarded-For")),
		zap.String("host", req.Host),
		zap.String("request", fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
		zap.String("user_agent", req.UserAgent()),
		zap.Any("user_id", userID),
		zap.Int("status", res.Status),
		zap.Int64("size", res.Size),
	}
}

func isIgnoreUserAgent(str string) bool {
	for _, v := range ignoreUserAgents {
		if strings.Contains(strings.ToLower(str), v) {
			return true
		}
	}
	return false
}
