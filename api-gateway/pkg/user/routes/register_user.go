package routes

import (
	"context"

	pb "github.com/TuanAnnn/api-gateway/pkg/user/pb"
	"github.com/gofiber/fiber/v2"
)

// type AccountType string
// const (
// 	Savings AccountType = "SAVINGS"
// 	Current             = "CURRENT"
// )

type UserRegisterRequestBody struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func RegisterUser(ctx *fiber.Ctx, c pb.UserCommandServiceClient) error {
	body := UserRegisterRequestBody{}

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := c.RegisterUser(context.Background(), &pb.RegisterUserRequest{
		UserName: body.UserName,
		Password: body.Password,
		Email:    body.Email,
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	return ctx.Status(int(res.Status)).JSON(&res)
}
