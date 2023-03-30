package main

import (
	"github.com/gofiber/fiber/v2"
	middlewares "github.com/husnulnawafil/online-learning-platform/global/middlewares"
	courseHandlers "github.com/husnulnawafil/online-learning-platform/handlers/courses"
	statisticHandlers "github.com/husnulnawafil/online-learning-platform/handlers/statistics"
	userHandlers "github.com/husnulnawafil/online-learning-platform/handlers/users"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")
	users := api.Group("users")
	auths := api.Group("auths")
	courses := api.Group("courses")

	// COURSES
	courses.Post("/", courseHandlers.CreateCourse)
	courses.Get("/:uuid", courseHandlers.GetCourseDetail)
	courses.Get("/", courseHandlers.GetCourseList)
	courses.Put("/:uuid", courseHandlers.UpdateCourse)
	courses.Delete("/:uuid", courseHandlers.DeleteCourse)
	courses.Get("/count", statisticHandlers.CountCourse)

	// AUTHS
	auths.Post("/register", userHandlers.RegisterUser)
	auths.Post("/login", userHandlers.LoginUser)

	// USERS
	users.Delete("/:uuid", userHandlers.DeleteUser)
	users.Get("/count", statisticHandlers.CountUser)

	// ADDITIONAL
	app.Get("/test", middlewares.Access()(userHandlers.MiddlewareTest))
	app.Listen(":3000")
}
