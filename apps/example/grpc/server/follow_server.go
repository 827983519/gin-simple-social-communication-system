package server

import (
	"golang.org/x/net/context"
	"sbs-entrytask-template/apps/example/grpc/proto/follow_proto"
	"sbs-entrytask-template/apps/example/repository/follow"
	"sbs-entrytask-template/libs/response/Errorcode"
)

type FollowServer struct{
	grpc_follow.FollowServiceServer
}



func (f *FollowServer) FollowUser(ctx context.Context, req *grpc_follow.FollowRequest) (*grpc_follow.FollowResponse, error) {
	user_id := req.UserId
	follow_by := req.FollowBy

	follow_record, _ := follow_db.Search_follow_record(user_id, follow_by)
	if follow_record != nil {
		res := &grpc_follow.FollowResponse{
			Retcode:  &grpc_follow.RetCodeResponse{
				Errorcode: Errorcode.ALREADY_EXIST,
				Message:   "Already follow user.",
			},
		}
		return res, nil
	}

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


func (f *FollowServer) GetOffUser(ctx context.Context, req *grpc_follow.FollowRequest) (*grpc_follow.FollowResponse, error) {
	user_id := req.UserId
	follow_by := req.FollowBy

	follow_record, _ := follow_db.Search_follow_record(user_id, follow_by)
	if follow_record != nil {
		err := follow_db.Get_off_user(user_id, follow_by)
		if err != nil {
			res := &grpc_follow.FollowResponse{
				Retcode: &grpc_follow.RetCodeResponse{
					Errorcode: err.Errorcode,
					Message:   err.Message,
				},
			}
			return res, nil
		}
	}
	res := &grpc_follow.FollowResponse{
		Retcode: nil,
	}
	return res, nil
}