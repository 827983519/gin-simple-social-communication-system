package server

import (
	"golang.org/x/net/context"
	"sbs-entrytask-template/apps/example/grpc/proto/user_proto"
	follow_db "sbs-entrytask-template/apps/example/repository/follow"
	"sbs-entrytask-template/apps/example/repository/user"
	"sbs-entrytask-template/apps/example/constant"
)

type UserServer struct{
	grpc_user.UserServiceServer
}




func (u *UserServer) GetUserInfo(ctx context.Context, req *grpc_user.UserInfoRequest) (*grpc_user.UserInfoResponse, error){
	user_id := req.UserId
	user_info, err := user_db.Search_by_id(user_id)
	if err != nil {
		res := &grpc_user.UserInfoResponse{
			Retcode: &grpc_user.RetCodeResponse{
				Errorcode: err.Errorcode,
				Message:   err.Message,
			},
		}
		return res, nil
	}
	res := &grpc_user.UserInfoResponse{
		Id:        user_info.ID,
		Image:     user_info.Image,
		Username:  user_info.Username,
		Nickname:  user_info.Nickname,
		Email:     user_info.Email,
		Attention: user_info.Attention,
		Status:    user_info.Status,
		Retcode:   nil,
	}

	return res, nil
}




func (u *UserServer) ViewOtherUserInfo(ctx context.Context, req *grpc_user.ViewOtherUserInfoRequest) (*grpc_user.ViewOtherUserInfoResponse, error){
	user_id := req.UserId
	view_id := req.ViewBy

	user_info, err := user_db.Search_by_id(user_id)
	if err != nil {
		res := &grpc_user.ViewOtherUserInfoResponse{
			Retcode:    &grpc_user.RetCodeResponse{
				Errorcode: err.Errorcode,
				Message: err.Message,
			},
		}
		return res, nil
	}
	view_follow_record, _ := follow_db.Search_follow_record(view_id, user_id)
	follow_record, _ := follow_db.Search_follow_record(user_id, view_id)

	var follow_status int32
	if follow_record != nil {
		follow_status = SelfConstant.Follow_status_follow
		if view_follow_record != nil {
			follow_status = SelfConstant.Follow_status_followed_eo
		}
	}else{
		follow_status = SelfConstant.Follow_status_no
		if view_follow_record != nil {
			follow_status = SelfConstant.Follow_status_followed
		}
	}

	res := &grpc_user.ViewOtherUserInfoResponse{
		Id:        user_info.ID,
		Image:     user_info.Image,
		Nickname:  user_info.Nickname,
		Attention: user_info.Attention,
		FollowStatus: follow_status,
		Retcode:   nil,
	}
	return res, nil
}


func (u *UserServer) GetUserFollowList(ctx context.Context, req *grpc_user.UserFollowListRequest) (*grpc_user.UserFollowListResponse, error) {
	user_id := req.UserId
	pageno := req.Pageno
	count := req.Count

	_, err := user_db.Search_by_id(user_id)
	if err != nil {
		res := &grpc_user.UserFollowListResponse{
			Retcode:    &grpc_user.RetCodeResponse{
				Errorcode: err.Errorcode,
				Message: err.Message,
			},
		}
		return res, nil
	}

	user_data, total := follow_db.Search_user_follow_list(user_id, pageno, count)

	res := &grpc_user.UserFollowListResponse{
		Total:      total,
		Count:      count,
		Pageno:     pageno,
		Retcode:    nil,
	}

	var followList []*grpc_user.FollowList
	for _, user_obj := range user_data {
		followList = append(followList, &grpc_user.FollowList{
			Nickname:  user_obj.Nickname,
			Image:     user_obj.Image,
			Attention: user_obj.Attention,
		})
	}
	res.Followlist = followList
	return res, nil
}

