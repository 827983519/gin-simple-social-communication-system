package user_services

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sbs-entrytask-template/apps/example/grpc/proto/user_proto"
	Validator "sbs-entrytask-template/apps/example/services/validator"
	"sbs-entrytask-template/libs/response"
	"sbs-entrytask-template/libs/response/Errorcode"
)


const (
	address  = "localhost:50051"
)



func Get_user_info(c *gin.Context) {
	data, err := Validator.CheckGetUserForm(c)
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
	client := grpc_user.NewUserServiceClient(conn)

	r, _ := client.GetUserInfo(context.Background(), &grpc_user.UserInfoRequest{UserId: data.UserId})
	if r.Retcode != nil {
		response.Response_data(c, gin.H{}, r.Retcode.Errorcode, r.Retcode.Message)
		return
	}
	return_map := gin.H{
		"id": 		  r.Id,
		"username":	  r.Username,
		"nickname":   r.Nickname,
		"image":      r.Image,
		"attention":  r.Attention,
		"email":      r.Email,
	}
	response.Response_data(c, return_map, Errorcode.OK, "")
}


func Get_user_follow_list(c *gin.Context) {
	data, err := Validator.CheckGetUserForm(c)
	if err != nil {
		response.Response_data(c, nil, Errorcode.PARAMS_VALIDATE_WRONG, "Parameter verification failed.")
		return
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		response.Response_data(c, nil, Errorcode.CONNECT_GRPC_FAILED, "Cannot connect grpc sever.")
		return
	}
	defer conn.Close()
	client := grpc_user.NewUserServiceClient(conn)

	r, _ := client.GetUserFollowList(context.Background(), &grpc_user.UserFollowListRequest{UserId: data.UserId})
	if r.Retcode != nil {
		response.Response_data(c, gin.H{}, r.Retcode.Errorcode, r.Retcode.Message)
		return
	}

	var followList []*grpc_user.FollowList
	if len(r.Followlist) == 0 {
		followList = []*grpc_user.FollowList{}
	}else{
		followList = r.Followlist
	}
	return_map := gin.H{
		"count": r.Count,
		"pageno": r.Pageno,
		"total":  r.Total,
		"list":  followList,
	}
	response.Response_data(c, return_map, Errorcode.OK, "")
}


func View_other_user_info(c *gin.Context) {
	data, err := Validator.CheckGetOtherUserForm(c)
	if err != nil {
		response.Response_data(c, nil, Errorcode.PARAMS_VALIDATE_WRONG, "Parameter verification failed.")
		return
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		response.Response_data(c, nil, Errorcode.CONNECT_GRPC_FAILED, "Cannot connect grpc sever.")
		return
	}
	defer conn.Close()
	client := grpc_user.NewUserServiceClient(conn)

	r, _ := client.ViewOtherUserInfo(context.Background(), &grpc_user.ViewOtherUserInfoRequest{UserId: data.UserId, ViewBy: data.ViewBy})
	if r.Retcode != nil {
		response.Response_data(c, gin.H{}, r.Retcode.Errorcode, r.Retcode.Message)
		return
	}
	return_map := gin.H{
		"id": 		  r.Id,
		"follow_status":r.FollowStatus,
		"nickname":   r.Nickname,
		"image":      r.Image,
		"attention":  r.Attention,
	}
	response.Response_data(c, return_map, Errorcode.OK, "")
}
