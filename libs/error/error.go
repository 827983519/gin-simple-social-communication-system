package Error

import (
	"sbs-entrytask-template/libs/response/Errorcode"
)

type APIException struct {
	Errorcode        int32      `json:"errorcode"`
	Message         string    `json:"message"`
}



func (err *APIException) Error() (int32, string) {
	return err.Errorcode, err.Message
}


func New(Errorcode int32, message string) *APIException{
	return  &APIException{Errorcode: Errorcode, Message: message}
}


func ErrorParams(message string) *APIException{
	return  &APIException{Errorcode: Errorcode.PARAMS_WRONG, Message: message}
}

func ParamsValidateWrong(message string) *APIException{
	return &APIException{Errorcode: Errorcode.PARAMS_VALIDATE_WRONG, Message: message}
}

func LoginFailed(message string) *APIException{
	return &APIException{Errorcode: Errorcode.LOGIN_FAILED, Message:message}
}

func NotExist(message string) *APIException{
	return &APIException{Errorcode: Errorcode.NOT_EXIST, Message: message}
}


func DBOperateWrong(message string) *APIException{
	return &APIException{Errorcode: Errorcode.DB_OPERATE_WRONG, Message: message}
}