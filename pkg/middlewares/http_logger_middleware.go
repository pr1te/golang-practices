package middlewares

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/logger"
	"github.com/valyala/bytebufferpool"
)

type HttpLogConfig struct {
	Logger *logger.Logger
	Next   func(c *fiber.Ctx) bool
}

func NewHttpLogger(config HttpLogConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if config.Next != nil && config.Next(c) {
			return c.Next()
		}

		start := time.Now()

		// handle request
		chainErr := c.Next()

		if chainErr != nil {
			if err := c.App().ErrorHandler(c, chainErr); err != nil {
				c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		stop := time.Now()
		buf := bytebufferpool.Get()

		_, _ = buf.WriteString(
			// ':remote-addr [:date] ":method :url HTTP/:http-version" :status :res[content-length] ":referrer" ":user-agent" ":response-time ms"
			fmt.Sprintf("%s - %s [%s] \"%s %s %s\" %d %d \"%s\" \"%s\" \"%s\"",
				c.IP(),
				"-", // this line for authenticated user
				start.UTC().Format("02/01/2006 15:04:05-0700"),
				c.Method(),
				c.OriginalURL(),
				c.Request().Header.Protocol(),
				c.Response().StatusCode(),
				c.Response().Header.ContentLength(),
				c.Request().Header.Referer(),
				c.Request().Header.UserAgent(),
				stop.Sub(start).Round(time.Millisecond),
			),
		)

		config.Logger.Infoln(buf)
		bytebufferpool.Put(buf)

		return nil
	}
}
