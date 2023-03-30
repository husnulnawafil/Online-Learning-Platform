package access

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/husnulnawafil/online-learning-platform/global/constants"
	"github.com/husnulnawafil/online-learning-platform/global/helpers"
)

type handlerFunc func(*fiber.Ctx) error

type ReqHeader struct {
	Authorization string `reqHeader:"authorization"`
}

func Access(allowedRole ...string) func(handlerFunc) handlerFunc {
	return func(callback handlerFunc) handlerFunc {
		return func(ctx *fiber.Ctx) error {
			reqHeader := new(ReqHeader)
			if err := ctx.ReqHeaderParser(reqHeader); err != nil {
				return ctx.
					Status(http.StatusUnauthorized).
					JSON(helpers.ReponseWithoutData(
						http.StatusUnauthorized,
						"oops unable to authorize, error in parsing authorization header",
					))
			}

			if reqHeader.Authorization == "" {
				return ctx.
					Status(http.StatusUnauthorized).
					JSON(helpers.ReponseWithoutData(
						http.StatusUnauthorized,
						"oops you need to be logged in first to perform this endpoint",
					))
			}

			token := helpers.GetToken(reqHeader.Authorization)
			claims, statusCode, err := helpers.ValidateToken(token)
			if err != nil {
				return ctx.
					Status(statusCode).
					JSON(helpers.ReponseWithoutData(
						statusCode,
						err.Error(),
					))
			}

			var users map[string]interface{} = *claims
			role, ok := users["role"]
			if !ok {
				return ctx.
					Status(http.StatusForbidden).
					JSON(helpers.ReponseWithoutData(
						http.StatusForbidden,
						"oops do not trick me, your token does not mention your role",
					))
			}

			if ok := helpers.ContainString(constants.Roles, role.(string)); !ok {
				return ctx.
					Status(http.StatusForbidden).
					JSON(helpers.ReponseWithoutData(
						http.StatusForbidden,
						"oops you are not allowed",
					))
			}

			allRoleAllowed := len(allowedRole) == 1 && allowedRole[0] == constants.AllRole
			if allRoleAllowed {
				return callback(ctx)
			}

			if ok := helpers.ContainString(allowedRole, role.(string)); !ok {
				return ctx.
					Status(http.StatusForbidden).
					JSON(helpers.ReponseWithoutData(
						http.StatusForbidden,
						"oops you are forbidden for perform this action",
					))
			}

			return callback(ctx)
		}
	}
}
