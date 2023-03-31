package handlers

import (
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/husnulnawafil/online-learning-platform/global/helpers"
	"github.com/husnulnawafil/online-learning-platform/models"
	services "github.com/husnulnawafil/online-learning-platform/services/categories"
	"github.com/jinzhu/copier"
)

var (
	service services.CategoryService
	once    sync.Once
)

func init() {
	once.Do(func() {
		service = services.NewCategoryService()
	})
}

func CreateCategory(ctx *fiber.Ctx) error {
	data := new(models.CourseCategory)
	ctx.BodyParser(data)

	if data.Name == "" {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ResponseWithoutData(
				http.StatusBadRequest,
				"oops please fill the name",
			))
	}

	err := service.CreateCategory(ctx.Context(), data)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ResponseWithoutData(
				http.StatusInternalServerError,
				"oops there is something wrong when creating new category",
			))
	}

	return ctx.
		Status(http.StatusOK).
		JSON(helpers.ResponseWithoutData(
			http.StatusOK, "successfully created",
		),
		)
}

func GetCategoryDetail(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	if name == "" {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ResponseWithoutData(
				http.StatusBadRequest,
				"oops please provide name please",
			))
	}
	category, err := service.GetCategoryDetail(ctx.Context(), name)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ResponseWithoutData(
				http.StatusInternalServerError,
				"oops server unable to retrieve the category detail",
			))
	}

	if category.Name == "" {
		return ctx.
			Status(http.StatusOK).
			JSON(helpers.ResponseWithoutData(
				http.StatusOK,
				"the category not found",
			))
	}

	categoryResponse := new(ResponseData)
	copier.Copy(categoryResponse, category)
	return ctx.
		Status(http.StatusOK).
		JSON(helpers.ReponseWithData(
			http.StatusOK,
			"yay success",
			categoryResponse,
		))
}

func GetCategories(ctx *fiber.Ctx) error {
	sortBy := ctx.Query("sortBy", "popularity")
	sortDir := ctx.Query("sortDir", "asc")
	search := ctx.Query("search")

	if sortBy == "popularity" {
		sortBy = "total_subcriber"
	}

	if sortBy == "rating" {
		sortBy = "average_rating"
	}

	categories, err := service.GetCategories(ctx.Context(), sortBy, sortDir, search)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ResponseWithoutData(
				http.StatusBadRequest,
				"oops error in retrieving course list",
			))
	}

	categoryList := []ResponseData{}
	copier.Copy(&categoryList, &categories)
	return ctx.
		Status(http.StatusOK).
		JSON(helpers.ReponseWithData(
			http.StatusOK,
			"yay success",
			categoryList,
		))
}
