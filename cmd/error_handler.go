package main

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/exceptions"
)

func errorHandler(ctx *fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError

	var e *exceptions.Exception
	if errors.As(err, &e) {
		prefix := strconv.Itoa(e.Code)[0:3]
		status, _ = strconv.Atoi(prefix)
	}

	err = ctx.Status(status).JSON(err)

	if err != nil {
		internalErr := exceptions.NewInternalServerErrorException("oops! something went wrong")

		return ctx.Status(status).JSON(internalErr)
	}

	return nil
}
