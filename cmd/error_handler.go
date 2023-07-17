package main

import (
	"strconv"

	goerrors "github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/left-it-api/pkg/errors"
	"github.com/pr1te/left-it-api/pkg/logger"
)

type CustomError struct {
	Type    string        `json:"type"`
	Status  int           `json:"status"`
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Details []interface{} `json:"details" default:"[]"`
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
				Details: exception.Details,
				Status:  status,
				Type:    errors.ERROR_TYPE[status],
			})

			statusCategory := strconv.Itoa(e.Err.(*errors.Exception).Code)[0:1]

			if statusCategory == "4" {
				logger.Debugf("%s, code: %d, errors: %+v", exception.Message, exception.Code, exception.Details)
			} else {
				logger.Errorln(err.Error(), e.ErrorStack())
			}

			return nil
		}

		if err != nil {
			logger.Errorln(err.Error())
		}

		err = ctx.Status(status).JSON(&CustomError{
			Code:    errors.INTERNAL_SERVER_ERROR,
			Message: err.Error(),
			Details: []any{},
			Status:  status,
			Type:    errors.ERROR_TYPE[status],
		})

		if err != nil {
			internalErr := errors.NewInternalServerError("oops! something went wrong")

			if err != nil {
				logger.Errorln(err.Error(), internalErr.ErrorStack())
			}

			return ctx.Status(status).JSON(internalErr.Err)
		}

		return nil
	}
}
