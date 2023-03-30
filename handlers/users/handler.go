package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/husnulnawafil/online-learning-platform/configs/cld"
	"github.com/husnulnawafil/online-learning-platform/global/helpers"
	"github.com/husnulnawafil/online-learning-platform/models"
	services "github.com/husnulnawafil/online-learning-platform/services/users"
	"github.com/jinzhu/copier"
)

var (
	service services.UserService
	once    sync.Once
)

func init() {
	once.Do(func() {
		service = services.NewUserService()
	})
}

func RegisterUser(ctx *fiber.Ctx) error {
	data := &models.User{}
	name := ctx.FormValue("name")
	if name == "" {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ReponseWithoutData(
				http.StatusBadRequest,
				"oops please fill the name",
			))
	}
	data.Name = name

	email := ctx.FormValue("email")
	if email == "" {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ReponseWithoutData(
				http.StatusBadRequest,
				"oops please fill the email",
			))
	}
	data.Email = email

	exist, err := service.GetUserDetailByEmail(ctx.Context(), email)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ReponseWithoutData(
				http.StatusInternalServerError,
				err.Error(),
			))
	}

	if exist.UUID != "" {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ReponseWithoutData(
				http.StatusBadRequest,
				"oops seems you have been registered yet, please login",
			))
	}

	password := ctx.FormValue("password")
	if password == "" {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ReponseWithoutData(
				http.StatusBadRequest,
				"oops please fill the password",
			))
	}

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ReponseWithoutData(
				http.StatusInternalServerError,
				"oops register can not be continue, we are not being able to secure your password",
			))
	}
	data.Password = hashedPassword

	uuid := uuid.Must(uuid.NewRandom())
	data.UUID = uuid.String()

	profileImageHeader, _ := ctx.FormFile("profile_image")
	if profileImageHeader != nil {
		profileImage, _ := profileImageHeader.Open()
		res, err := cld.UploadImage(ctx.Context(), "users", uuid.String(), profileImage)
		if err != nil {
			return ctx.
				Status(http.StatusInternalServerError).
				JSON(helpers.ReponseWithoutData(
					http.StatusInternalServerError,
					"oops, error in uploading user profile image, err : %s"+err.Error(),
				))
		}
		data.ProfileImage = res.URL
	}

	err = service.RegisterUser(ctx.Context(), data)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ReponseWithoutData(
				http.StatusInternalServerError,
				"oops, sorry we are unable registering your data, err : %s"+err.Error(),
			))
	}

	responseUser := ResponseData{}
	copier.Copy(&responseUser, &data)
	return ctx.
		Status(http.StatusOK).
		JSON(helpers.ReponseWithData(
			http.StatusOK,
			"yay registration success",
			responseUser,
		))
}

func LoginUser(ctx *fiber.Ctx) error {
	requestLogin := new(RequestLogin)
	if err := ctx.BodyParser(requestLogin); err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ReponseWithoutData(
				http.StatusInternalServerError,
				"oops, sorry we are having trouble to continue your request, err : %s"+err.Error(),
			))
	}

	email := requestLogin.Email
	password := requestLogin.Password
	if email == "" {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ReponseWithoutData(
				http.StatusBadRequest,
				"oops, please fill the email please to login",
			))
	}

	if password == "" {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ReponseWithoutData(
				http.StatusBadRequest,
				"oops, you can not login with empty password",
			))
	}

	user, err := service.GetUserDetailByEmail(ctx.Context(), email)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ReponseWithoutData(
				http.StatusInternalServerError,
				err.Error(),
			))
	}

	if user.UUID == "" {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ReponseWithoutData(
				http.StatusBadRequest,
				"oops seems you have not been registered yet, please register",
			))
	}

	storedPassword := user.Password

	if ok := helpers.CheckPasswordHash(password, storedPassword); !ok {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(helpers.ReponseWithoutData(
				http.StatusBadRequest,
				"oops, sorry you input the incorrect password",
			))
	}

	token, err := helpers.GenerateToken(user.UUID, user.Role)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ReponseWithoutData(
				http.StatusInternalServerError,
				"oops, we are unable to get you login, please try again",
			))
	}

	responseUser := ResponseData{}
	copier.Copy(&responseUser, user)

	res := map[string]interface{}{"token": token, "user": responseUser}
	return ctx.
		Status(http.StatusOK).
		JSON(helpers.ReponseWithData(
			http.StatusOK,
			"success",
			res,
		))
}

func DeleteUser(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid")
	user, err := service.GetUserDetailByUUID(ctx.Context(), uuid)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ReponseWithoutData(
				http.StatusInternalServerError,
				"oops server unable to retrieve the user detail",
			))
	}

	if user.UUID == "" {
		return ctx.
			Status(http.StatusOK).
			JSON(helpers.ReponseWithoutData(
				http.StatusOK,
				"the user not found",
			))
	}

	delCount, err := service.DeleteUser(ctx.Context(), uuid)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(helpers.ReponseWithoutData(
				http.StatusInternalServerError,
				"oops error to delete the user, err : "+err.Error(),
			))
	}

	if delCount <= 0 {
		return ctx.
			Status(http.StatusOK).
			JSON(helpers.ReponseWithoutData(
				http.StatusOK,
				"unfortunately, the user not yet been deleted or not found",
			))
	}

	return ctx.
		Status(http.StatusOK).
		JSON(helpers.ReponseWithoutData(
			http.StatusOK,
			fmt.Sprintf(`yay user "%s" has been deleted`, user.Name),
		))
}
