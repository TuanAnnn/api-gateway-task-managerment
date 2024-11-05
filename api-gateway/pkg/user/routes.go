package user

import (
	"github.com/TuanAnnn/api-gateway/pkg/common/config"
	"github.com/TuanAnnn/api-gateway/pkg/user/routes"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, c config.Config) {
	svc := InitServiceClient(&c)

	r := app.Group("/api/v1/user")
	r.Post("/open", svc.RegisterUser)
}

func (svc *ServiceClient) RegisterUser(ctx *fiber.Ctx) error {
	return routes.RegisterUser(ctx, svc.CommandClient)
}

func (svc *ServiceClient) FindUser(ctx *fiber.Ctx) error {
	return routes.FindUser(ctx, svc.QueryClient)
}
