package login_services

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	grpc_login "sbs-entrytask-template/apps/example/grpc/proto/login_proto"
	Validator "sbs-entrytask-template/apps/example/services/validator"
	"sbs-entrytask-template/libs/response"
	"sbs-entrytask-template/libs/response/Errorcode"
)


const (
	address  = "localhost:50051"
)



func Login(c *gin.Context){
	data, err := Validator.CheckLoginForm(c)
	if err != nil {
		response.Response_data(c, gin.H{}, Errorcode.PARAMS_VALIDATE_WRONG, "Parameter verification failed.")
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		response.Response_data(c, gin.H{}, Errorcode.CONNECT_GRPC_FAILED, "Cannot connect grpc sever.")
	}
	defer conn.Close()
	client := grpc_login.NewLoginServiceClient(conn)

	r, _ := client.Login(context.Background(), &grpc_login.LoginRequest{
				Username:  data.Username,
				Password:  data.Password,
	})
	if r.Retcode != nil {
		response.Response_data(c, gin.H{}, r.Retcode.Errorcode, r.Retcode.Message)
		return
	}

	response.Response_data(c, gin.H{}, Errorcode.OK, "")
}
