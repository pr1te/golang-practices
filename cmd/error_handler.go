package main

import (
	"strconv"

	goerrors "github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/errors"
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

		var e *goerrors.Error
		if goerrors.As(err, &e) {
			prefix := strconv.Itoa(e.Err.(*errors.Exception).Code)[0:3]
			status, _ = strconv.Atoi(prefix)

			exception := e.Err.(*errors.Exception)
			ctx.Status(status).JSON(&CustomError{
				Code:    exception.Code,
				Message: exception.Message,
				Errors:  exception.Errors,
				Status:  status,
				Type:    errors.ERROR_TYPE[status],
			})

			statusCategory := strconv.Itoa(e.Err.(*errors.Exception).Code)[0:1]

			if statusCategory == "4" {
				logger.Debugf("%s, code: %d\n%+v", exception.Message, exception.Code, exception.Errors)
			} else {
				logger.Errorln(e.ErrorStack())
			}

			return nil
		}

		err = ctx.Status(status).JSON(err)

		if err != nil {
			internalErr := errors.NewInternalServerError("oops! something went wrong")

			return ctx.Status(status).JSON(internalErr.Err)
		}

		return nil
	}
}
