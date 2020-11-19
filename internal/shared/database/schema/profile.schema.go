package schema

// Profile ...
type Profile struct {
	Provider     string // google | apple
	AppID        string // hashedAppId
	ProfileImage string // Image URL
	Nickname     string
}

// TaskAlertSetting ...
type TaskAlertSetting struct {
	Hour   int `json:"hour" form:"hour" binding:"required"`
	Minute int `json:"minute" form:"minute" binding:"required"`
	Offset int `json:"offset" form:"offset" binding:"required"`
}

// ProfileWithSetting ...
type ProfileWithSetting struct {
	Profile
	TaskAlertSetting TaskAlertSetting `dynamo:",set"`
}

// ProfileSchema ...
type ProfileSchema struct {
	Key // USER#uuid PROFILE#email
	ProfileWithSetting
}
