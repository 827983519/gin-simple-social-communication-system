package Follow


import (
	 //"github.com/jinzhu/gorm",
	 "sbs-entrytask-template/apps/example/model"
	 "sbs-entrytask-template/agent/db"
)


func Search(user_id int) model.UserTab{
	var user model.UserTab
	DBagent.DB.Where("id = ?", "user_id").First(&user)
	return user
}



