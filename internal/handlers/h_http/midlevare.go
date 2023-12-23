package h_http

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

func LoggerMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger = logger.With(
			slog.String("component", "middleware/logger.go"),
		)

		fn := func() {
			entry := logger.With(
				slog.String("method", c.Request.Method),
				slog.String("path", c.Request.URL.Path),
				slog.String("request_id", requestid.Get(c)),
			)

			t := time.Now()
			defer func() {
				entry.Info(
					"request complited",
					slog.Int("request status", c.Writer.Status()),
					slog.String("duration", time.Since(t).String()),
				)
			}()
			c.Next()
		}
		fn()
	}
}
