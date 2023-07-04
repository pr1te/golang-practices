package main

import (
	"strconv"

	"github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/exceptions"
	"github.com/pr1te/announcify-api/pkg/logger"
)

type CustomError struct {
	Type    string        `json:"type"`
	Status  int           `json:"status"`
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Errors  []interface{} `json:"errors" default:"[]"`
}

func errorHandler(logger *logger.Logger) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		status := fiber.StatusInternalServerError

		var e *errors.Error
		if errors.As(err, &e) {
			prefix := strconv.Itoa(e.Err.(*exceptions.Exception).Code)[0:3]
			status, _ = strconv.Atoi(prefix)

			exception := e.Err.(*exceptions.Exception)
			ctx.Status(status).JSON(&CustomError{
				Code:    exception.Code,
				Message: exception.Message,
				Errors:  exception.Errors,
				Status:  status,
				Type:    exceptions.EXCEPTION_TYPE[status],
			})

			logger.Errorln(e.ErrorStack())

			return nil
		}

		err = ctx.Status(status).JSON(err)

		if err != nil {
			internalErr := exceptions.NewInternalServerErrorException("oops! something went wrong")

			return ctx.Status(status).JSON(internalErr.Err)
		}

		return nil
	}
}
