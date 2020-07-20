package model

import (

)

type (
	FollowTab struct {
		ID           	 int32 `gorm:"primary_key"`
		UserId     		 int32 `json:"user_id"`
		FollowBy     	 int32 `json:"followby"`
		Ctime        	 int32  `json:"ctime"`
	}
)



func (FollowTab) TableName() string {
	return "follow_tab"
}
