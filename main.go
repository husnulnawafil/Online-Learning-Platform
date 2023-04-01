package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/husnulnawafil/online-learning-platform/global/constants"
	middlewares "github.com/husnulnawafil/online-learning-platform/global/middlewares"
	categoryHandlers "github.com/husnulnawafil/online-learning-platform/handlers/categories"
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
	categories := api.Group("categories")
	statistics := api.Group("statistics")

	// COURSES
	courses.Post("/", middlewares.Access(constants.RoleAdmin)(courseHandlers.CreateCourse))
	courses.Get("/:uuid", middlewares.Access(constants.AllRole)(courseHandlers.GetCourseDetail))
	courses.Get("/", middlewares.Access(constants.AllRole)(courseHandlers.GetCourseList))
	courses.Put("/:uuid", middlewares.Access(constants.RoleAdmin)(courseHandlers.UpdateCourse))
	courses.Delete("/:uuid", middlewares.Access(constants.RoleAdmin)(courseHandlers.DeleteCourse))

	// CATEGORIES
	categories.Post("/", middlewares.Access(constants.RoleAdmin)(categoryHandlers.CreateCategory))
	categories.Get("/:name", middlewares.Access(constants.AllRole)(categoryHandlers.GetCategoryDetail))
	categories.Get("/", middlewares.Access(constants.AllRole)(categoryHandlers.GetCategories))

	// AUTHS
	auths.Post("/register", userHandlers.RegisterUser)
	auths.Post("/login", userHandlers.LoginUser)

	// USERS
	users.Delete("/:uuid", middlewares.Access(constants.RoleAdmin)(userHandlers.DeleteUser))

	// STATISTICS
	statistics.Get("/users/count", middlewares.Access(constants.RoleAdmin)(statisticHandlers.CountUser))
	statistics.Get("/courses/count", middlewares.Access(constants.RoleAdmin)(statisticHandlers.CountCourse))

	app.Listen(":3000")
}
