package user_db


import (
	"sbs-entrytask-template/apps/example/model"
	"sbs-entrytask-template/agent/db"
	"sbs-entrytask-template/libs/error"
)


func Search_by_id(user_id int32) (*model.UserTab, *Error.APIException){
	var user model.UserTab
	err := DBagent.DB.Where("id = ?", user_id).First(&user).Error

	if err != nil {
		return nil, Error.NotExist("Cannot find user.")
	}
	return &user, nil
}

func Search_by_username(username string) (*model.UserTab, *Error.APIException){
	var user model.UserTab
	err := DBagent.DB.Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, Error.NotExist("Cannot find user.")
	}
	return &user, nil
}


func Increase_attention(user_id int32) *Error.APIException{
	user_info, err := Search_by_id(user_id)

	if err != nil {
		return err
	}
	user_info.Attention += 1

	save_err := DBagent.DB.Save(&user_info).Error

	if save_err != nil {
		return Error.DBOperateWrong("Increase attention wrong")
	}
	return nil
}


func Decrease_attention(user_id int32) *Error.APIException{
	user_info, err := Search_by_id(user_id)

	if err != nil {
		return err
	}
	user_info.Attention -= 1

	save_err := DBagent.DB.Save(&user_info).Error

	if save_err != nil {
		return Error.DBOperateWrong("Increase attention wrong")
	}
	return nil

}



