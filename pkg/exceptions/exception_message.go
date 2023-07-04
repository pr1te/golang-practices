package exceptions

import "github.com/gofiber/fiber/v2/utils"

var Messages = []string{
	// 5000xxx
	500000: utils.StatusMessage(500),

	// 400xxx
	400000: utils.StatusMessage(400),

	// 404xxx
	404000: utils.StatusMessage(404),
}
