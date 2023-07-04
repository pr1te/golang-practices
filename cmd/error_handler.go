package main

import (
	"strconv"

	"github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/exceptions"
	"github.com/pr1te/announcify-api/pkg/logger"
)

func errorHandler(logger *logger.Logger) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		status := fiber.StatusInternalServerError

		var e *errors.Error
		if errors.As(err, &e) {
			prefix := strconv.Itoa(e.Err.(*exceptions.Exception).Code)[0:3]
			status, _ = strconv.Atoi(prefix)

			logger.Errorln(e.ErrorStack())
		}

		err = ctx.Status(status).JSON(e.Err)

		if err != nil {
			internalErr := exceptions.NewInternalServerErrorException("oops! something went wrong")

			return ctx.Status(status).JSON(internalErr.Err)
		}

		return nil
	}
}
