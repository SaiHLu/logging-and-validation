package main

import (
	"github.com/SaiHLu/logging-and-validation/presenter"
	"github.com/SaiHLu/logging-and-validation/validation"
	"github.com/gin-gonic/gin"
)

type ParamRequest struct {
	ID string `uri:"id" validate:"required,min=3"`
}

type CreateRequest struct {
	Username string `json:"username" validate:"required"`
}

type UpdateRequest struct {
	Username string `json:"username" updatereq:"omitempty"`
	Email    string `json:"email" updatereq:"required"`
}

func main() {
	r := gin.Default()

	r.POST("/create", func(c *gin.Context) {
		var body CreateRequest

		if err := c.BindJSON(&body); err != nil {
			c.JSON(400, presenter.DefaultErrorJsonResponse("Invalid request body", err))
			return
		}

		if err := validation.ValidateBody(c.Request.Context(), body); err != nil {
			c.JSON(400, presenter.DefaultErrorJsonResponse("Invalid request body", err))
			return
		}

		c.JSON(200, presenter.DefaultSuccessJsonResponse(nil, "Success"))
	})

	r.PUT("/update/:id", func(c *gin.Context) {
		var (
			param ParamRequest
			body  UpdateRequest
		)

		if err := c.ShouldBindUri(&param); err != nil {
			c.JSON(400, presenter.DefaultErrorJsonResponse("Invalid param", err))
			return
		}

		if err := validation.ValidateParams(c.Request.Context(), param); err != nil {
			c.JSON(400, presenter.DefaultErrorJsonResponse("Invalid param", err))
			return
		}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(400, presenter.DefaultErrorJsonResponse("Invalid param", err))
			return
		}

		if err := validation.ValidateUpdateBody(c.Request.Context(), body); err != nil {
			c.JSON(400, presenter.DefaultErrorJsonResponse("Invalid request body", err))
			return
		}

		c.JSON(200, presenter.DefaultSuccessJsonResponse(nil, "Success"))
	})

	r.Run(":8080")
}
