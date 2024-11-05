package user

import (
	"fmt"

	"github.com/TuanAnnn/api-gateway/pkg/common/config"
	pb "github.com/TuanAnnn/api-gateway/pkg/user/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	CommandClient pb.UserCommandServiceClient
	QueryClient   pb.UserQueryServiceClient
}

func InitServiceClient(c *config.Config) *ServiceClient {
	queryConnection, err := grpc.Dial(c.UserQSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	cmdConnection, err := grpc.Dial(c.UserCSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &ServiceClient{
		QueryClient:   pb.NewUserQueryServiceClient(queryConnection),
		CommandClient: pb.NewUserCommandServiceClient(cmdConnection),
	}
}
