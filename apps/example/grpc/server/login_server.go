package server

import (
	"golang.org/x/net/context"
	"sbs-entrytask-template/agent/cache"
	grpc_login "sbs-entrytask-template/apps/example/grpc/proto/login_proto"
	user_db "sbs-entrytask-template/apps/example/repository/user"
	"sbs-entrytask-template/libs/response/Errorcode"
)

type LoginServer struct{
	grpc_login.LoginServiceServer
}

var LOGIN_EXPIRATION_TIME  int32 = 86400


func (u *LoginServer) Login(ctx context.Context, req *grpc_login.LoginRequest) (*grpc_login.LoginResponse, error){
	username := req.Username
	password := req.Password

	user_info, err := user_db.Search_by_username(username)
	if err != nil || user_info.Password != password{
		res := &grpc_login.LoginResponse{
			Retcode:    &grpc_login.RetCodeResponse{
				Errorcode: Errorcode.LOGIN_FAILED,
				Message:   "Username or Password wrong",
			},
		}
		return res, nil
	}

	res := &grpc_login.LoginResponse{
		Id:   user_info.ID,
		Retcode: nil,
	}
	return res, nil
}



func (u *LoginServer) Register(ctx context.Context, req *grpc_login.RegisterRequest) (*grpc_login.RegisterResponse, error){
	//err := check_register_info(req)
	//if err != nil {
	//	res := &grpc_login.RegisterResponse{
	//		Retcode: &grpc_login.RetCodeResponse{
	//			Errorcode: err.Errorcode,
	//			Message:   err.Message,
	//		},
	//	}
	//	return res, nil
	//}

	send_email(req.Email)
	_ = cache.CommonCache.Set(req.Email, 1, 3600)
	res :=  &grpc_login.RegisterResponse{
			Retcode: nil,
	}
	return res, nil
}



//func check_register_info(req *grpc_login.RegisterRequest) *Error.APIException{
//	username := req.Username
//	password := req.Password
//	nickname := req.Nickname
//	email := req.Email
//	password_confirm := req.PasswordConfirmation
//	return nil
//}

func send_email(email string) {

}


