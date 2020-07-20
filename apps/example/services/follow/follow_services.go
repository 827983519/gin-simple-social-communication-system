package follow_services

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	grpc_follow "sbs-entrytask-template/apps/example/grpc/proto/follow_proto"
	Validator "sbs-entrytask-template/apps/example/services/validator"
	"sbs-entrytask-template/libs/response"
	"sbs-entrytask-template/libs/response/Errorcode"
)

const (
	address  = "localhost:50051"
)

func Follow_user(c *gin.Context)  {
	data, err := Validator.CheckFollowForm(c)
	if err != nil {
		response.Response_data(c, gin.H{}, Errorcode.PARAMS_VALIDATE_WRONG, "Parameter verification failed.")
		return
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		response.Response_data(c, gin.H{}, Errorcode.CONNECT_GRPC_FAILED, "Cannot connect grpc sever.")
		return
	}
	defer conn.Close()

	client := grpc_follow.NewFollowServiceClient(conn)

	res, _ := client.FollowUser(context.Background(), &grpc_follow.FollowRequest{UserId: data.UserId, FollowBy: data.FollowBy})

	if res.Retcode != nil {
		response.Response_data(c, gin.H{}, res.Retcode.Errorcode, res.Retcode.Message)
		return
	}

	response.Response_data(c, gin.H{}, Errorcode.OK, "")
}


func Get_off_user(c *gin.Context)  {
	data, err := Validator.CheckFollowForm(c)
	if err != nil {
		response.Response_data(c, gin.H{}, Errorcode.PARAMS_VALIDATE_WRONG, "Parameter verification failed.")
		return
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		response.Response_data(c, gin.H{}, Errorcode.CONNECT_GRPC_FAILED, "Cannot connect grpc sever.")
		return
	}
	defer conn.Close()
	client := grpc_follow.NewFollowServiceClient(conn)

	res, _ := client.GetOffUser(context.Background(), &grpc_follow.FollowRequest{
		UserId: data.UserId,
		FollowBy: data.FollowBy,
	})

	if res.Retcode != nil {
		response.Response_data(c, gin.H{}, res.Retcode.Errorcode, res.Retcode.Message)
		return
	}

	response.Response_data(c, gin.H{}, Errorcode.OK, "")

}


