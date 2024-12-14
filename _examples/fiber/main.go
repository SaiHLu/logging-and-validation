package main

import (
	"github.com/SaiHLu/logging-and-validation/presenter"
	"github.com/SaiHLu/logging-and-validation/validation"
	"github.com/gofiber/fiber/v2"
)

type ParamRequest struct {
	ID string `json:"id" validate:"required,min=3"`
}

type CreateRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"omitempty"`
}

type UpdateRequest struct {
	Username string `json:"username" updatereq:"omitempty"`
	Email    string `json:"email" updatereq:"required"`
}

func main() {
	app := fiber.New()

	app.Post("/create", func(c *fiber.Ctx) error {
		var body CreateRequest

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.DefaultErrorJsonResponse("Invalid request body", err))
		}

		if err := validation.ValidateBody(c.Context(), body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.DefaultErrorJsonResponse("Invalid request body", err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.DefaultSuccessJsonResponse(nil, "Success"))
	})

	app.Put("/update/:id", func(c *fiber.Ctx) error {
		var (
			param ParamRequest
			body  UpdateRequest
		)

		if err := c.ParamsParser(&param); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.DefaultErrorJsonResponse("Invalid param", err))
		}

		if err := validation.ValidateParams(c.Context(), param); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.DefaultErrorJsonResponse("Invalid param", err))
		}

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.DefaultErrorJsonResponse("Invalid request body", err))
		}

		if err := validation.ValidateUpdateBody(c.Context(), body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.DefaultErrorJsonResponse("Invalid request body", err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.DefaultSuccessJsonResponse(nil, "Success"))
	})

	app.Listen(":3000")
}
