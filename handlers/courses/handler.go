package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/husnulnawafil/online-learning-platform/configs/cld"
	"github.com/husnulnawafil/online-learning-platform/global/helpers"
	"github.com/husnulnawafil/online-learning-platform/models"
	services "github.com/husnulnawafil/online-learning-platform/services/courses"
	"github.com/jinzhu/copier"
)

var (
	service services.CourseService
	once    sync.Once
)

func init() {
	once.Do(func() {
		service = services.NewCourseService()
	})
}

func CreateCourse(ctx *fiber.Ctx) error {
	data := &models.Course{}
	name := ctx.FormValue("name")
	if name == "" {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ResponseWithoutData(
				http.StatusBadRequest,
				"oops please fill the name",
			))
	}
	data.Name = name

	category := ctx.FormValue("category")
	if category == "" {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ResponseWithoutData(
				http.StatusBadRequest,
				"oops please fill the category",
			))
	}
	data.Category = category

	priceString := ctx.FormValue("price")
	if priceString == "" {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ResponseWithoutData(
				http.StatusBadRequest,
				"oops please provide price, fill with 0 if it's free",
			))
	}

	price, err := strconv.Atoi(priceString)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ResponseWithoutData(
				http.StatusBadRequest,
				"oops price should be number",
			))
	}
	if price == 0 {
		data.IsFree = true
	}

	if price < 0 {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ResponseWithoutData(
				http.StatusBadRequest,
				"oops price should not be negative value",
			))
	}
	data.Price = float64(price)

	uuid := uuid.Must(uuid.NewRandom())
	data.UUID = uuid.String()

	fileImageHeader, _ := ctx.FormFile("fileImage")
	if fileImageHeader != nil {
		fileImage, _ := fileImageHeader.Open()
		res, err := cld.UploadImage(ctx.Context(), "courses", uuid.String(), fileImage)
		if err != nil {
			return ctx.
				Status(http.StatusInternalServerError).
				JSON(helpers.ResponseWithoutData(
					http.StatusInternalServerError,
					"oops, error in uploading course image, err : %s"+err.Error(),
				))
		}
		data.ImageURL = res.URL
	}

	err = service.CreateCourse(ctx.Context(), data)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ResponseWithoutData(
				http.StatusInternalServerError,
				"oops got error when inserted data to database, err : "+err.Error(),
			))
	}

	courseResponse := ResponseData{}
	copier.Copy(&courseResponse, &data)
	return ctx.
		Status(http.StatusOK).
		JSON(helpers.ReponseWithData(
			http.StatusOK,
			"yay course is succesfully created",
			courseResponse,
		))
}

func GetCourseDetail(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid")
	if uuid == "" {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ResponseWithoutData(
				http.StatusBadRequest,
				"oops please provide uuid please",
			))
	}
	course, err := service.GetCourseByUUID(ctx.Context(), uuid)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ResponseWithoutData(
				http.StatusInternalServerError,
				"oops server unable to retrieve the course detail",
			))
	}

	if course.UUID == "" {
		return ctx.
			Status(http.StatusOK).
			JSON(helpers.ResponseWithoutData(
				http.StatusOK,
				"the course not found",
			))
	}

	courseResponse := new(ResponseData)
	copier.Copy(courseResponse, course)
	return ctx.
		Status(http.StatusOK).
		JSON(helpers.ReponseWithData(
			http.StatusOK,
			"yay success",
			courseResponse,
		))
}

func GetCourseList(ctx *fiber.Ctx) error {
	isFree := ctx.QueryBool("isFree", false)
	sortBy := ctx.Query("sortBy", "name")
	sortDir := ctx.Query("sortDir", "asc")
	search := ctx.Query("search")
	courses, err := service.GetCourses(ctx.Context(), sortBy, sortDir, search, isFree)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ResponseWithoutData(
				http.StatusBadRequest,
				"oops error in retrieving course list",
			))
	}

	courseList := []ResponseData{}
	copier.Copy(&courseList, &courses)
	return ctx.
		Status(http.StatusOK).
		JSON(helpers.ReponseWithData(
			http.StatusOK,
			"yay success",
			courseList,
		))
}

func UpdateCourse(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid")
	course, err := service.GetCourseByUUID(ctx.Context(), uuid)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ResponseWithoutData(
				http.StatusInternalServerError,
				"oops server unable to retrieve the course detail",
			))
	}

	if course.UUID == "" {
		return ctx.
			Status(http.StatusOK).
			JSON(helpers.ResponseWithoutData(
				http.StatusOK,
				"the course not found",
			))
	}

	name := ctx.FormValue("name")
	if name != "" {
		course.Name = name
	}

	priceString := ctx.FormValue("price")
	if priceString != "" {
		price, err := strconv.Atoi(priceString)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(helpers.ResponseWithoutData(
					http.StatusBadRequest,
					"oops price should be number",
				))
		}
		course.Price = float64(price)
		if price == 0 {
			course.IsFree = true
		}

		if price > 0 {
			course.IsFree = true
		}

		if price < 0 {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(helpers.ResponseWithoutData(
					http.StatusBadRequest,
					"oops price should not be negative value",
				))
		}
	}

	fileImageHeader, _ := ctx.FormFile("fileImage")
	if fileImageHeader != nil {
		fileImage, _ := fileImageHeader.Open()
		res, err := cld.UploadImage(ctx.Context(), "courses", uuid, fileImage)
		if err != nil {
			return ctx.
				Status(http.StatusInternalServerError).
				JSON(helpers.ResponseWithoutData(
					http.StatusInternalServerError,
					"oops, error in uploading course image, err : %s"+err.Error(),
				))
		}
		course.ImageURL = res.URL
	}
	updateCount, err := service.UpdateCourse(ctx.Context(), course.UUID, course)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ResponseWithoutData(
				http.StatusInternalServerError,
				"oops, error updated the course, err : %s"+err.Error(),
			))
	}

	if updateCount <= 0 {
		return ctx.
			Status(http.StatusOK).
			JSON(helpers.ResponseWithoutData(
				http.StatusOK,
				"unfortunately, the course has not been updated yet or not found",
			))
	}

	courseResponse := ResponseData{}
	copier.Copy(&courseResponse, &course)
	return ctx.
		Status(http.StatusOK).
		JSON(helpers.ReponseWithData(
			http.StatusOK,
			"yay course has been updated",
			courseResponse,
		))

}

func DeleteCourse(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid")
	course, err := service.GetCourseByUUID(ctx.Context(), uuid)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ResponseWithoutData(
				http.StatusInternalServerError,
				"oops server unable to retrieve the course detail",
			))
	}

	if course.UUID == "" {
		return ctx.
			Status(http.StatusOK).
			JSON(helpers.ResponseWithoutData(
				http.StatusOK,
				"the course not found",
			))
	}

	delCount, err := service.DeleteCourse(ctx.Context(), uuid)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ResponseWithoutData(
				http.StatusInternalServerError,
				"oops error to delete the course, err : "+err.Error(),
			))
	}

	if delCount <= 0 {
		return ctx.
			Status(http.StatusOK).
			JSON(helpers.ResponseWithoutData(
				http.StatusOK,
				"unfortunately, the course not yet been deleted or not found",
			))
	}

	return ctx.
		Status(http.StatusOK).
		JSON(helpers.ResponseWithoutData(
			http.StatusOK,
			fmt.Sprintf(`yay course "%s" has been deleted`, course.Name),
		))
}
