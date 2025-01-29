package handlers

import (
	"context"
	"net/http"

	"carbonwize-backend-assessment/repository"
	"carbonwize-backend-assessment/requests"
	"carbonwize-backend-assessment/responses"

	"github.com/gofiber/fiber/v2"
)

type CalculateHandler struct {
	CalculateRepo repository.CalculateRepo
}

func NewCalculateHandler(calculateRepo repository.CalculateRepo) (auth *CalculateHandler) {
	return &CalculateHandler{calculateRepo}
}

func (ch *CalculateHandler) Calculate(c *fiber.Ctx) error {
	var inp = &requests.CalculateRequest{}
	var res = &responses.CalculateResponse{}
	var err error

	if err = c.BodyParser(inp); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err = ch.CalculateRepo.GetEmissionFactors(context.Background(), inp)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(res)
}
