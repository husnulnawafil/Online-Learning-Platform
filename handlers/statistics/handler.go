package handlers

import (
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/husnulnawafil/online-learning-platform/global/helpers"
	services "github.com/husnulnawafil/online-learning-platform/services/statistics"
)

var (
	service services.StatisticService
	once    sync.Once
)

func init() {
	once.Do(func() {
		service = services.NewStatisticService()
	})
}

func CountCourse(ctx *fiber.Ctx) error {
	isFree := ctx.QueryBool("isFree", false)
	count, err := service.CountCourse(ctx.Context(), isFree)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ResponseWithoutData(
				http.StatusInternalServerError,
				"oops error while retrieving the data , err :"+err.Error(),
			))
	}

	return ctx.
		Status(http.StatusOK).
		JSON(helpers.ReponseWithData(
			http.StatusOK,
			"success",
			map[string]int64{
				"count": count,
			},
		))
}

func CountUser(ctx *fiber.Ctx) error {
	count, err := service.CountUser(ctx.Context())
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ResponseWithoutData(
				http.StatusInternalServerError,
				"oops error while retrieving the data , err :"+err.Error(),
			))
	}

	return ctx.
		Status(http.StatusOK).
		JSON(helpers.ReponseWithData(
			http.StatusOK,
			"success",
			map[string]int64{
				"count": count,
			},
		))
}
