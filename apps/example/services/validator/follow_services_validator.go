package Validator



type FollowForm struct {
	UserId   int32 `form:"UserId" binding:"required"`
	FollowBy int32 `form:"FollowBy" binding:"required"`
}