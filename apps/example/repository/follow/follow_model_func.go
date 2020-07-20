package follow_db

import (
	"sbs-entrytask-template/agent/db"
	"sbs-entrytask-template/apps/example/model"
	"sbs-entrytask-template/apps/example/repository/user"
	"sbs-entrytask-template/libs/error"
)


func Search_user_follow_list(follow_by int32, count int32, pageno int32) ([]model.UserTab, int32){
	var follow_records []model.FollowTab
	var user_data []model.UserTab
	var total int32


	DBagent.DB.Select("count(*)").Where("follow_by = ?", follow_by).Count(&total)
	Db := DBagent.DB.Where("follow_by = ?", follow_by)
	if count > 0 && pageno > 0 {
		Db = Db.Limit(count).Offset((pageno - 1) * count)
	}
	Db.Find(&follow_records)

	var use_id_list []int32
	for _, follow_record := range follow_records {
		use_id_list = append(use_id_list, follow_record.UserId)
	}

	DBagent.DB.Select("nickname, image, attention").Where("id IN (?)", use_id_list).Find(&user_data)
	return user_data, total
}


func Search_follow_record(user_id int32, follow_by int32) (*model.FollowTab, *Error.APIException){
	var follow_record model.FollowTab
	err := DBagent.DB.Where("user_id = ? and follow_by = ?", user_id, follow_by).First(&follow_record).Error
	if err != nil{
		return nil, Error.NotExist("Cannot find follow record.")
	}
	return &follow_record, nil
}


func Create(user_id int32, follow_by int32) *Error.APIException{
	follow := model.FollowTab{
		UserId:   user_id,
		FollowBy: follow_by,
	}
	err := DBagent.DB.Create(&follow).Error

	if err != nil {
		return Error.DBOperateWrong("Cannot create follow record.")
	}
	return nil
}


func Delete(user_id int32, follow_by int32) {
	follow := model.FollowTab{
		UserId:   user_id,
		FollowBy: follow_by,
	}
	DBagent.DB.Delete(&follow)
}



func Follow_user(user_id int32, follow_by int32) *Error.APIException{
	tx := DBagent.DB.Begin()


	err := Create(user_id, follow_by)
	if err != nil {
		return err
	}

	ret_err := user_db.Increase_attention(user_id)
	if ret_err != nil {
		return ret_err
	}

	txn_err := tx.Commit().Error
	if txn_err != nil {
		return Error.DBOperateWrong("Transaction wrong.")
	}
	return nil
}


func Get_off_user(user_id int32, follow_by int32) *Error.APIException{
	tx := DBagent.DB.Begin()

	Delete(user_id, follow_by)

	ret_err := user_db.Decrease_attention(user_id)
	if ret_err != nil {
		return ret_err
	}

	txn_err := tx.Commit().Error
	if txn_err != nil {
		return Error.DBOperateWrong("Transaction wrong.")
	}
	return nil
}






