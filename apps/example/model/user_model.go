package model

type (
	UserTab struct {
		ID           int32 	`gorm:"primary_key"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		Nickname     string `json:"nickname"`
		Image        string	`json:"image"`
		Status       int32 `json:"status"`
		Attention    int32	`json:"attention"`
		Email 		 string	`json:"email"`
		Ctime        int64	`json:"ctime"`
		Mtime   	 int64	`json:"mtime"`
	}
)


func (UserTab) TableName() string {
	return "user_tab"
}

