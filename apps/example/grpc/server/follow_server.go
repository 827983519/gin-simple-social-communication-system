package server

import (
	"golang.org/x/net/context"
	"sbs-entrytask-template/apps/example/grpc/proto/follow_proto"
	"sbs-entrytask-template/apps/example/repository/follow"
	Error "sbs-entrytask-template/libs/error"
)

type FollowServer struct{}



func (f *FollowServer) FollowUser(ctx context.CancelFunc, req *grpc_follow.FollowRequest) (*grpc_follow.FollowResponse, *Error.APIException) {
	user_id := req.UserId
	follow_by := req.FollowBy
	err := follow_db.Follow_user(user_id, follow_by)
	if err != nil {
		res := &grpc_follow.FollowResponse{
			Retcode:  &grpc_follow.RetCodeResponse{
				Errorcode: err.Errorcode,
				Message:   err.Message,
			},
		}
		return res, nil
	}
	res := &grpc_follow.FollowResponse{
		Retcode: nil,
	}
	return res, nil
}


func (f *FollowServer) GetOffUser(ctx context.CancelFunc, req *grpc_follow.FollowRequest) (*grpc_follow.FollowResponse, *Error.APIException) {
	user_id := req.UserId
	follow_by := req.FollowBy
	err := follow_db.Get_off_user(user_id, follow_by)
	if err != nil {
		res := &grpc_follow.FollowResponse{
			Retcode:  &grpc_follow.RetCodeResponse{
				Errorcode: err.Errorcode,
				Message:   err.Message,
			},
		}
		return res, nil
	}
	res := &grpc_follow.FollowResponse{
		Retcode: nil,
	}
	return res, nil
}