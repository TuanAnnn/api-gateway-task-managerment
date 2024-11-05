package routes

import (
	"context"

	pb "github.com/TuanAnnn/api-gateway/pkg/user/pb"
	"github.com/gofiber/fiber/v2"
)

func FindUser(ctx *fiber.Ctx, c pb.UserQueryServiceClient) error {
	res, err := c.FindAccount(context.Background(), &pb.FindUserRequest{
		Id: ctx.Params("id"),
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	return ctx.Status(int(res.Status)).JSON(&res)
}
